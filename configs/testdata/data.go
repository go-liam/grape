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
)