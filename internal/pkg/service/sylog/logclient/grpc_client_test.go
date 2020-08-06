package logclient

import (
	"grape/internal/pkg/config"
	"testing"
	"time"
)

func TestName_WriteLog(t *testing.T) {
	ClientInit("localhost" + config.ServerSvLog.PortGRPC)
	ClientWrite(1, 12345, "你好啊！12232")
}

func TestWriteLogAsy(t *testing.T) {
	ClientInit("localhost" + config.ServerSvLog.PortGRPC)
	ClientWriteAsy(2, 12345, "你好啊！12232")
	time.Sleep(1 * time.Second)
}
