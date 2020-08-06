package svgrpc

import (
	"context"
	"google.golang.org/grpc"
	pb2 "grape/api/proto/notice"
	"grape/internal/pkg/config"
	"log"
	"time"
)

var (
	svNoticeClient        pb2.NoticeServiceClient
	noticeClientIsConnect = false
)

func initNotice() (int, error) {
	if noticeClientIsConnect {
		return 2, nil
	}
	address := config.ServerSvNotice.Host + config.ServerSvNotice.PortGRPC
	conn, err := grpc.Dial(address, grpc.WithInsecure())
	if err != nil {
		log.Printf("did not connect: %v", err)
		return -2, err
	}
	//defer conn.Close()
	svNoticeClient = pb2.NewNoticeServiceClient(conn)
	noticeClientIsConnect = true
	return 1, nil
}

func NoticeEmail(info *pb2.EmailReq) (int, error) {
	rs, err1 := initNotice()
	if rs <= 0 {
		return rs, err1
	}
	// Contact the server and print out its response.
	ctx, cancel := context.WithTimeout(context.Background(), time.Second)
	defer cancel()
	r, err := svNoticeClient.Email(ctx, info)
	if err != nil {
		log.Printf("[ERROR]could not greet: %v", err)
		return -1, err
	}
	log.Printf("testEmail: %+v \n ", r) //r.Message)
	return 1, nil
}

func NoticeDingTalkMarkdown(info *pb2.DingTalkReq) (int, error) {
	rs, err1 := initNotice()
	if rs <= 0 {
		return rs, err1
	}
	// Contact the server and print out its response.
	ctx, cancel := context.WithTimeout(context.Background(), time.Second)
	defer cancel()
	r, err := svNoticeClient.DingTalkMarkdown(ctx, info)
	if err != nil {
		log.Printf("[ERROR]could not greet: %v", err)
		return -2, err
	}
	log.Printf("dingTalkMarkdown: %+v \n ", r) //r.Message)
	return 1, nil
}
