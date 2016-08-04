package sdb

import (
	"database/sql"

	_ "github.com/go-sql-driver/mysql"
)

type Table struct {
	name string
	db   *Database
	stmt *sql.Stmt
}

func (t *Table) Select(sql string, args ...interface{}) (c *Cursor, err error) {

	return nil, err
}

func (t *Table) SelectOne(obj interface{}) (err error) {
	return err

}

func (t *Table) SelectBatch(obj []interface{}) (err error) {
	return err
}

func (t *Table) Exec(sql string, args ...interface{}) (ci ChangeInfo, err error) {
	return ci, err
}

func (t *Table) Insert(sql string, args ...interface{}) (ci ChangeInfo, err error) {
	return ci, err
}

func (t *Table) InsertObject(obj interface{}) (ci ChangeInfo, err error) {
	return ci, err
}
