package env

import (
	en "github.com/timest/env"
	"log"
)

type mailConfig struct {
	UserEmail    string `env:"NOTICE_Mail_USER_EMAIL" default:"***@qq.com"` // 授权的用户 1234@qq.com
	MailPassword string `env:"NOTICE_Mail_PASSWORD" default:"jdqhpfjzmzywbh**"` // 邮箱的授权码 xxx
	MailSMTPHost string `env:"NOTICE_Mail_SMTP_HOST" default:"smtp.**.com"` //"smtp.qq.com"
	MailSMTPPort string `env:"NOTICE_Mail_SMTP_PORT" default:":587"` // 端口号，:587  :25也行
}

var MailConfig = new(mailConfig) //*MailConfig

func InitMailConfig() {
	en.IgnorePrefix()
	err := en.Fill(MailConfig)
	if ConstEnvUnit == EnvConfig.ProjectEnv {
		MailConfig.UserEmail = ""
	}
	log.Printf("[INFO] MailConfig :%+v\n", MailConfig)
	if err != nil {
		log.Printf("[ERROR] MailConfig :%+v\n", err)
	}
}
