package dingtalk

import (
	"grape/internal/pkg/config/env"
	"log"
	"testing"
)

var (
	robotToken  = ""
	robotSecret = ""
)

func getConfig() {
	robotToken = env.DefaultDingTalkConfigRobotToken
	robotSecret = env.DefaultDingTalkConfigRobotSecret
	println(robotToken, robotSecret)
}

func TestRobot_SendMessage(t *testing.T) {
	getConfig()
	tsSendMessage()
}

func tsSendMessage() {
	//t.SkipNow()
	msg := map[string]interface{}{
		"msgtype": "text",
		"text": map[string]string{
			"content": "这是一条golang钉钉消息测试.",
		},
		"at": map[string]interface{}{
			"atMobiles": []string{},
			"isAtAll":   true, //false,
		},
	}
	//robot := New(os.Getenv(robotToken), os.Getenv(robotSecret))
	robot := New(robotToken, robotSecret)
	if err := robot.SendMessage(msg); err != nil {
		log.Printf("%+v\n", err)
	}
}

func TestRobot_SendTextMessage(t *testing.T) {
	getConfig()
	tsSendTextMessage()
}

func tsSendTextMessage() {
	robot := New(robotToken, robotSecret)
	if err := robot.SendTextMessage("[提示]普通文本消息", []string{}, false); err != nil {
		log.Printf("%+v\n", err)
	}
}

func TestRobot_SendMarkdownMessage(t *testing.T) {
	getConfig()
	tsSendMarkdownMessage()
}

func tsSendMarkdownMessage() {
	robot := New(robotToken, robotSecret)
	err := robot.SendMarkdownMessage(
		"[提示]Markdown Test Title",
		"### Markdown 测试消息 **@指定手机号和@所有人** \n* 百度: [baidu](https://www.baidu.com/)\n* 一张图片\n ![](http://image.biaobaiju.com/uploads/20181025/19/1540467434-IhiJNbyXak.jpg)",
		[]string{},
		true,
	)
	if err != nil {
		log.Printf("%+v\n", err)
	}
}

func TestRobot_SendLinkMessage(t *testing.T) {
	getConfig()
	tsSendLinkMessage()
}

func tsSendLinkMessage() {
	robot := New(robotToken, robotSecret)
	err := robot.SendLinkMessage(
		"[提示]Link Test Title",
		"这是一条链接测试消息",
		"https://github.com/go-liam",
		"http://image.biaobaiju.com/uploads/20181025/19/1540467434-IhiJNbyXak.jpg",
	)
	if err != nil {
		log.Printf("%+v\n", err)
	}
}
