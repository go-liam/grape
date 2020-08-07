package mysql

import (
	"fmt"
	"grape/configs/env"
	"log"
	"testing"
)

func TestMySQL_Connect(t *testing.T) {
	//initMySQLData(engine, configs.EnvConfig.MysqlDBName)
	//initMySQLData(EngineOrg, configs.EnvConfig.MysqlDBNameOrg)
	sv := new(SvMySql)
	sv.URL = "root:1234qwer@tcp(127.0.0.1:3306)/test?charset=utf8mb4" //"%s:%s@tcp(%s:%s)/%s?charset=utf8mb4"
	sv.NewEngine()
	println("statue:", sv.IsConnect)
	//log.Printf("v:%+v\n",sv.engine)
}

func TestHealth(t *testing.T) {
	v1 := Health()
	log.Printf("v1= %+v\n", v1)
}

func TestName_New(t *testing.T) {
	sv := new(SvMySql)
	sv.URL = fmt.Sprintf("%s:%s@tcp(%s:%s)/%s?charset=utf8mb4&parseTime=true&loc=Local", env.MysqlConfig.MysqlUser, env.MysqlConfig.MysqlPwd,
		env.MysqlConfig.MysqlHost, env.MysqlConfig.MysqlPort, env.MysqlConfig.MysqlDBName)
	sv.NewEngine()
	log.Printf("v1= %+v\n", sv.IsConnect)
}
