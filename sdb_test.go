package sdb

import (
	"math/rand"
	"testing"
	"time"
)

func TestFunction(t *testing.T) {
	rand.Seed(time.Now().Unix())

	db := New("mysql")
	err := db.Open("root:root@tcp(127.0.0.1:3306)/world")
	if err != nil {
		t.Error(err)
	}

	sql := "select * from tb_new_table"
	q, err := db.CreateQuery(sql)
	if err != nil {
		t.Error(err)
	}

	err = q.Exec()
	if err != nil {
		t.Error(err)
	}

	type result struct {
		SId  int `orm:"id"`
		Name string
	}
	res := &result{}
	err = q.Next(res)
	if err != nil {
		t.Error(err)
	}
	t.Log(*res)

	err = q.Next(res)
	if err != nil {
		t.Error(err)
	}
	t.Log(*res)

	err = q.Next(res)
	if err != nil {
		t.Error(err)
	}
	t.Log(*res)

	rs := make([]interface{}, 10)
	rs[0] = res
	rs, err = q.BatchNext(rs, 10)
	if err != nil {
		t.Error(err)
	}

	t.Log(rs)
}

/*
type ChangeInfo struct {
}

type DataBase interface {
	BindTable(tablename string) (Table, error)
	CreateQuery(sql string) (Query, error)
	CreateTable(sql string) (ChangeInfo, error)
	AlterTable(sql string) (ChangeInfo, error)
	Drop(tablename string) (ChangeInfo, error)
	Exec(sql string, args ...interface{}) (ChangeInfo, error)
}

type Query interface {
	Select(sql string, args ...interface{}) (Cursor, error)
	SelectOne(obj interface{}) error
	SelectBatch(obj []interface{}) error
}

type Table interface {
	Query
	Exec(sql string, args ...interface{}) (ChangeInfo, error)
	Insert(sql string, args ...interface{}) (ChangeInfo, error)
	InsertObject(obj interface{}) (ChangeInfo, error)
}

type Cursor interface {
	Next(obj interface{}) error
	GetAll(obj []interface{}) error
}
*/
