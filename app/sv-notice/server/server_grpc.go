package server

import (
	"context"
	"google.golang.org/grpc"
	"google.golang.org/grpc/reflection"
	pb "grape/proto/configproto"
	"log"
	"net"
)

const (
	portGRPC = ":9405"
)

// svGRPC is used to implement helloworld.GreeterServer.
type svGRPC struct{}

func (s*svGRPC) Health(ctx context.Context, in *pb.HealthReq ) (*pb.HealthResp,error) {
	mess:= in.ClientName
	return &pb.HealthResp{ Message: mess ,Code: 0 ,Version: 123456789012},nil
}

func (s*svGRPC) FindOne(ctx context.Context, in *pb.FindOneReq ) (*pb.FindOneResp,error) {
	return &pb.FindOneResp{ },nil
}

func (s*svGRPC) Update(ctx context.Context, in *pb.UpdateReq ) (*pb.UpdateResp,error) {
	return &pb.UpdateResp{ },nil
}

func (s*svGRPC) List(ctx context.Context, in *pb.ListReq ) (*pb.ListResp,error) {
	return &pb.ListResp{ },nil
}

// RunServerGRPC :
func RunServerGRPC() {
	println("svGRPC run on:", portGRPC)
	lis, err := net.Listen("tcp", portGRPC)
	if err != nil {
		log.Printf("failed to listen: %v \n", err)
	}
	s := grpc.NewServer()
	pb.RegisterConfigServiceServer(s, &svGRPC{})
	// Register reflection service on gRPC svGRPC.
	reflection.Register(s)
	if err := s.Serve(lis); err != nil {
		log.Printf("failed to serve: %v \n", err)
	}
}
