package sdb

type DataBaseInterface interface {
	Open(conn string) error
	Close() error
	BindTable(tablename string) (Table, error)
	CreateQuery(sql string) (Query, error)
	CreateTable(sql string) (ChangeInfo, error)
	AlterTable(sql string) (ChangeInfo, error)
	Drop(tablename string) (ChangeInfo, error)
	Exec(sql string, args ...interface{}) (ChangeInfo, error)
}

type QueryInterface interface {
	Select(sql string, args ...interface{}) (Cursor, error)
	SelectOne(obj interface{}) error
	SelectBatch(obj []interface{}) error
}

type TableInterface interface {
	QueryInterface
	Exec(sql string, args ...interface{}) (ChangeInfo, error)
	Insert(sql string, args ...interface{}) (ChangeInfo, error)
	InsertObject(obj interface{}) (ChangeInfo, error)
}

type CursorInterface interface {
	Next(obj interface{}) error
	GetAll(obj []interface{}) error
}
