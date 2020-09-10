package jwt2

import (
	"log"
	"testing"
	"time"
)

const SignKey  = "xxxxxxx"

func TestJWT_CreateToken(t *testing.T) {
	// creat
	j := new(JWT)
	j.SetKey(SignKey)
	claims := &CustomClaims{
		UserID:     12345678901234567,
		ClientType: 1,
		LoginFlag:  1,
	}
	minute := 1000 * 24 * 3600
	claims.ExpiresAt = time.Now().Add(time.Minute * time.Duration(minute)).Unix()
	token, err := j.CreateToken(claims)
	if err == nil {
		log.Printf("token1=%+v\n", token)
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
	refresh, err3 := j.RefreshToken(token, 123410)
	log.Printf("token2=%+v\n", refresh)
	log.Printf("err3=%+v\n", err3)
}

func TestName_ParseToken(t *testing.T) {
	j := new(JWT)
	j.SetKey(SignKey)
	info, err2 := j.ParseToken("token")
	log.Printf("info=%+v\n", info)
	log.Printf("err2=%+v\n", err2)
}
