package email

import (
	"log"
	"net/smtp"
	"strings"
)

const (
	TypeMailHtml  = "html"
	TypeMailPlain = "plain"
)

type Mail struct {
	UserEmail    string // 授权的用户 1234@qq.com
	MailSMTPPort string // 端口号，:587  :25也行
	MailPassword string // 邮箱的授权码
	MailSMTPHost string //"smtp.qq.com"
}

type MailInfo struct {
	To       []string //  []string{"***87685@qq.com"}
	From     string   // 1234@qq.com , QQ免费的 无法改，授权用户
	FromName string   // 发送者昵称
	Subject  string   //"我是-标题"
	Body     string   // 发送内容
	Type     string   // html
	//Msg string
}

func New(userEmail string, mailPassword string, mailSMTPHost string, mailSMTPPort string, ) *Mail {
	return &Mail{
		UserEmail:    userEmail,
		MailSMTPPort: mailSMTPPort,
		MailPassword: mailPassword,
		MailSMTPHost: mailSMTPHost,
	}
}

func (sv *Mail) Send(item *MailInfo) (bool, error) {
	from := sv.UserEmail
	auth := smtp.PlainAuth("", sv.UserEmail, sv.MailPassword, sv.MailSMTPHost)
	contentType := "Content-Type: text/plain; charset=UTF-8"
	if item.Type == TypeMailHtml {
		contentType = "Content-Type: text/html; charset=UTF-8"
	}
	msg := []byte("To: " + strings.Join(item.To, ",") + "\r\nFrom: " + item.FromName +
		"<" + from + ">\r\nSubject: " + item.Subject + "\r\n" + contentType + "\r\n\r\n" + item.Body)
	addr := sv.MailSMTPHost + sv.MailSMTPPort
	err := smtp.SendMail(addr, auth, from, item.To, msg)
	if err != nil {
		log.Printf("[ERROR] send mail error: %v", err)
		return false, err
	}
	return true, nil
}
