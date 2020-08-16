package role

import (
	"errors"
	"github.com/DATA-DOG/go-sqlmock"
	"github.com/stretchr/testify/assert"
	"grape/configs/testdata"
	"grape/internal/pkg/database/mysql"
	"testing"
)

const (
	sqlFindOne = "select .+ from rb_role"
)

var (
	tsRow  = []string{"id", "created_at"}
	tsItem = &Model{Name: testdata.ConstWantString, ID: 1, CreatedAt: 0}
)

func TestSqlRole_Create(t *testing.T) {
	db, mock, _ := mysql.MockEngine(mysql.ServerAPI) //getMockDB(t)
	defer db.Close()
	sql := "INSERT INTO `rb_role`" // regexp.QuoteMeta("SELECT * from cp_admin")

	t.Run(testdata.ConstFail, func(t *testing.T) {
		mock.ExpectBegin()
		mock.ExpectExec(sql).
			WillReturnError(testdata.ErrorDBConnect)
		mock.ExpectRollback()
		got, err := new(SrvRole).Create(tsItem)
		assert.NotNil(t, err)
		assert.NotEqual(t, tsItem.ID, got)
	})

	t.Run(testdata.ConstSuccess, func(t *testing.T) {
		mock.ExpectBegin()
		mock.ExpectExec(sql).WillReturnResult(sqlmock.NewResult(testdata.ConstWantOne, testdata.ConstWantOne))
		mock.ExpectCommit()
		got, err := new(SrvRole).Create(tsItem)
		assert.Nil(t, err)
		assert.EqualValues(t, testdata.ConstWantOne, got)
	})

	if err := mock.ExpectationsWereMet(); err != nil {
		t.Errorf("%s %s", testdata.ErrDbUnfulfilled, err)
	}
}

func TestDao_FindOne(t *testing.T) {
	db, mock, _ := mysql.MockEngine(mysql.ServerAPI) //getMockDB(t)
	defer db.Close()
	sql := sqlFindOne

	t.Run(testdata.ConstFail, func(t *testing.T) {
		mock.ExpectQuery(sql).
			WillReturnError(errors.New(testdata.ErrDbErrResult))
		got, err := new(SrvRole).FindOne(tsItem.ID)
		assert.NotNil(t, err)
		assert.EqualValues(t, 0, got.ID)
	})

	t.Run(testdata.ConstSuccess, func(t *testing.T) {
		mock.ExpectQuery(sql). //WithArgs(item.ID).
					WillReturnRows(sqlmock.NewRows(tsRow).
						AddRow(tsItem.ID, tsItem.CreatedAt))
		got, err := new(SrvRole).FindOne(tsItem.ID)
		assert.Nil(t, err)
		assert.EqualValues(t, tsItem.ID, got.ID)
	})

	if err := mock.ExpectationsWereMet(); err != nil {
		t.Errorf("%s %s", testdata.ErrDbUnfulfilled, err)
	}
}

func TestDao_FindMulti(t *testing.T) {
	db, mock, _ := mysql.MockEngine(mysql.ServerAPI) //getMockDB(t)
	defer db.Close()
	sql := sqlFindOne // regexp.QuoteMeta("SELECT * from cp_admin")

	t.Run(testdata.ConstFail, func(t *testing.T) {
		mock.ExpectQuery(sql).
			WillReturnError(errors.New(testdata.ErrDbErrResult))
		got, err := new(SrvRole).FindMultiByNil()
		assert.NotNil(t, err)
		assert.NotEqual(t, tsItem.ID, len(got))
	})

	t.Run(testdata.ConstSuccess, func(t *testing.T) {
		mock.ExpectQuery(sql). //WithArgs(item.ID).
					WillReturnRows(sqlmock.NewRows(tsRow).
						AddRow(tsItem.ID, tsItem.CreatedAt))
		got, err := new(SrvRole).FindMultiByNil()
		assert.Nil(t, err)
		assert.EqualValues(t, testdata.ConstWantOne, len(got))
	})

	if err := mock.ExpectationsWereMet(); err != nil {
		t.Errorf("%s %s", testdata.ErrDbUnfulfilled, err)
	}
}

func TestDao_Update(t *testing.T) {
	db, mock, _ := mysql.MockEngine(mysql.ServerAPI) //getMockDB(t)
	defer db.Close()
	//sql :=  regexp.QuoteMeta("update `cp_admin` ")
	sql := "update rb_role " //

	t.Run(testdata.ConstFail, func(t *testing.T) {
		//mock.ExpectBegin()
		mock.ExpectExec(sql).
			WillReturnError(testdata.ErrorDBConnect)
		//mock.ExpectRollback()
		got, err := new(SrvRole).Update(tsItem)
		assert.NotNil(t, err)
		assert.NotEqual(t, tsItem.ID, got)
	})

	t.Run(testdata.ConstSuccess, func(t *testing.T) {
		//mock.ExpectBegin()
		mock.ExpectExec(sql).WillReturnResult(sqlmock.NewResult(testdata.ConstWantOne, testdata.ConstWantOne))
		//mock.ExpectCommit()
		got, err := new(SrvRole).Update(tsItem)
		assert.Nil(t, err)
		assert.EqualValues(t, testdata.ConstWantOne, got)
	})

	if err := mock.ExpectationsWereMet(); err != nil {
		t.Errorf("%s %s", testdata.ErrDbUnfulfilled, err)
	}
}

func TestDao_UpdateStatus(t *testing.T) {
	db, mock, _ := mysql.MockEngine(mysql.ServerAPI) //getMockDB(t)
	defer db.Close()
	//sql :=  regexp.QuoteMeta("update `cp_admin` ")
	sql := "update rb_role " //

	t.Run(testdata.ConstFail, func(t *testing.T) {
		//mock.ExpectBegin()
		mock.ExpectExec(sql).
			WillReturnError(testdata.ErrorDBConnect)
		//mock.ExpectRollback()
		got, err := new(SrvRole).UpdateStatus(tsItem)
		assert.NotNil(t, err)
		assert.NotEqual(t, tsItem.ID, got)
	})

	t.Run(testdata.ConstSuccess, func(t *testing.T) {
		//mock.ExpectBegin()
		mock.ExpectExec(sql).WillReturnResult(sqlmock.NewResult(testdata.ConstWantOne, testdata.ConstWantOne))
		//mock.ExpectCommit()
		got, err := new(SrvRole).UpdateStatus(tsItem)
		assert.Nil(t, err)
		assert.EqualValues(t, testdata.ConstWantOne, got)
	})

	if err := mock.ExpectationsWereMet(); err != nil {
		t.Errorf("%s %s", testdata.ErrDbUnfulfilled, err)
	}
}
