package test

import (
	"testing"

	"github.com/golang/mock/gomock"
	mock_database "paru.net/gosimpleapp/internal/test/mocks/database"
)

type FooServiceMocks struct {
	Db *mock_database.MockDatabase
}

func SetUpFooServiceTest(t *testing.T) (*gomock.Controller, FooServiceMocks) {
	ctrl := gomock.NewController(t)
	dbm := mock_database.NewMockDatabase(ctrl)
	mocks := FooServiceMocks{
		Db: dbm,
	}
	return ctrl, mocks
}
