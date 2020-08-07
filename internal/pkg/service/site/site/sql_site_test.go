package site

import (
	"github.com/DATA-DOG/go-sqlmock"
	"github.com/stretchr/testify/assert"
	"grape/internal/pkg/database/mysql"
	"testing"
)

const (
	sqlFindOne = "select .+ from cp_admin"
)

var (
	tsRow  = []string{"ID", "Name"}
	tsItem = &Model{Title: mysql.ConstWantName, Description: ""}
)

func init() {
	tsItem.ID = mysql.ConstWantOne
}

func TestSqlSite_Create(t *testing.T) {
	db, mock, _ := mysql.MockEngine(mysql.ServerAPI) //getMockDB(t)
	defer db.Close()
	sql := "INSERT INTO `ws_site`" // regexp.QuoteMeta("SELECT * from cp_admin")

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
