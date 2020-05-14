package email

import (
	"log"
	"grape/pkg/config"
	"net/smtp"
	"strings"
)

const TypeMailHtml = "html"

type MailInfo struct {
	To       []string //  []string{"***87685@qq.com"}
	From     string   // 1234@qq.com , QQ免费的 无法改，授权用户
	FromName string   // 发送者昵称
	Subject  string   //"我是-标题"
	Body     string   // 发送内容
	Type     string   // html
	//Msg string
}

func Send(item *MailInfo) (bool, error) {
	from := config.Mail.UserEmail
	auth := smtp.PlainAuth("", config.Mail.UserEmail, config.Mail.MailPassword, config.Mail.MailSMTPHost)
	contentType := "Content-Type: text/plain; charset=UTF-8"
	if item.Type == TypeMailHtml {
		contentType = "Content-Type: text/html; charset=UTF-8"
	}
	msg := []byte("To: " + strings.Join(item.To, ",") + "\r\nFrom: " + item.FromName +
		"<" + from + ">\r\nSubject: " + item.Subject + "\r\n" + contentType + "\r\n\r\n" + item.Body)
	addr := config.Mail.MailSMTPHost + config.Mail.MailSMTPPort
	err := smtp.SendMail(addr, auth, from, item.To, msg)
	if err != nil {
		log.Printf("[ERROR] send mail error: %v", err)
		return false, err
	}
	return true, nil
}
