package dingtalk

import (
	"grape/pkg/config/testdata"
	"log"
	"testing"
)

var (
	ROBOT_TOKEN  = testdata.DingTalkConfigRobotToken
	ROBOT_SECRET = testdata.DingTalkConfigRobotSecret
)

func TestRobot_SendMessage(t *testing.T) {
	println(ROBOT_TOKEN)
	tsSendMessage()
}
func tsSendMessage() {
	//t.SkipNow()
	msg := map[string]interface{}{
		"msgtype": "text",
		"text": map[string]string{
			"content": "这是一条golang钉钉消息测试2.",
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
