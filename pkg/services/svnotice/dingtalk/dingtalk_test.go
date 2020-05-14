package dingtalk

import (
	dingtalk_robot "github.com/JetBlink/dingtalk-notify-go-sdk"
	"log"
	"testing"
)

/*
https://github.com/JetBlink/dingtalk-notify-go-sdk
*/

func TestName_Ding(t *testing.T) {
	send2()
}

func send2()  {
	robot := dingtalk_robot.NewRobot("4fc5802ab76ad19377457b961f4d8c9142a0e90c768d33c73cdac304f70841a4", "")
	msg := map[string]interface{}{
		"msgtype": "text",
		"text": map[string]string{
			"content": "[提示]这是一条golang钉钉消息测试. @all ",
		},
		"at": map[string]interface{}{
			"atMobiles": []string{},
			"isAtAll":   false,
		},
	}

	//robot := dingtalk_robot.NewRobot(os.Getenv("ROBOT_TOKEN"), os.Getenv("ROBOT_SECRET"))
	if err := robot.SendMessage(msg); err != nil {
		//t.Error(err)
		log.Printf("[ERROR] %+v\n",err)
	}

}

//func testSend() {
//
//	// 钉钉机器人webhook
//	url := "https://oapi.dingtalk.com/robot/send?access_token=4fc5802ab76ad19377457b961f4d8c9142a0e90c768d33c73cdac304f70841a4"
//		//"https://oapi.dingtalk.com/robot/send?access_token=your_access_token"
//
//	// 添加头信息
//	headers := map[string]string{
//		"Content-Type":  "application/json",
//	}
//
//	/*
//	   发送简单文本消息
//	   content := make(map[string]string)
//	   content["content"] = "hello"
//	   postData := map[string]interface{}{
//	       "msgtype": "text",
//	       "text": content ,
//	   }
//	*/
//
//	// 发送markdown消息
//	content := make(map[string]string)
//	content["content"] = "[提示]hello"
//	postData := map[string]interface{}{
//		"msgtype": "markdown" ,
//		"markdown": map[string]string{
//			"title":  "ERROR",
//			"text":  "## Please check now\n " +
//				"> 1. first\n" +
//				"> 2. second\n",
//		},
//		"at": map[string]interface{}{
//			"atMobiles":[]string{"18888888888","18866666666"},
//			"isAtAll": false,
//		},
//	}
//
//	// 链式操作
//	req := curl.NewRequest()
//	resp, err := req.
//		SetUrl(url).
//		SetHeaders(headers).
//		SetPostData(postData).
//		Post()
//
//	// 返回处理
//	if err != nil {
//		fmt.Println(err)
//	} else {
//		if resp.IsOk() {
//			fmt.Println(resp.Body)
//		} else {
//			fmt.Println(resp.Raw)
//		}
//	}
//
//}