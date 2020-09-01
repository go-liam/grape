package jwt2

import (
	jwtgo "github.com/dgrijalva/jwt-go"
	"time"
)

var (
	Success = 0
	// ErrorTokenExpired : 过期了
	ErrorTokenExpired = -1 // errors.New("Token is expired")
	// ErrorTokenNotValidYet ：还没有生效
	ErrorTokenNotValidYet = -2 // errors.New("Token not active yet")
	// ErrorTokenMalformed  ：That's not even a token
	ErrorTokenMalformed = -3 // errors.New("That's not even a token")
	// ErrorTokenInvalid ：Couldn't handle this token
	ErrorTokenInvalid = -4 //errors.New("Couldn't handle this token")
)

// JWT 签名结构
type JWT struct {
	SigningKey []byte
}

type CustomClaims struct {
	UserID     string `json:"uid"` // userID,18位长度，用int64 js取值会有问题
	ClientType int    `json:"uty"` // 客户端类型
	jwtgo.StandardClaims
}

// NewJWT 新建一个jwt实例
func NewJWT(key string) *JWT {
	return &JWT{
		[]byte(key),
	}
}

// CreateToken 生成一个token
func (j *JWT) CreateToken(claims *CustomClaims) (string, error) {
	token := jwtgo.NewWithClaims(jwtgo.SigningMethodHS256, claims)
	return token.SignedString(j.SigningKey)
}

// ParseToken 解析Tokne
func (j *JWT) ParseToken(tokenString string) (*CustomClaims, int) {
	token, err := jwtgo.ParseWithClaims(tokenString, &CustomClaims{}, func(token *jwtgo.Token) (interface{}, error) {
		return j.SigningKey, nil
	})
	if err != nil {
		if ve, ok := err.(*jwtgo.ValidationError); ok {
			if ve.Errors&jwtgo.ValidationErrorMalformed != 0 {
				return nil, ErrorTokenMalformed
			} else if ve.Errors&jwtgo.ValidationErrorExpired != 0 {
				// Token is expired
				return nil, ErrorTokenExpired
			} else if ve.Errors&jwtgo.ValidationErrorNotValidYet != 0 {
				return nil, ErrorTokenNotValidYet
			} else {
				return nil, ErrorTokenInvalid
			}
		}
	}
	if claims, ok := token.Claims.(*CustomClaims); ok && token.Valid {
		return claims, Success
	}
	return nil, ErrorTokenInvalid
}

// RefreshToken 更新token,认加 x
func (j *JWT) RefreshToken(tokenString string, second int32) (string, int) {
	info, err2 := j.ParseToken(tokenString)
	if err2 != Success {
		return "", err2
	}
	info.ExpiresAt = time.Now().Add(time.Second * time.Duration(second)).Unix()
	v, _ := j.CreateToken(info)
	return v, Success
}
