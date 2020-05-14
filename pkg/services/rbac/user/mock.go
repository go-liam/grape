package user

import "errors"

var ServerMock *SvUserMock

func init() {
	ServerMock = new(SvUserMock)
}

type SvUserMock struct {
	list    []*ModelUser
	hadData bool
}

func (sv *SvUserMock) dataInit() {
	if !sv.hadData {
		sv.hadData = true
		sv.list = make([]*ModelUser, 0)
		sv.list = append(sv.list, &ModelUser{UserID: 1, Name: "name1", RoleFlag: 1, RoleIDs: "[1]"})
		sv.list = append(sv.list, &ModelUser{UserID: 2, Name: "name2", RoleFlag: 0, RoleIDs: "[1,2]"})
		sv.list = append(sv.list, &ModelUser{UserID: 3, Name: "name3", RoleFlag: 0, RoleIDs: "[3]"})
	}
}

func (sv *SvUserMock) List() ([]*ModelUser, error) {
	sv.dataInit()
	return sv.list, nil
}

func (sv *SvUserMock) Add(item *ModelUser) (int, error) {
	sv.dataInit()
	sv.list = append(sv.list, item)
	return 1, nil
}

func (sv *SvUserMock) Edit(item *ModelUser) (int, error) {
	sv.dataInit()
	i := 0
	for _, v := range sv.list {
		if v.UserID == item.UserID {
			sv.list[i] = item
		}
		i++
	}
	return 1, nil
}

func (sv *SvUserMock) UpdateRole(item *ModelUser) (int, error) {
	sv.dataInit()
	i := 0
	for _, v := range sv.list {
		if v.UserID == item.UserID {
			sv.list[i].RoleIDs = item.RoleIDs
		}
		i++
	}
	return 1, nil
}

func (sv *SvUserMock) Delete(item *ModelUser) (int, error) {
	sv.dataInit()
	ls := make([]*ModelUser, 0)
	for _, v := range sv.list {
		if v.UserID != item.UserID {
			ls = append(ls, v)
		}
	}
	sv.list = ls
	return 1, nil
}

func (sv *SvUserMock) FindOne(userID int64) (*ModelUser, error) {
	sv.dataInit()
	for _, v := range sv.list {
		if v.UserID == userID {
			return v, nil
		}
	}
	return nil, errors.New("No data")
}
