package mysql

import (
	"fmt"
	"grape/internal/pkg/config"
	"grape/internal/pkg/config/env"
)

var (
	Server    *SvMySql
	ServerAPI *SvMySql
	ServerAct *SvMySql
	ServerWX  *SvMySql
)

const formatSt = "%s:%s@tcp(%s:%s)/%s?charset=utf8mb4&parseTime=true&loc=Local"

func init() {
	if config.EnvConfig.ProjectEnv == env.ConstEnvUnit { // "unit"
		unitTest()
	} else {
		initNew()
	}
}

func initNew() {
	//  "root:1234qwer@tcp(127.0.0.1:3306)/test?charset=utf8mb4"
	Server = new(SvMySql)
	Server.URL = fmt.Sprintf(formatSt, env.MysqlConfig.MysqlUser, env.MysqlConfig.MysqlPwd,
		env.MysqlConfig.MysqlHost, env.MysqlConfig.MysqlPort, env.MysqlConfig.MysqlDBName)
	Server.NewEngine()
	//Server.engine.LogMode(true)
	ServerAPI = new(SvMySql)
	ServerAPI.URL = fmt.Sprintf(formatSt, env.MysqlConfig.MysqlUser, env.MysqlConfig.MysqlPwd,
		env.MysqlConfig.MysqlHost, env.MysqlConfig.MysqlPort, env.MysqlConfig.MysqlDBNameBoxAPI)
	ServerAPI.NewEngine()
	//act
	ServerAct = new(SvMySql)
	ServerAct.URL = fmt.Sprintf(formatSt, env.MysqlConfig.MysqlUser, env.MysqlConfig.MysqlPwd,
		env.MysqlConfig.MysqlHost, env.MysqlConfig.MysqlPort, env.MysqlConfig.MysqlDBNameBoxAct)
	ServerAct.NewEngine()
	// wx
	ServerWX = new(SvMySql)
	ServerWX.URL = fmt.Sprintf(formatSt, env.MysqlConfig.MysqlUser, env.MysqlConfig.MysqlPwd,
		env.MysqlConfig.MysqlHost, env.MysqlConfig.MysqlPort, env.MysqlConfig.MysqlDBNameBoxWX)
	ServerWX.NewEngine()
}

func unitTest() {
	Server = new(SvMySql)
	Server.URL = ""
	Server.NewEngine()

	ServerWX = new(SvMySql)
	ServerWX.URL = ""
	ServerWX.NewEngine()

	ServerAct = new(SvMySql)
	ServerAct.URL = ""
	ServerAct.NewEngine()

	ServerAPI = new(SvMySql)
	ServerAPI.URL = ""
	ServerAPI.NewEngine()
}
