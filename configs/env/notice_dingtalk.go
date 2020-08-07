package env

import (
	en "github.com/timest/env"
	"log"
)

type dingTalkConfig struct {
	RobotToken  string `env:"NOTICE_DINGTALK_ROBOT_TOKEN1" default:""`
	RobotSecret string `env:"NOTICE_DINGTALK_ROBOT_SECRET" default:""`
}

var DingTalkConfig = new(dingTalkConfig)

func InitDingTalkConfig() {
	en.IgnorePrefix()
	err := en.Fill(DingTalkConfig)
	if DingTalkConfig.RobotToken == "" {
		DingTalkConfig.RobotToken = DefaultDingTalkConfigRobotToken
		DingTalkConfig.RobotSecret = DefaultDingTalkConfigRobotSecret
	}
	if ConstEnvUnit == EnvConfig.ProjectEnv {
		DingTalkConfig.RobotToken = ""
	}
	log.Printf("[INFO] DingTalkConfig :%+v\n", DingTalkConfig)
	if err != nil {
		log.Printf("[ERROR] DingTalkConfig :%+v\n", err)
	}
}

func init() {
	InitDingTalkConfig()
}
