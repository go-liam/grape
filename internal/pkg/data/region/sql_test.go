package region

import (
	"errors"
	"github.com/DATA-DOG/go-sqlmock"
	"github.com/go-liam/util/response"
	"github.com/stretchr/testify/assert"
	"grape/configs/testdata"
	"grape/internal/pkg/database/mysql"
	models "grape/internal/pkg/model"
	"regexp"
	"testing"
)

const (
	sqlFindOne = "select .+ from ar_user_region"
)

var (
	tsRow  = []string{"user_id", "created_at"}
	tsItem = &Model{UserID: testdata.ConstWantOne, RegionID: 0, CreatedAt: 0}
)

func TestSqlRegion_Create(t *testing.T) {
	db, mock, _ := mysql.MockEngine(mysql.ServerAPI) //getMockDB(t)
	defer db.Close()
	sql := "INSERT INTO `ar_user_region`" // regexp.QuoteMeta("SELECT * from cp_admin")

	t.Run(testdata.ConstFail, func(t *testing.T) {
		mock.ExpectBegin()
		mock.ExpectExec(sql).
			WillReturnError(testdata.ErrorDBConnect)
		mock.ExpectRollback()
		got, err := Server.Create(tsItem)
		assert.NotNil(t, err)
		assert.NotEqual(t, tsItem.UserID, got)
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
		got, err := Server.FindOne(tsItem.UserID)
		assert.NotNil(t, err)
		assert.EqualValues(t, testdata.ConstWant0, got.UserID)
	})

	t.Run(testdata.ConstSuccess, func(t *testing.T) {
		mock.ExpectQuery(sql). //WithArgs(item.ID).
					WillReturnRows(sqlmock.NewRows(tsRow).
						AddRow(tsItem.UserID, tsItem.CreatedAt))
		got, err := Server.FindOne(tsItem.UserID)
		assert.Nil(t, err)
		assert.EqualValues(t, tsItem.UserID, got.UserID)
	})

	if err := mock.ExpectationsWereMet(); err != nil {
		t.Errorf("%s %s", testdata.ErrDbUnfulfilled, err)
	}
}

func TestDao_FindMulti(t *testing.T) {
	db, mock, _ := mysql.MockEngine(mysql.ServerAPI) //getMockDB(t)
	defer db.Close()
	sql := sqlFindOne // regexp.QuoteMeta("SELECT * from cp_admin")
	sqlCount := regexp.QuoteMeta("SELECT count(*) FROM `ar_user_region`")
	page := &response.Pagination{PageSize: 10, Current: 1}
	s := &response.ListParameter{WhereSt: " 1=1 ", OrderSt: " order by id "}

	t.Run(testdata.ConstFail, func(t *testing.T) {
		mock.ExpectQuery(sqlCount).
			WillReturnError(errors.New(testdata.ErrDbErrResult))
		mock.ExpectQuery(sql).
			WillReturnError(errors.New(testdata.ErrDbErrResult))
		got, err := Server.FindMulti(page, s)
		assert.NotNil(t, err)
		assert.NotEqual(t, testdata.ConstWantOne, len(got))
	})

	t.Run(testdata.ConstSuccess, func(t *testing.T) {
		mock.ExpectQuery(sqlCount). //WithArgs(item.ID).
						WillReturnRows(sqlmock.NewRows(tsRow).
							AddRow(tsItem.UserID, tsItem.CreatedAt))
		mock.ExpectQuery(sql). //WithArgs(item.ID).
					WillReturnRows(sqlmock.NewRows(tsRow).
						AddRow(tsItem.UserID, tsItem.CreatedAt))
		got, err := Server.FindMulti(page, s)
		assert.Nil(t, err)
		assert.EqualValues(t, testdata.ConstWantOne, len(got))
	})

	if err := mock.ExpectationsWereMet(); err != nil {
		t.Errorf("%s %s", testdata.ErrDbUnfulfilled, err)
	}
}
