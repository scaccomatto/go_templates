package database

type MysqlDb struct {
}

func newMysqlDb() *MysqlDb {
	// init
	return &MysqlDb{}
}

func (ms MysqlDb) SelectFoo(id string) (Foo, error) {
	return Foo{}, nil
}

func (ms MysqlDb) InsertFoo(name string, value int) error {
	return nil
}
