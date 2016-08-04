package sdb

import (
	"database/sql"
	"fmt"

	_ "github.com/go-sql-driver/mysql"
)

func New(dbtype string) *Database {
	return &Database{dbtype: dbtype}
}

type ChangeInfo struct {
	LastInsertId int64
	RowsAffected int64
}

type Database struct {
	dbtype string // just now only support "mysql"
	dsn    string // connection string
	db     *sql.DB
}

func (db *Database) Open(conn string) (err error) {
	if conn == "" {
		return fmt.Errorf("invalid db connection string")
	}
	db.dsn = conn

	db.db, err = sql.Open("mysql", conn)
	if err != nil {
		return err
	}

	err = db.db.Ping()
	if err != nil {
		return err
	}
	return nil
}

func (db *Database) Close() (err error) {
	if db.db != nil {
		err = db.db.Close()
		if err == nil {
			db.db = nil
		}
	}
	return err
}

func (db *Database) BindTable(tablename string) (t *Table, err error) {
	sql := "select * from " + tablename
	stmt, err := db.db.Prepare(sql)
	if err != nil {
		return nil, err
	}
	t = &Table{name: tablename, db: db, stmt: stmt}
	return t, nil
}

func (db *Database) CreateQuery(sql string) (q *Query, err error) {
	stmt, err := db.db.Prepare(sql)
	if err != nil {
		return nil, err
	}
	q = &Query{sql: sql, db: db, stmt: stmt}
	return q, nil
}

func (db *Database) CreateTable(sql string) (ci ChangeInfo, err error) {
	return db.Exec(sql)
}

func (db *Database) AlterTable(sql string) (ci ChangeInfo, err error) {
	return db.Exec(sql)
}

func (db *Database) Drop(tablename string) (ci ChangeInfo, err error) {
	sql := "drop table " + tablename
	return db.Exec(sql)
}

func (db *Database) Exec(sql string, args ...interface{}) (ci ChangeInfo, err error) {
	res, err := db.db.Exec(sql, args)
	if err != nil {
		return ci, err
	}
	ci.LastInsertId, _ = res.LastInsertId()
	ci.RowsAffected, _ = res.RowsAffected()
	return ci, err
}
