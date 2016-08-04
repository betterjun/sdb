package sdb

import (
	"database/sql"

	_ "github.com/go-sql-driver/mysql"
)

type Cursor struct {
	rows *sql.Rows
	db   *Database
}

func (c *Cursor) Next(obj interface{}) (err error) {
	return err
}

func (c *Cursor) GetAll(obj []interface{}) (err error) {
	return err
}
