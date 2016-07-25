package sdb

func New(conn string) (*DataBase, error) {
	return nil, nil
}

type ChangeInfo struct {
}

type DataBase interface {
	BindTable(tablename string) (Table, error)
	CreateTable(sql string) (ChangeInfo, error)
	AlterTable(sql string) (ChangeInfo, error)
	Drop(tablename string) (ChangeInfo, error)
	Exec(sql string, args ...interface{}) (ChangeInfo, error)
}

type Table interface {
	Exec(sql string, args ...interface{}) (ChangeInfo, error)

	Insert(sql string, args ...interface{}) (ChangeInfo, error)
	InsertObject(obj interface{}) (ChangeInfo, error)
	Select(sql string, args ...interface{}) (Cursor, error)
	SelectOne(obj interface{}) error
	SelectBatch(obj []interface{}) error
}

type Cursor interface {
	Next(obj interface{}) error
	GetAll(obj []interface{}) error
}
