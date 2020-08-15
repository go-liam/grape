package menu

import (
	"errors"
	"github.com/DATA-DOG/go-sqlmock"
	"github.com/stretchr/testify/assert"
	"grape/configs/testdata"
	"grape/internal/pkg/database/mysql"
	"testing"
)

const (
	sqlFindOne = "select .+ from ws_menu"
)

var (
	tsRow  = []string{"id", "name"}
	tsItem = &Model{Name: testdata.ConstWantString, ParentID: 0}
)

func init() {
	tsItem.ID = testdata.ConstWantOne
}

func TestSqlSite_Create(t *testing.T) {
	db, mock, _ := mysql.MockEngine(mysql.ServerAPI) //getMockDB(t)
	defer db.Close()
	sql := "INSERT INTO `ws_menu`" // regexp.QuoteMeta("SELECT * from cp_admin")

	t.Run(testdata.ConstFail, func(t *testing.T) {
		mock.ExpectBegin()
		mock.ExpectExec(sql).
			WillReturnError(testdata.ErrorDBConnect)
		mock.ExpectRollback()
		got, err := Server.Create(tsItem)
		assert.NotNil(t, err)
		assert.NotEqual(t, tsItem.ID, got)
	})

	t.Run(testdata.ConstSuccess, func(t *testing.T) {
		mock.ExpectBegin()
		mock.ExpectExec(sql).WillReturnResult(sqlmock.NewResult(testdata.ConstWantOne, testdata.ConstWantOne))
		mock.ExpectCommit()
		got, err := Server.Create(tsItem)
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
		got, err := Server.FindOne(tsItem.ID)
		assert.NotNil(t, err)
		assert.EqualValues(t, 0, got.ID)
	})

	t.Run(testdata.ConstSuccess, func(t *testing.T) {
		mock.ExpectQuery(sql). //WithArgs(item.ID).
					WillReturnRows(sqlmock.NewRows(tsRow).
						AddRow(tsItem.ID, tsItem.Name))
		got, err := Server.FindOne(tsItem.ID)
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
		got, err := Server.FindMulti()
		assert.NotNil(t, err)
		assert.NotEqual(t, tsItem.ID, len(got))
	})

	t.Run(testdata.ConstSuccess, func(t *testing.T) {
		mock.ExpectQuery(sql). //WithArgs(item.ID).
					WillReturnRows(sqlmock.NewRows(tsRow).
						AddRow(tsItem.ID, tsItem.Name))
		got, err := Server.FindMulti()
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
	sql := "update ws_menu " //

	t.Run(testdata.ConstFail, func(t *testing.T) {
		//mock.ExpectBegin()
		mock.ExpectExec(sql).
			WillReturnError(testdata.ErrorDBConnect)
		//mock.ExpectRollback()
		got, err := Server.Update(tsItem)
		assert.NotNil(t, err)
		assert.NotEqual(t, tsItem.ID, got)
	})

	t.Run(testdata.ConstSuccess, func(t *testing.T) {
		//mock.ExpectBegin()
		mock.ExpectExec(sql).WillReturnResult(sqlmock.NewResult(testdata.ConstWantOne, testdata.ConstWantOne))
		//mock.ExpectCommit()
		got, err := Server.Update(tsItem)
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
	sql := "update ws_menu " //

	t.Run(testdata.ConstFail, func(t *testing.T) {
		//mock.ExpectBegin()
		mock.ExpectExec(sql).
			WillReturnError(testdata.ErrorDBConnect)
		//mock.ExpectRollback()
		got, err := Server.UpdateStatus(tsItem)
		assert.NotNil(t, err)
		assert.NotEqual(t, tsItem.ID, got)
	})

	t.Run(testdata.ConstSuccess, func(t *testing.T) {
		//mock.ExpectBegin()
		mock.ExpectExec(sql).WillReturnResult(sqlmock.NewResult(testdata.ConstWantOne, testdata.ConstWantOne))
		//mock.ExpectCommit()
		got, err := Server.UpdateStatus(tsItem)
		assert.Nil(t, err)
		assert.EqualValues(t, testdata.ConstWantOne, got)
	})

	if err := mock.ExpectationsWereMet(); err != nil {
		t.Errorf("%s %s", testdata.ErrDbUnfulfilled, err)
	}
}