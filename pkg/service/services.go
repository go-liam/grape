package service

import (
	"grape/pkg/config/env"
	"grape/pkg/service/svnotice/dingtalk"
)

var (
	// 钉钉
	SvDingTalk = dingtalk.NewRobot(env.DingTalkConfig.RobotToken, env.DingTalkConfig.RobotSecret)
)
