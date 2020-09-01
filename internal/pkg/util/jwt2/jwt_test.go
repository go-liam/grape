package jwt2

import (
	"log"
	"testing"
	"time"
)

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

func TestName_Create(t *testing.T) {
	token, refresh := Create(123456789012345, 2, 0)
	println("token=", token)
	println("refresh=", refresh)
}

func TestName_Refresh(t *testing.T) {
	//t1 := "eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJ1aWQiOiIxMjM0NTY3ODkwMTIzNDUiLCJ1dHkiOjIsImV4cCI6MTU5ODkyNjA4NX0.RTjitjEoUAPllBU_U4uJmxmJyoVskn5F0d5VQ6xA25g"
	f1 := "eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJ1aWQiOiIxMjM0NTY3ODkwMTIzNDUiLCJ1dHkiOjIsImV4cCI6MTYwMTUxNzc4NX0.TBbgQawAJJT5GCkHV0wYG5R7ILcz7NC0G_lbzErjtz8"
	token, refresh := Refresh(f1)
	println("token=", token)
	println("refresh=", refresh)
}

func TestName_Refresh2(t *testing.T) {
	j := new(JWT)
	j.SetKey(SignKeyRefresh)
	claims := &CustomClaims{
		UserID:     1234567890123456,
		ClientType: 1,
	}
	minute := 10*24*3600 - 10
	claims.ExpiresAt = time.Now().Add(time.Second * time.Duration(minute)).Unix()
	f1, _ := j.CreateToken(claims)
	println("refresh old=", f1)
	token, refresh := Refresh(f1)
	println("token=", token)
	println("refresh=", refresh)
}
