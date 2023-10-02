package service

import (
	"errors"
	"fmt"

	"paru.net/gosimpleapp/internal/database"
	"paru.net/gosimpleapp/pkg/dto"
)

type Foo struct {
	db database.Database
}

func NewFooService(inDb database.Database) Foo {
	return Foo{
		db: inDb,
	}
}

func (sf Foo) DoSomething() error {
	sf.InsertFoo("first", 1)
	sf.InsertFoo("second", 2)

	exp, err := sf.GetFoo("2")
	if err != nil {
		return err
	}

	fmt.Printf("%v \n", exp)
	sf.GetFoo("1")
	return nil
}

func (sf Foo) InsertFoo(name string, val int) error {
	// some validation...
	return sf.db.InsertFoo(name, val)
}

func (sf Foo) GetFoo(id string) (dto.Foo, error) {
	if id == "" {
		return dto.Foo{}, errors.New("empty string")
	}

	target, err := sf.db.SelectFoo(id)
	if err != nil {
		return dto.Foo{}, errors.New("error getting data")
	}

	return newFoo(target), nil
}

func newFoo(dbFoo database.Foo) dto.Foo {
	return dto.Foo{
		Id:    dbFoo.Id,
		Name:  dbFoo.Name,
		Value: dbFoo.Value,
	}
}
