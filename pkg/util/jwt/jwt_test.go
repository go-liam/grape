package jwt

import (
	"log"
	"testing"
	"time"
)

func TestJWT_CreateToken(t *testing.T) {
	// creat
	j := NewJWT("")
	claims := &CustomClaims{
		UserID: 110,
		Name:   "中国",
	}
	minute := 10
	claims.NotBefore = time.Now().Unix() - 10000 // 签名生效时间
	claims.ExpiresAt = time.Now().Add(time.Minute * time.Duration(minute)).Unix()
	claims.Issuer = TokenIssuer //签名的发行者
	token, err := j.CreateToken(claims)
	if err == nil {
		log.Printf("token1=%+v\n",token)
		t.Log("PASS")
	} else {
		println(err)
		t.Fatal("Test_JWT creat jwt failure")
	}
	//check
	info, err2 := j.ParseToken(token)
	log.Printf("info=%+v\n", info)
	log.Printf("err2=%+v\n", err2)
	//refresh
	refresh,err3 :=  j.RefreshToken(token,123410)
	log.Printf("token2=%+v\n", refresh)
	log.Printf("err3=%+v\n", err3)
}
