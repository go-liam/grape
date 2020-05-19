package email

import (
	"grape/pkg/config/env"
	"log"
	"testing"
)

func TestNameEmail(t *testing.T) {
	sendTest()
}

func sendTest() {
	log.Printf("config=%+v\n", env.Mail)
	i := new(MailInfo)
	i.To = []string{"liam.jiang@***.com"}
	i.Body = `
<html>
		<body>
		<h1 style="color: red; font-size: 16px;">标题</h1>
		<h1 >测试的：12345678</h1>
		</body>
		</html>
`
	i.FromName = "帅哥007"
	i.Subject = "我是我-标题10"
	i.Type = TypeMailHtml
	Send(i)
}
