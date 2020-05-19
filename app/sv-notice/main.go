//go:generate protoc -I ../helloworld --go_out=plugins=grpc:../helloworld ../helloworld/helloworld.proto

package main

import (
	"log"
	"grape/app/sv-notice/server"
	"os"
	"os/signal"
)

const (
	portGRPC = ":9405"
	partGin    = ":7405"
)

func main() {
	println(":run server")
	go server.RunServerGin(partGin)
	server.RunServerGRPC(portGRPC)
	quit := make(chan os.Signal)
	signal.Notify(quit, os.Interrupt)
	<-quit
	log.Println("Shutdown Server ...")
}

//
//const (
//	port = ":7405"
//)
//
//// server is used to implement helloworld.GreeterServer.
//type server struct{}
//
//// SayHello implements helloworld.GreeterServer
//func (s *server) SayHello(ctx context.Context, in *pb.HelloRequest) (*pb.HelloReply, error) {
//	return &pb.HelloReply{Message: "Hello " + in.Name}, nil
//}
//
//func main() {
//	println("server run on:", port)
//	lis, err := net.Listen("tcp", port)
//	if err != nil {
//		log.Fatalf("failed to listen: %v", err)
//	}
//	s := grpc.NewServer()
//	pb.RegisterGreeterServer(s, &server{})
//	// Register reflection service on gRPC server.
//	reflection.Register(s)
//	if err := s.Serve(lis); err != nil {
//		log.Fatalf("failed to serve: %v", err)
//	}
//}
