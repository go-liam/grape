package server

import (
	"context"
	"google.golang.org/grpc"
	"grape/pkg/config"
	pb2 "grape/proto/svlog"
	"log"
	"os"
	"testing"
	"time"
)

var portGRPCTest = config.ServerSvLog.PortGRPC // ":9403"

//func TestMain(m *testing.M) {
//	println("test- TestMain")
//}

func TestSvGRPC_WriteLog(t *testing.T) {
	testWriteLog()
}

func TestSvGRPC_Health(t *testing.T) {
	testHealth()
}

func testWriteLog() {
	conn, err := grpc.Dial("localhost"+PortGRPC, grpc.WithInsecure())
	if err != nil {
		log.Printf("did not connect: %v", err)
		return
	}
	defer conn.Close()
	c := pb2.NewSvLogServiceClient(conn)
	// Contact the server and print out its response.
	ctx, cancel := context.WithTimeout(context.Background(), time.Second)
	defer cancel()
	r, err := c.WriteLog(ctx, &pb2.WriteLogReq{LogID: 123, Level: 3, Msg: "msg--xxxx"})
	if err != nil {
		log.Printf("could not greet: %v", err)
		return
	}
	log.Printf("testWriteLog-Greeting: %+v \n ", r) //r.Message)
}

func testHealth() {
	println("test-health-")
	conn, err := grpc.Dial("localhost"+PortGRPC, grpc.WithInsecure())
	if err != nil {
		log.Printf("did not connect: %v", err)
		return
	}
	defer conn.Close()
	c := pb2.NewSvLogServiceClient(conn)

	// Contact the server and print out its response.
	name := serverName
	if len(os.Args) > 1 {
		name = os.Args[1]
	}
	ctx, cancel := context.WithTimeout(context.Background(), time.Second)
	defer cancel()
	r, err := c.Health(ctx, &pb2.HealthReq{ClientName: name})
	if err != nil {
		log.Printf("could not greet: %v", err)
		return
	}
	log.Printf("Greeting: %+v \n ", r) //r.Message)
}
