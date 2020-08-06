package role

import "errors"

var ServerMock *SvRoleMock

func init() {
	ServerMock = new(SvRoleMock)
}

type SvRoleMock struct {
	list    []*ModelRole
	hadData bool
}

func (sv *SvRoleMock) dataInit() {
	if !sv.hadData {
		sv.hadData = true
		sv.list = make([]*ModelRole, 0)
		sv.list = append(sv.list, &ModelRole{RoleID: 1, RoleName: "name1", PowerIDS: "[1,2,3]"})
		sv.list = append(sv.list, &ModelRole{RoleID: 2, RoleName: "name1", PowerIDS: "[2,3]"})
		sv.list = append(sv.list, &ModelRole{RoleID: 3, RoleName: "name1", PowerIDS: "[4,5]"})
	}
}

func (sv *SvRoleMock) List() ([]*ModelRole, error) {
	sv.dataInit()
	return sv.list, nil
}

func (sv *SvRoleMock) Add(item *ModelRole) (int, error) {
	sv.dataInit()
	sv.list = append(sv.list, item)
	return 1, nil
}

func (sv *SvRoleMock) Edit(item *ModelRole) (int, error) {
	i := 0
	for _, v := range sv.list {
		if v.RoleID == item.RoleID {
			sv.list[i] = item
		}
		i++
	}
	return 1, nil
}
func (sv *SvRoleMock) UpdatePower(item *ModelRole) (int, error) {
	i := 0
	for _, v := range sv.list {
		if v.RoleID == item.RoleID {
			sv.list[i].PowerIDS = item.PowerIDS
		}
		i++
	}
	return 1, nil
}

func (sv *SvRoleMock) Delete(item *ModelRole) (int, error) {
	ls := make([]*ModelRole, 0)
	for _, v := range sv.list {
		if v.RoleID != item.RoleID {
			ls = append(ls, v)
		}
	}
	sv.list = ls
	return 1, nil
}

func (sv *SvRoleMock) FindOne(roleID int) (*ModelRole, error) {
	for _, v := range sv.list {
		if v.RoleID == roleID {
			return v, nil
		}
	}
	return nil, errors.New("No data")
}
