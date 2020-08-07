package page

import (
	"errors"
	"github.com/DATA-DOG/go-sqlmock"
	"github.com/go-liam/util/response"
	"github.com/stretchr/testify/assert"
	"grape/internal/pkg/database/mysql"
	models "grape/internal/pkg/model"
	"regexp"
	"testing"
)

const (
	sqlFindOne = "select .+ from ws_page"
)

var (
	tsRow  = []string{"id", "title"}
	tsItem = &Model{Title: mysql.ConstWantName, Description: ""}
)

func init() {
	tsItem.ID = mysql.ConstWantOne
}

func TestSqlSite_Create(t *testing.T) {
	db, mock, _ := mysql.MockEngine(mysql.ServerAPI) //getMockDB(t)
	defer db.Close()
	sql := "INSERT INTO `ws_page`" // regexp.QuoteMeta("SELECT * from cp_admin")

	t.Run(mysql.ConstFail, func(t *testing.T) {
		mock.ExpectBegin()
		mock.ExpectExec(sql).
			WillReturnError(mysql.ErrorDBConnect)
		mock.ExpectRollback()
		got, err := Server.Create(tsItem)
		assert.NotNil(t, err)
		assert.NotEqual(t, tsItem.ID, got)
	})

	t.Run(mysql.ConstSuccess, func(t *testing.T) {
		mock.ExpectBegin()
		mock.ExpectExec(sql).WillReturnResult(sqlmock.NewResult(mysql.ConstWantOne, mysql.ConstWantOne))
		mock.ExpectCommit()
		got, err := Server.Create(tsItem)
		assert.Nil(t, err)
		assert.EqualValues(t, mysql.ConstWantOne, got)
	})

	if err := mock.ExpectationsWereMet(); err != nil {
		t.Errorf("%s %s", mysql.ErrDbUnfulfilled, err)
	}
}

func TestDao_FindOne(t *testing.T) {
	db, mock, _ := mysql.MockEngine(mysql.ServerAPI) //getMockDB(t)
	defer db.Close()
	sql := sqlFindOne

	t.Run(mysql.ConstFail, func(t *testing.T) {
		mock.ExpectQuery(sql).
			WillReturnError(errors.New(mysql.ErrDbErrResult))
		got, err := Server.FindOne(tsItem.ID)
		assert.NotNil(t, err)
		assert.EqualValues(t, 0, got.ID)
	})

	t.Run(mysql.ConstSuccess, func(t *testing.T) {
		mock.ExpectQuery(sql). //WithArgs(item.ID).
					WillReturnRows(sqlmock.NewRows(tsRow).
						AddRow(tsItem.ID, tsItem.Title))
		got, err := Server.FindOne(tsItem.ID)
		assert.Nil(t, err)
		assert.EqualValues(t, tsItem.ID, got.ID)
	})

	if err := mock.ExpectationsWereMet(); err != nil {
		t.Errorf("%s %s", mysql.ErrDbUnfulfilled, err)
	}
}

func TestDao_FindMulti(t *testing.T) {
	db, mock, _ := mysql.MockEngine(mysql.ServerAPI) //getMockDB(t)
	defer db.Close()
	sql := sqlFindOne // regexp.QuoteMeta("SELECT * from cp_admin")
	sqlCount := regexp.QuoteMeta("SELECT count(*) FROM `ws_page`")
	page := &response.Pagination{PageSize: 10, Current: 1}
	s := &models.ListParameter{WhereSt: " 1=1 ", OrderSt: " order by id "}

	t.Run(mysql.ConstFail, func(t *testing.T) {
		mock.ExpectQuery(sqlCount).
			WillReturnError(errors.New(mysql.ErrDbErrResult))
		mock.ExpectQuery(sql).
			WillReturnError(errors.New(mysql.ErrDbErrResult))
		got, err := Server.FindMulti(page, s)
		assert.NotNil(t, err)
		assert.NotEqual(t, tsItem.ID, len(got))
	})

	t.Run(mysql.ConstSuccess, func(t *testing.T) {
		mock.ExpectQuery(sqlCount). //WithArgs(item.ID).
						WillReturnRows(sqlmock.NewRows(tsRow).
							AddRow(tsItem.ID, tsItem.Title))
		mock.ExpectQuery(sql). //WithArgs(item.ID).
					WillReturnRows(sqlmock.NewRows(tsRow).
						AddRow(tsItem.ID, tsItem.Title))
		got, err := Server.FindMulti(page, s)
		assert.Nil(t, err)
		assert.EqualValues(t, mysql.ConstWantOne, len(got))
	})

	if err := mock.ExpectationsWereMet(); err != nil {
		t.Errorf("%s %s", mysql.ErrDbUnfulfilled, err)
	}
}

func TestDao_Update(t *testing.T) {
	db, mock, _ := mysql.MockEngine(mysql.ServerAPI) //getMockDB(t)
	defer db.Close()
	//sql :=  regexp.QuoteMeta("update `cp_admin` ")
	sql := "update ws_page " //

	t.Run(mysql.ConstFail, func(t *testing.T) {
		//mock.ExpectBegin()
		mock.ExpectExec(sql).
			WillReturnError(mysql.ErrorDBConnect)
		//mock.ExpectRollback()
		got, err := Server.Update(tsItem)
		assert.NotNil(t, err)
		assert.NotEqual(t, tsItem.ID, got)
	})

	t.Run(mysql.ConstSuccess, func(t *testing.T) {
		//mock.ExpectBegin()
		mock.ExpectExec(sql).WillReturnResult(sqlmock.NewResult(mysql.ConstWantOne, mysql.ConstWantOne))
		//mock.ExpectCommit()
		got, err := Server.Update(tsItem)
		assert.Nil(t, err)
		assert.EqualValues(t, mysql.ConstWantOne, got)
	})

	if err := mock.ExpectationsWereMet(); err != nil {
		t.Errorf("%s %s", mysql.ErrDbUnfulfilled, err)
	}
}

func TestDao_UpdateStatus(t *testing.T) {
	db, mock, _ := mysql.MockEngine(mysql.ServerAPI) //getMockDB(t)
	defer db.Close()
	//sql :=  regexp.QuoteMeta("update `cp_admin` ")
	sql := "update ws_page " //

	t.Run(mysql.ConstFail, func(t *testing.T) {
		//mock.ExpectBegin()
		mock.ExpectExec(sql).
			WillReturnError(mysql.ErrorDBConnect)
		//mock.ExpectRollback()
		got, err := Server.UpdateStatus(tsItem)
		assert.NotNil(t, err)
		assert.NotEqual(t, tsItem.ID, got)
	})

	t.Run(mysql.ConstSuccess, func(t *testing.T) {
		//mock.ExpectBegin()
		mock.ExpectExec(sql).WillReturnResult(sqlmock.NewResult(mysql.ConstWantOne, mysql.ConstWantOne))
		//mock.ExpectCommit()
		got, err := Server.UpdateStatus(tsItem)
		assert.Nil(t, err)
		assert.EqualValues(t, mysql.ConstWantOne, got)
	})

	if err := mock.ExpectationsWereMet(); err != nil {
		t.Errorf("%s %s", mysql.ErrDbUnfulfilled, err)
	}
}
