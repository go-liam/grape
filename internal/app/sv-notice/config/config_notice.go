package config

import (
	"grape/pkg/config/env"
	"grape/pkg/service/svnotice/dingtalk"
	"grape/pkg/service/svnotice/email"
	"log"
)

var (
	// 钉钉
	SvDingTalk *dingtalk.Robot
	SvMail     *email.Mail
)

func init() {
	//env.InitDingTalkConfig()
	SvDingTalk = dingtalk.New(env.DingTalkConfig.RobotToken, env.DingTalkConfig.RobotSecret)
	//env.InitMailConfig()
	SvMail = email.New(env.MailConfig.UserEmail, env.MailConfig.MailPassword, env.MailConfig.MailSMTPHost, env.MailConfig.MailSMTPPort)
	log.Printf("[INFO] DingTalkConfig %+v\n", env.DingTalkConfig)
	log.Printf("[INFO] MailConfig %+v\n", env.MailConfig)
}
