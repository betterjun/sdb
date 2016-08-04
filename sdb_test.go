package sdb

import (
	"math/rand"
	"testing"
	"time"
)

func TestFunction(t *testing.T) {
	rand.Seed(time.Now().Unix())
}

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
