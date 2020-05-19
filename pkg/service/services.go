package service

import (
	"grape/pkg/config/env"
	"grape/pkg/service/svnotice/dingtalk"
	"grape/pkg/service/svnotice/email"
)

var (
	// 钉钉
	SvDingTalk *dingtalk.Robot
	SvMail *email.Mail
)

func init()  {
	env.InitDingTalkConfig()
	SvDingTalk = dingtalk.New(env.DingTalkConfig.RobotToken, env.DingTalkConfig.RobotSecret)
	env.InitMailConfig()
	SvMail = email.New(env.MailConfig.UserEmail,env.MailConfig.MailPassword,env.MailConfig.MailSMTPHost,env.MailConfig.MailSMTPPort)
}