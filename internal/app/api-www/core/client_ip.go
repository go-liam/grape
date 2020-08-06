package core

import (
	"github.com/gin-gonic/gin"
	"log"
	"net"
	"net/http"
	"strings"
)

type ModelIP struct {
	IP string
	Request interface{}
	Header interface{}
}

func ClientIPHeader(c *gin.Context) {
	i :=new(ModelIP)
	i.IP =  ClientIP(c.Request)
	//i.Request = c.Request
	//log.Println(strings.Split(c.r.LocalAddr().String(), ":")[0])
	//i.ServerIP = strings.Split(c.Request.LocalAddr().String(), ":")[0])
	i.Header = c.Request.Header
	c.JSON(http.StatusOK, i )
}

// ClientIP 尽最大努力实现获取客户端 IP 的算法。
// 解析 X-Real-IP 和 X-Forwarded-For 以便于反向代理（nginx 或 haproxy）可以正常工作。
func ClientIP(r *http.Request) string {
	log.Printf("header: %+v\n",r )
	xForwardedFor := r.Header.Get("X-Forwarded-For")
	ip := strings.TrimSpace(strings.Split(xForwardedFor, ",")[0])
	if ip != "" {
		return ip
	}
	ip = strings.TrimSpace(r.Header.Get("X-Real-Ip"))
	if ip != "" {
		return ip
	}
	ip = strings.TrimSpace(r.Header.Get("X-Forwarded-Host"))
	if ip != "" {
		return ip
	}
	ip = strings.TrimSpace(r.Header.Get("X-Test-Host"))
	if ip != "" {
		return ip
	}
	if ip, _, err := net.SplitHostPort(strings.TrimSpace(r.RemoteAddr)); err == nil {
		return ip
	}
	return ""
}
