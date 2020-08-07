package mysql

import (
	"github.com/jinzhu/gorm"
	"grape/configs"
	"log"
	"time"
	// 引用数据库驱动初始化
	_ "github.com/jinzhu/gorm/dialects/mysql"
)

const connectTime = 600 //秒
const (
	MaxOpenConns    = 150
	MaxIdleConns    = 100
	ConnMaxLifetime = 100
)

type SvMySql struct {
	URL         string //
	IsConnect   bool
	connectTime int64
	engine      *gorm.DB
	Prefix      string // table 前缀
}

func (sv *SvMySql) Engine() *gorm.DB {
	if !sv.IsConnect {
		sv.AutoConnect()
	}
	return sv.engine
}

// AutoConnect :
func (sv *SvMySql) AutoConnect() {
	println("[INFO] mysql AutoConnect - 1")
	if sv.IsConnect {
		return
	}
	now := time.Now().Unix()
	// 10*60 秒
	if now-sv.connectTime < connectTime {
		return
	}
	println("[INFO] mysql AutoConnect - 2")
	sv.NewEngine()
}

// NewEngine :
func (sv *SvMySql) NewEngine() bool {
	sv.connectTime = time.Now().Unix()
	var err error
	sv.engine, err = gorm.Open("mysql", sv.URL)
	if err != nil {
		log.Println("[ERROR] connect to mysql fail, ", sv.URL, err)
		//panic(err)
		sv.IsConnect = false
		return false
	}
	logSQL := true
	if configs.EnvConfig.ProjectEnv == "prod" {
		logSQL = false
	}
	sv.engine.LogMode(logSQL)
	if sv.engine != nil && sv.engine.DB() != nil {
		sv.engine.DB().SetConnMaxLifetime(ConnMaxLifetime * time.Second)
		sv.engine.DB().SetMaxOpenConns(MaxOpenConns)
		sv.engine.DB().SetMaxIdleConns(MaxIdleConns)
	}
	// 注册创建/更新回调函数，自动插入时间戳
	//engine.Callback().Create().Replace("gorm:update_time_stamp", updateTimeStampForCreateCallback)
	//engine.Callback().Update().Replace("gorm:update_time_stamp", updateTimeStampForUpdateCallback)
	// 禁止update/delete传空对象
	sv.engine.BlockGlobalUpdate(true)
	sv.IsConnect = true
	log.Println("[INFO] mysql connect ok")
	return true
}
