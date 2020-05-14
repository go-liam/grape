package core

import (
	"fmt"
	"github.com/go-liam/util/request"
	"log"
	"grape/pkg/util/conv"
	"grape/pkg/util/jwt"
	"net/http"
)

var badIpList = [...] string{"190.0.0.0"}

// 鉴权，黑名单,封IP
func auth(req *http.Request) bool {
	log.Printf("head:%+v\n", req.Header)
	setUserIDHead(req)
	ip := request.ClientIP(req)
	println("ip:", ip)
	return !isBadIP(ip)
}

func isBadIP(ip string) bool {
	for _,v := range  badIpList{
		if v == ip {
			return true
		}
	}
	return false
}

func setUserIDHead( req *http.Request )  {
	token := req.Header.Get("tokenjwt")
	if token !=""{
		info,_ := jwt.Server.ParseToken(token)
		if info != nil && conv.StringToInt64(info.UserID,0) > 0{
			req.Header.Set("UserID", fmt.Sprintf("%d", info.UserID))
		}
	}
}