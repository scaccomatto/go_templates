package service

import (
	"testing"

	"github.com/golang/mock/gomock"
	"github.com/stretchr/testify/assert"
	"paru.net/gosimpleapp/internal/test"
)

// We need to mock the dependency(db) if we want to be able to run the tests
func Test_GivenFooWhenDoingSomethingThenNoError(t *testing.T) {

	t.Helper()
	ctrl, mocks := test.SetUpFooServiceTest(t)
	defer ctrl.Finish()

	mocks.Db.EXPECT().
		InsertFoo("aa", gomock.Any()).
		Times(1)

	fooSvc := NewFooService(mocks.Db)

	err := fooSvc.InsertFoo("aa", 1)
	assert.Nil(t, err)
}
