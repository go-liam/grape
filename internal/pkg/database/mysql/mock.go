package mysql

import (
	"database/sql"
	"github.com/DATA-DOG/go-sqlmock"
	"github.com/jinzhu/gorm"
	"grape/configs"
	"grape/configs/env"
)

func MockEngine(server *SvMySql) (*sql.DB, sqlmock.Sqlmock, error) {
	db, mock, err := sqlmock.New()
	if err != nil {
		panic(err)
	}
	server.EngineMock(db)
	return db, mock, err
}

func (sv *SvMySql) EngineMock(db *sql.DB) {
	var err error
	sv.engine, err = gorm.Open("mysql", db)
	if err != nil {
		panic(err)
	}
	if configs.EnvConfig.ProjectEnv != env.ConstEnvUnit {
		sv.engine.LogMode(true)
	}
	sv.IsConnect = true
	//println("[INFO] EngineMock")
}
