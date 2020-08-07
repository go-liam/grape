package mysql

import (
	"database/sql"
	"errors"
	"github.com/DATA-DOG/go-sqlmock"
)

const (
	ErrDbErrResult   = "数据库错误"
	ErrDbUnfulfilled = "there were unfulfilled expectations:"

	ConstSuccess  = "ok"
	ConstFail     = "fail"
	ConstWant0    = 0
	ConstWantOne  = 1
	ConstWantMore = 2
	ConstWantName = "tsName"
)

var ErrorDBConnect = errors.New("数据库错误")

func MockEngine(server *SvMySql) (*sql.DB, sqlmock.Sqlmock, error) {
	db, mock, err := sqlmock.New()
	if err != nil {
		panic(err)
	}
	server.EngineMock(db)
	return db, mock, err
}
