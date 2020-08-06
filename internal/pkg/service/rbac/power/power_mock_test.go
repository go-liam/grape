package power

import (
	"github.com/go-liam/util/conv"
	"testing"
)

// 合并 各组权限
func TestPower_List(t *testing.T) {
	pid := 1
	group1 := []int{1, 2, 3, 4}
	group2 := []int{1, 2, 3, 4}
	join := conv.ArrayIntJoin(group1, group2)
	check := conv.ArrayIntContains(join, pid)
	println("用户权限：= ", check)
}

func TestSvPowerMock_List(t *testing.T) {
	ServerMock.List(1)
}

func TestSvPowerMock_Add(t *testing.T) {
	item := new(ModelPower)
	ServerMock.Add(item)
}

func TestSvPowerMock_Update(t *testing.T) {
	item := new(ModelPower)
	ServerMock.Edit(item)
}

func TestSvPowerMock_Delete(t *testing.T) {
	item := new(ModelPower)
	ServerMock.Delete(item)
}
