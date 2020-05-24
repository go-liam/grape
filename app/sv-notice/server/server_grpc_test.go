package server

import (
	"context"
	"google.golang.org/grpc"
	"grape/pkg/config/env"
	"grape/pkg/config/testdata"
	pb2 "grape/proto/notice"
	"log"
	"os"
	"testing"
	"time"
)

const (
	tsGRPCPort    = ":4401"
	tsGRPCAddress = "localhost" + tsGRPCPort
	defaultName   = "client:"
)

func init()  {
	go RunServerGRPC(tsGRPCPort)
}

func TestSvGRPC_Health(t *testing.T) {
	time.Sleep(time.Microsecond * 10)
	testHealth()
}

func TestSvGRPC_Email(t *testing.T) {
	//go RunServerGRPC(tsGRPCPort)
	time.Sleep(time.Microsecond * 10)
	testEmail()
}

func TestSvGRPC_DingTalkMarkdown(t *testing.T) {
	//go RunServerGRPC(tsGRPCPort)
	time.Sleep(time.Microsecond * 10)
	testDingTalkMarkdown()
}

func testHealth() {
	conn, err := grpc.Dial(tsGRPCAddress, grpc.WithInsecure())
	if err != nil {
		log.Printf("did not connect: %v", err)
		return
	}
	defer conn.Close()
	c := pb2.NewNoticeServiceClient(conn)
	// Contact the server and print out its response.
	name := defaultName
	if len(os.Args) > 1 {
		name += os.Args[1]
	}
	ctx, cancel := context.WithTimeout(context.Background(), time.Second)
	defer cancel()
	r, err := c.Health(ctx, &pb2.HealthReq{ClientName: name})
	if err != nil {
		log.Printf("[ERROR]could not greet: %v", err)
	}
	log.Printf("testHealth: %+v \n ", r) //r.Message)
}

func testEmail() {
	conn, err := grpc.Dial(tsGRPCAddress, grpc.WithInsecure())
	if err != nil {
		log.Printf("did not connect: %v", err)
		return
	}
	defer conn.Close()
	c := pb2.NewNoticeServiceClient(conn)
	// Contact the server and print out its response.
	ctx, cancel := context.WithTimeout(context.Background(), time.Second)
	defer cancel()
	r, err := c.Email(ctx, &pb2.EmailReq{
		To:       testdata.NoticeEmailTo ,
		FromName: "xxxFrom",
		Subject:  "你好！",
		Type:     "html",
		Body:     `<h5 style="color: red; font-size: 16px;">测试-标题</h5>`,
	})
	if err != nil {
		log.Printf("[ERROR]could not greet: %v", err)
		return
	}
	log.Printf("testEmail: %+v \n ", r) //r.Message)
}

func testDingTalkMarkdown() {
	conn, err := grpc.Dial(tsGRPCAddress, grpc.WithInsecure())
	if err != nil {
		log.Printf("did not connect: %v", err)
		return
	}
	defer conn.Close()
	c := pb2.NewNoticeServiceClient(conn)
	// Contact the server and print out its response.
	name := defaultName
	if len(os.Args) > 1 {
		name += os.Args[1]
	}
	ctx, cancel := context.WithTimeout(context.Background(), time.Second)
	defer cancel()
	r, err := c.DingTalkMarkdown(ctx, &pb2.DingTalkReq{
		Token:       env.DingTalkConfig.RobotToken,
		RobotSecret: env.DingTalkConfig.RobotSecret,
		Title:       "title-标题-1",
		Text:        "text:这是一条golang钉钉消息测试",
		IsAtAll:     true,
	})
	if err != nil {
		log.Printf("[ERROR]could not greet: %v", err)
		return
	}
	log.Printf("dingTalkMarkdown: %+v \n ", r) //r.Message)
}
