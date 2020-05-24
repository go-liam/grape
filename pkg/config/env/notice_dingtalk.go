package env

import (
	en "github.com/timest/env"
	"log"
)

type dingTalkConfig struct {
	RobotToken  string `env:"NOTICE_DINGTALK_ROBOT_TOKEN1" default:"4fc5802ab76ad19377457b961f**"`
	RobotSecret string `env:"NOTICE_DINGTALK_ROBOT_SECRET" default:"SEC52300a280cbbb18103f433ce82ab7**"`
}

var DingTalkConfig = new(dingTalkConfig)

func InitDingTalkConfig() {
	en.IgnorePrefix()
	err := en.Fill(DingTalkConfig)
	if ConstEnvUnit == EnvConfig.ProjectEnv {
		DingTalkConfig.RobotToken = ""
	}
	log.Printf("[INFO] DingTalkConfig :%+v\n", DingTalkConfig)
	if err != nil {
		log.Printf("[ERROR] DingTalkConfig :%+v\n", err)
	}
}
