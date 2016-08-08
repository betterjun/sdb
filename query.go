package sdb

import (
	"database/sql"
	"fmt"
	"reflect"
	"strings"

	_ "github.com/go-sql-driver/mysql"
)

type Query struct {
	sql  string
	db   *Database
	stmt *sql.Stmt
	rows *sql.Rows
	cols []string
}

func (q *Query) unpack(ptr interface{}) (scanArgs []interface{}) {
	fields := make(map[string]interface{})
	v := reflect.ValueOf(ptr).Elem() // the struct variable
	for i := 0; i < v.NumField(); i++ {
		fieldInfo := v.Type().Field(i) // a reflect.StructField
		tag := fieldInfo.Tag           // a reflect.StructTag
		name := tag.Get("orm")
		if name == "" {
			name = strings.ToLower(fieldInfo.Name)
		}
		// take the addr and as interface{}, this is required by sql Scan
		fields[name] = v.Field(i).Addr().Interface()
	}

	for _, name := range q.cols {
		f := fields[name]
		scanArgs = append(scanArgs, f)
	}
	fmt.Println(scanArgs)
	return scanArgs
}

func (q *Query) Exec(args ...interface{}) (err error) {
	if len(args) == 0 {
		q.rows, err = q.stmt.Query()
	} else {
		q.rows, err = q.stmt.Query(args)
	}
	if err != nil {
		return err
	}
	//c = &Cursor{rows: rows, db: q.db}

	// 根据反射，处理obj的字段，在rows.Column中就输出
	q.cols, err = q.rows.Columns()
	if err != nil {
		return err
	}
	return err
}

func (q *Query) Next(ptr interface{}) (err error) {
	scanArgs := q.unpack(ptr)
	if scanArgs == nil {
		return fmt.Errorf("no fields found in the objptr")
	}

	if q.rows.Next() {
		err = q.rows.Scan(scanArgs...)
		fmt.Println("ptr=", ptr)
		return err
	} else {
		return fmt.Errorf("the last one found")
	}
}

// objs must have one object
func (q *Query) BatchNext(objs []interface{}, size int) (ret []interface{}, err error) {
	arg := objs[0]
	scanArgs := q.unpack(arg)
	if scanArgs == nil {
		return nil, fmt.Errorf("no fields found in the objptr")
	}

	var c int = 0
	for q.rows.Next() {
		fmt.Println("BatchNext")
		err = q.rows.Scan(scanArgs...)
		if err != nil {
			fmt.Println("return", objs)
			return objs, err
		}

		//objs = append(objs, arg)
		objs[c] = arg
		c++
	}
	fmt.Println(objs)
	return objs, err
}
