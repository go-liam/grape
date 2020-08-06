package server

import (
	"context"
	"google.golang.org/grpc"
	"google.golang.org/grpc/reflection"
	pb "grape/api/proto/svlog"
	"grape/pkg/service/sylog"
	"log"
	"net"
)

const (
	PortGRPC   = ":9403"
	serverName = "sv-log"
)

// svGRPC is used to implement helloworld.GreeterServer.
type svGRPC struct{}

func (s *svGRPC) Health(ctx context.Context, in *pb.HealthReq) (*pb.HealthResp, error) {
	mess := in.ClientName
	return &pb.HealthResp{Message: mess, Code: 0, Version: 123456789012}, nil
}

func (s *svGRPC) WriteLog(ctx context.Context, in *pb.WriteLogReq) (*pb.WriteLogResp, error) {
	log.Printf("info %+v\n", in)
	go svlog.WriteLog(in.LogID, in.Msg, in.Level)
	return &pb.WriteLogResp{Code: 0, Message: "OK"}, nil
}

func ServerGRPC(port string) {
	println("svGRPC run on:", port)
	lis, err := net.Listen("tcp", port)
	if err != nil {
		log.Printf("failed to listen: %v \n", err)
	}
	s := grpc.NewServer()
	pb.RegisterSvLogServiceServer(s, &svGRPC{})
	// Register reflection service on gRPC svGRPC.
	reflection.Register(s)
	if err := s.Serve(lis); err != nil {
		log.Printf("failed to serve: %v \n", err)
	}
}
