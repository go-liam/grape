package env

import (
	en "github.com/timest/env"
	"log"
)

type dingTalkConfig struct {
	RobotToken string `env:"NOTICE_DINGTALK_ROBOT_TOKEN1" default:""`
	RobotSecret string `env:"NOTICE_DINGTALK_ROBOT_SECRET" default:""`
}

var DingTalkConfig *dingTalkConfig

func init()  {
	DingTalkConfig = new(dingTalkConfig)
	en.IgnorePrefix()
	err := en.Fill(DingTalkConfig)
	if ConstEnvUnit == EnvConfig.ProjectEnv{
		DingTalkConfig.RobotToken = ""
	}
	log.Printf("[INFO] DingTalkConfig :%+v\n", DingTalkConfig)
	if err != nil {
		log.Printf("[ERROR] DingTalkConfig :%+v\n", err)
	}
}