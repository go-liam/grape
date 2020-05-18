package dingtalk

import (
	"log"
	"testing"
)

var (
	ROBOT_TOKEN  = "4fc5802ab76ad19377457b961f4d8c9142a0e90c768d33c73cdac304f70841a4"
	ROBOT_SECRET = ""
)

func TestRobot_SendMessage(t *testing.T) {
	tsSendMessage()
}
func tsSendMessage() {
	//t.SkipNow()
	msg := map[string]interface{}{
		"msgtype": "text",
		"text": map[string]string{
			"content": "[提示]这是一条golang钉钉消息测试.",
		},
		"at": map[string]interface{}{
			"atMobiles": []string{},
			"isAtAll":   true ,//false,
		},
	}
	//robot := NewRobot(os.Getenv(ROBOT_TOKEN), os.Getenv(ROBOT_SECRET))
	robot := NewRobot(ROBOT_TOKEN, ROBOT_SECRET)
	if err := robot.SendMessage(msg); err != nil {
		log.Printf("%+v\n", err)
	}
}

func TestRobot_SendTextMessage(t *testing.T) {
	tsSendTextMessage()
}

func tsSendTextMessage() {
	robot := NewRobot(ROBOT_TOKEN, ROBOT_SECRET)
	if err := robot.SendTextMessage("[提示]普通文本消息", []string{}, false); err != nil {
		log.Printf("%+v\n", err)
	}
}

func TestRobot_SendMarkdownMessage(t *testing.T) {
	tsSendMarkdownMessage()
}
func tsSendMarkdownMessage() {
	robot := NewRobot(ROBOT_TOKEN, ROBOT_SECRET)
	err := robot.SendMarkdownMessage(
		"[提示]Markdown Test Title",
		"### Markdown 测试消息 **@指定手机号和@所有人** \n* 百度: [baidu](https://www.baidu.com/)\n* 一张图片\n ![](http://image.biaobaiju.com/uploads/20181025/19/1540467434-IhiJNbyXak.jpg)",
		[]string{},
		false,
	)
	if err != nil {
		log.Printf("%+v\n", err)
	}
}

func TestRobot_SendLinkMessage(t *testing.T) {
	tsSendLinkMessage()
}
func tsSendLinkMessage() {
	robot := NewRobot(ROBOT_TOKEN, ROBOT_SECRET)
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
