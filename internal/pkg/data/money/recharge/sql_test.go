package recharge

import (
	"errors"
	"github.com/DATA-DOG/go-sqlmock"
	"github.com/go-liam/util/response"
	"github.com/stretchr/testify/assert"
	"grape/configs/testdata"
	"grape/internal/pkg/database/mysql"
	"regexp"
	"testing"
)

const (
	sqlFindOne = "select .+ from um_recharge"
)

var (
	tsRow  = []string{"id", "created_at"}
	tsItem = &Model{ID: testdata.ConstWantOne, CreatedAt: 0}
)

func TestSqlRecharge_Create(t *testing.T) {
	db, mock, _ := mysql.MockEngine(mysql.ServerAPI) //getMockDB(t)
	defer db.Close()
	sql := "INSERT INTO `um_recharge`" // regexp.QuoteMeta("SELECT * from cp_admin")

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

func TestSqlRecharge_FindOne(t *testing.T) {
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
						AddRow(tsItem.ID, tsItem.CreatedAt))
		got, err := Server.FindOne(tsItem.ID)
		assert.Nil(t, err)
		assert.EqualValues(t, tsItem.ID, got.ID)
	})

	if err := mock.ExpectationsWereMet(); err != nil {
		t.Errorf("%s %s", testdata.ErrDbUnfulfilled, err)
	}
}

func TestSqlRecharge_FindMulti(t *testing.T) {
	db, mock, _ := mysql.MockEngine(mysql.ServerAPI) //getMockDB(t)
	defer db.Close()
	sql := sqlFindOne // regexp.QuoteMeta("SELECT * from cp_admin")
	sqlCount := regexp.QuoteMeta("SELECT count(*) FROM `um_recharge`")
	page := &response.Pagination{PageSize: 10, Current: 1}
	s := &response.ListParameter{WhereSt: " 1=1 ", OrderSt: " order by id "}

	t.Run(testdata.ConstFail, func(t *testing.T) {
		mock.ExpectQuery(sqlCount).
			WillReturnError(errors.New(testdata.ErrDbErrResult))
		mock.ExpectQuery(sql).
			WillReturnError(errors.New(testdata.ErrDbErrResult))
		got, err := Server.FindMulti(page, s)
		assert.NotNil(t, err)
		assert.NotEqual(t, tsItem.ID, len(got))
	})

	t.Run(testdata.ConstSuccess, func(t *testing.T) {
		mock.ExpectQuery(sqlCount). //WithArgs(item.ID).
						WillReturnRows(sqlmock.NewRows(tsRow).
							AddRow(tsItem.ID, tsItem.CreatedAt))
		mock.ExpectQuery(sql). //WithArgs(item.ID).
					WillReturnRows(sqlmock.NewRows(tsRow).
						AddRow(tsItem.ID, tsItem.CreatedAt))
		got, err := Server.FindMulti(page, s)
		assert.Nil(t, err)
		assert.EqualValues(t, testdata.ConstWantOne, len(got))
	})

	if err := mock.ExpectationsWereMet(); err != nil {
		t.Errorf("%s %s", testdata.ErrDbUnfulfilled, err)
	}
}

func TestSqlRecharge_Update(t *testing.T) {
	db, mock, _ := mysql.MockEngine(mysql.ServerAPI) //getMockDB(t)
	defer db.Close()
	//sql :=  regexp.QuoteMeta("update `cp_admin` ")
	sql := "update um_recharge " //

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

func TestSqlRecharge_UpdateStatus(t *testing.T) {
	db, mock, _ := mysql.MockEngine(mysql.ServerAPI) //getMockDB(t)
	defer db.Close()
	//sql :=  regexp.QuoteMeta("update `cp_admin` ")
	sql := "update um_recharge " //

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
