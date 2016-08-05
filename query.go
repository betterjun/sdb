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
	// Build map of fields keyed by effective name.
	fields := make(map[string]reflect.Value)
	v := reflect.ValueOf(ptr).Elem() // the struct variable
	for i := 0; i < v.NumField(); i++ {
		fieldInfo := v.Type().Field(i) // a reflect.StructField
		tag := fieldInfo.Tag           // a reflect.StructTag
		name := tag.Get("orm")
		if name == "" {
			name = strings.ToLower(fieldInfo.Name)
		}
		fields[name] = v.Field(i)
	}

	// Update struct field for each parameter in the request.
	for _, name := range q.cols {
		f := fields[name]
		if !f.IsValid() {
			continue // ignore unrecognized orm parameters
		}
		scanArgs = append(scanArgs, fields[name].Addr())
	}
	return nil
}

func (q *Query) Exec(args ...interface{}) (err error) {
	rows, err := q.stmt.Query(args)
	if err != nil {
		return err
	}
	q.rows = rows
	//c = &Cursor{rows: rows, db: q.db}

	// 根据反射，处理obj的字段，在rows.Column中就输出
	q.cols, err = rows.Columns()
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
		return err
	} else {
		return fmt.Errorf("the last one found")
	}
}

// objs must have one object
func (q *Query) BatchNext(objs []interface{}, size int) (err error) {
	scanArgs := q.unpack(objs[0])
	if scanArgs == nil {
		return fmt.Errorf("no fields found in the objptr")
	}

	for q.rows.Next() {
		err = q.rows.Scan(scanArgs...)
		if err != nil {
			return err
		}

		objs = append(objs, objs[0])
	}
	return err
}
