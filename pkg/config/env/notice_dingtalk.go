package env

import (
	en "github.com/timest/env"
	"log"
)

type dingTalkConfig struct {
	RobotToken  string `env:"NOTICE_DINGTALK_ROBOT_TOKEN1" default:"4fc5802ab76ad19377457b961f4d8c9142a0e90c768d33c73cdac304f7084**"`
	RobotSecret string `env:"NOTICE_DINGTALK_ROBOT_SECRET" default:"SEC52300a280cbbb18103f433ce4a8d668850a3ecc9898b47d9cd895d2082ab7**"`
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
