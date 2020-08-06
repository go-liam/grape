package logclient

import (
	"context"
	"google.golang.org/grpc"
	pb2 "grape/api/proto/svlog"
	"grape/internal/pkg/config"
	"log"
	"time"
)

func init() {
	address := config.ServerSvLog.Host + config.ServerSvLog.PortGRPC
	ClientInit(address)
}

var client pb2.SvLogServiceClient
var isConnect bool

// address = localhost:9403
func ClientInit(address string) {
	conn, err := grpc.Dial(address, grpc.WithInsecure())
	if err != nil {
		log.Println("[ERROR]RPC Account Service ERROR:", err)
	}
	client = pb2.NewSvLogServiceClient(conn) //.NewAccountServiceClient(conn)
	isConnect = true
}

// ClientWrite :
func ClientWrite(level int32, logID int64, msg string) {
	if !isConnect {
		log.Printf("[ERROR] ClientWrite: 没有init数据")
		return
	}
	ctx, cancel := context.WithTimeout(context.Background(), time.Second)
	defer cancel()
	_, err := client.WriteLog(ctx, &pb2.WriteLogReq{LogID: logID, Level: level, Msg: msg})
	if err != nil {
		log.Printf("[ERROR] ClientWrite: %v", err)
		return
	}
}

// ClientWriteAsy :
func ClientWriteAsy(level int32, logID int64, msg string) {
	go ClientWrite(level, logID, msg)
}
