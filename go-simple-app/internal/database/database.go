package database

import (
	"time"
)

type Database interface {
	SelectFoo(id string) (Foo, error)
	InsertFoo(name string, value int) error
}

type Foo struct {
	Id         string
	Name       string
	Value      int
	LastUpdate time.Duration
}

func New() Database {
	return newMysqlDb()
}
