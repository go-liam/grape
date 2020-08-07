package svgrpc

import (
	pb2 "grape/api/proto/notice"
	"grape/configs/env"
	"log"
	"testing"
)

func TestNoticeDingTalkMarkdown(t *testing.T) {
	v1, v2 := NoticeDingTalkMarkdown(&pb2.DingTalkReq{
		Token:       env.DingTalkConfig.RobotToken,
		RobotSecret: env.DingTalkConfig.RobotSecret,
		Title:       "title-标题",
		Text:        "text-这是一条golang钉钉消息测试-3",
		IsAtAll:     true,
	})
	log.Printf("v1=%+v\n", v1)
	log.Printf("v1=%+v\n", v2)
}

func TestNoticeEmail(t *testing.T) {
	v1, v2 := NoticeEmail(&pb2.EmailReq{
		To:       env.DefaultNoticeEmailTo,
		FromName: "xxxFrom",
		Subject:  "你好-2-svgrpc",
		Type:     "html",
		Body:     `<h5 style="color: red; font-size: 16px;">测试-标题</h5>`,
	})
	log.Printf("v1=%+v\n", v1)
	log.Printf("v1=%+v\n", v2)
}
