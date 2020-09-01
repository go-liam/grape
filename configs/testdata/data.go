package testdata

import "errors"

const (
	ErrDbErrResult   = "数据库错误"
	ErrDbUnfulfilled = "there were unfulfilled expectations:"

	ConstSuccess  = "ok"
	ConstFail     = "fail"

	ConstWant0    = 0
	ConstWantOne  = 1
	ConstWantTwo = 2
	ConstWantBigInt = 1597027173
	ConstSecondTime = 1597027175

	ConstWantString = "tsName"
	ConstWantString2 = "Name-nul"
)

var ErrorDBConnect = errors.New("数据库错误")

const (
	UserKey = "jsrseida"
	AdminKey = "box2CELUbG1rVYcbzO32R7WEj0oE9ECu"
	JWTKey = "eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJ1aWQiOjEyMzQ1Njc4OTAxMjM0NTY3LCJ1dHkiOjEsImxnIjoxLCJleHAiOjY3ODI5NDMxMTJ9.AExdsA3KKA9VTiNPimDnM6PELPSeEG27lLWzKSpqUIc"
)