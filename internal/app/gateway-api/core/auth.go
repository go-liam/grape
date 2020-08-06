package core

import (
	"github.com/go-liam/util/conv"
	"github.com/go-liam/util/request"
	"grape/internal/pkg/util/jwt"
	"log"
	"net/http"
)

var badIpList = [...]string{"190.0.0.0"}

// 鉴权，黑名单,封IP
func auth(req *http.Request) bool {
	log.Printf("head:%+v\n", req.Header)
	setUserIDHead(req)
	ip := request.ClientIP(req)
	println("ip:", ip)
	return !isBadIP(ip)
}

func isBadIP(ip string) bool {
	for _, v := range badIpList {
		if v == ip {
			return true
		}
	}
	return false
}

func setUserIDHead(req *http.Request) {
	token := req.Header.Get("tokenjwt")
	if token != "" {
		info, _ := jwt.Server.ParseToken(token)
		if info != nil && conv.StringToInt64(info.UserID, 0) > 0 {
			req.Header.Set("UserID", info.UserID)
		}
	}
}
