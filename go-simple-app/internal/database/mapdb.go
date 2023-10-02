package database

import (
	"errors"
	"strconv"
)

type MapDb struct {
	db map[string]Foo
}

func NewMapDb() *MapDb {
	m := make(map[string]Foo)
	dbmap := MapDb{
		db: m,
	}
	return &dbmap
}

func (m *MapDb) SelectFoo(id string) (Foo, error) {
	target, ok := m.db[id]
	if !ok {
		return Foo{}, errors.New("id not found")
	}
	return target, nil
}

func (m *MapDb) InsertFoo(name string, value int) error {
	//fooId := uuid.New().String()
	fooId := strconv.FormatInt(int64(value), 10)
	target := Foo{
		Id:    fooId,
		Name:  name,
		Value: value,
	}
	m.db[fooId] = target

	return nil
}
