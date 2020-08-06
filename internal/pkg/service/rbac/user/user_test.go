package user

import (
	"log"
	"testing"
)

func TestSvUserMock_List(t *testing.T) {
	list, _ := Server.List()
	for _, v := range list {
		log.Printf("v=%+v\n", v)
	}
}

func TestSvUserMock_Add(t *testing.T) {
	Server.Add(&ModelUser{UserID: 1, Name: "name1", RoleFlag: 1, RoleIDs: "[1]"})
}

func TestSvUserMock_Edit(t *testing.T) {
	Server.Edit(&ModelUser{UserID: 1, Name: "name1", RoleFlag: 1, RoleIDs: "[1]"})
}

func TestSvUserMock_Delete(t *testing.T) {
	Server.Delete(&ModelUser{UserID: 1, Name: "name1", RoleFlag: 1, RoleIDs: "[1]"})
}

func TestSvUserMock_UpdateRole(t *testing.T) {
	Server.UpdateRole(&ModelUser{UserID: 1, Name: "name1", RoleFlag: 1, RoleIDs: "[1]"})
}

func TestSvUserMock_FindOne(t *testing.T) {
	Server.FindOne(1)
}

func TestInfo(t *testing.T) {
	v, v2 := Info(2)
	log.Printf("v2=%+v\n", v2)
	log.Printf("v=%+v\n", v)
}
