package env

type MailConfig struct {
	UserEmail string // 授权的用户 1234@qq.com
	MailSMTPPort string // 端口号，:587  :25也行
	MailPassword string // 邮箱的授权码
	MailSMTPHost string //"smtp.qq.com"
}

var Mail = new(MailConfig) //*MailConfig

//func init()  {
//	//Mail = new(MailConfig)
//	// 在这里设置你的账号
//}