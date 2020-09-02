package account

import (
	"log"
	"testing"
)

func TestName_RegByNamePassword(t *testing.T) {
	i := &RegModel{
		NickName: "管理员",
		Password: "123456",
		RegionID: 1,
		UserName: "root1",
		Phone:    "",
		Email:    "",
		Remark:   "公司员工",
		Extended: "", //`{"x":1,"b":"xxx"}`
	}
	v1, v2 := RegByNamePassword(i)
	println("v1=", v1)
	log.Printf("v2=%+v\n", v2)
}
