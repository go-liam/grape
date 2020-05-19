package server

import (
	"context"
	"google.golang.org/grpc"
	"google.golang.org/grpc/reflection"
	pb "grape/proto/notice"
	"log"
	"net"
)

// svGRPC is used to implement helloworld.GreeterServer.
type svGRPC struct{}

func (s*svGRPC) Health(ctx context.Context, in *pb.HealthReq ) (*pb.HealthResp,error) {
	mess:= in.ClientName
	return &pb.HealthResp{ Message: mess ,Code: 0 ,Version: "123456789012"},nil
}

func (s*svGRPC) Email(ctx context.Context, in *pb.EmailReq ) (*pb.EmailResp,error) {

	return &pb.EmailResp{ Code: 0,Message: in.Subject},nil
}

func (s*svGRPC) DingTalkMarkdown(ctx context.Context, in *pb.DingTalkReq ) (*pb.DingTalkResp,error) {
	return &pb.DingTalkResp{
		Code:                 0,
		Message:              "OK",
	},nil
}

// RunServerGRPC :
func RunServerGRPC(port string) {
	println("svGRPC run on:", port)
	lis, err := net.Listen("tcp", port)
	if err != nil {
		log.Printf("failed to listen: %v \n", err)
	}
	s := grpc.NewServer()
	pb.RegisterNoticeServiceServer(s, &svGRPC{})
	// Register reflection service on gRPC svGRPC.
	reflection.Register(s)
	if err := s.Serve(lis); err != nil {
		log.Printf("failed to serve: %v \n", err)
	}
}