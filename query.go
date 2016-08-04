package sdb

import (
	"database/sql"

	_ "github.com/go-sql-driver/mysql"
)

type Query struct {
	sql  string
	db   *Database
	stmt *sql.Stmt
	rows *sql.Rows
	cols []string
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

func (q *Query) Next(obj interface{}) (err error) {

	//row := q.stmt.QueryRow(args)
	//row.Scan()
	// 根据反射，处理obj的字段
	return err

}

func (q *Query) BatchNext(obj []interface{}, size int) (err error) {

	return err
}
