package env

import (
	en "github.com/timest/env"
	"log"
)

var (
	// MysqlConfig :
	MysqlConfig *mysqlConfig
)

type mysqlConfig struct {
	// Mysql :
	MysqlHost      string `env:"MYSQL_HOST" default:"localhost"` // localhost
	MysqlPort      string `env:"MYSQL_PORT" default:"3306"`
	MysqlDBName    string `env:"MYSQL_DB_NAME" default:"ops-dev"` // 目标数据库
	MysqlUser      string `env:"MYSQL_USER" default:"root"`
	MysqlPwd       string `env:"MYSQL_PWD" default:"1234qwer"` // 1234qwer
	// other
	MysqlDBNameBoxWX string `env:"MYSQL_DB_NAME_BOX_WX" default:"wx-dev"` // 微信数据源
	MysqlDBNameBoxAct string `env:"MYSQL_DB_NAME_BOX_ACT" default:"activation-dev"` // 激活码数据源
	MysqlDBNameBoxAPI string `env:"MYSQL_DB_NAME_BOX_ACT" default:"api-dev"` // 课程数据源
	//
}

func init() {
	MysqlConfig = new(mysqlConfig)
	en.IgnorePrefix()
	err := en.Fill(MysqlConfig)
	log.Printf("[INFO] MysqlConfig :%+v\n", MysqlConfig)
	if err != nil {
		log.Printf("[ERROR] MysqlConfig :%+v\n", err)
	}
}
