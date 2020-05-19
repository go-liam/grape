package server

import (
	"context"
	"google.golang.org/grpc"
	pb2 "grape/proto/notice"
	"log"
	"os"
	"testing"
	"time"
)

const (
	tsGRPCPort =":4401"
	tsGRPCAddress     = "localhost"+tsGRPCPort
	defaultName = "client:"
)


func TestSvGRPC_Health(t *testing.T) {
	go RunServerGRPC(tsGRPCPort)
	time.Sleep(time.Microsecond *10)
	testHealth()
}

func TestSvGRPC_Email(t *testing.T) {
	go RunServerGRPC(tsGRPCPort)
	time.Sleep(time.Microsecond *10)
	testEmail()
}

func TestSvGRPC_DingTalkMarkdown(t *testing.T) {
	go RunServerGRPC(tsGRPCPort)
	time.Sleep(time.Microsecond *10)
	testDingTalkMarkdown()
}

func testHealth() {
	conn, err := grpc.Dial(tsGRPCAddress, grpc.WithInsecure())
	if err != nil {
		log.Fatalf("did not connect: %v", err)
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
		log.Fatalf("did not connect: %v", err)
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
	r, err := c.Email(ctx, &pb2.EmailReq{
		To:                   "56487685@qq.com",
		FromName:             "xxxFrom",
		Subject:              "你好！",
		Type:                 "html",
		Body:                 `<h5 style="color: red; font-size: 16px;">测试-标题</h5>`,
	})
	if err != nil {
		log.Printf("[ERROR]could not greet: %v", err)
	}
	log.Printf("testEmail: %+v \n ", r) //r.Message)
}

func testDingTalkMarkdown() {
	conn, err := grpc.Dial(tsGRPCAddress, grpc.WithInsecure())
	if err != nil {
		log.Fatalf("did not connect: %v", err)
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
		Token:                "xxx1",
		RobotSecret:          "xxx2",
		Title:                "title-标题",
		Text:                 "text-",
		IsAtAll:              false,
	})
	if err != nil {
		log.Printf("[ERROR]could not greet: %v", err)
	}
	log.Printf("dingTalkMarkdown: %+v \n ", r) //r.Message)
}
