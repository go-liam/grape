package logclient

import (
	"grape/configs"
	"testing"
	"time"
)

func TestName_WriteLog(t *testing.T) {
	ClientInit("localhost" + configs.ServerSvLog.PortGRPC)
	ClientWrite(1, 12345, "你好啊！12232")
}

func TestWriteLogAsy(t *testing.T) {
	ClientInit("localhost" + configs.ServerSvLog.PortGRPC)
	ClientWriteAsy(2, 12345, "你好啊！12232")
	time.Sleep(1 * time.Second)
}
