package jwt2

import (
	"grape/internal/pkg/util/jwt2"
	"testing"
	"time"
)

func TestName_Create(t *testing.T) {
	token, refresh := Create(123456789012345, 2, 0)
	println("token=", token)
	println("refresh=", refresh)
}

func TestName_Refresh(t *testing.T) {
	f1 := "eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJ1aWQiOiIxMjM0NTY3ODkwMTIzNDUiLCJ1dHkiOjIsImV4cCI6MTYwMTUxNzc4NX0." +
		"TBbgQawAJJT5GCkHV0wYG5R7ILcz7NC0G_lbzErjtz8"
	token, refresh := Refresh(f1)
	println("token=", token)
	println("refresh=", refresh)
}

func TestName_Refresh2(t *testing.T) {
	j := new(jwt2.JWT)
	j.SetKey(SignKeyRefresh)
	claims := &jwt2.CustomClaims{
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

func TestName_token(t *testing.T) {
	j := new(jwt2.JWT)
	j.SetKey(SignKey)
	claims := &jwt2.CustomClaims{
		UserID:     1,
		ClientType: 1,
	}
	minute := 300 * 24 * 3600
	claims.ExpiresAt = time.Now().Add(time.Second * time.Duration(minute)).Unix()
	f1, _ := j.CreateToken(claims)
	println("token =", f1)
}
