package role

import "errors"

var ServerMock *SvRoleMock

func init() {
	ServerMock = new(SvRoleMock)
}

type SvRoleMock struct {
	list    []*Model
	hadData bool
}

func (sv *SvRoleMock) dataInit() {
	if !sv.hadData {
		sv.hadData = true
		sv.list = make([]*Model, 0)
		sv.list = append(sv.list, &Model{ID: 1, Name: "name1", PowerIDS: "[1,2,3]"})
		sv.list = append(sv.list, &Model{ID: 2, Name: "name1", PowerIDS: "[2,3]"})
		sv.list = append(sv.list, &Model{ID: 3, Name: "name1", PowerIDS: "[4,5]"})
	}
}

func (sv *SvRoleMock) List() ([]*Model, error) {
	sv.dataInit()
	return sv.list, nil
}

func (sv *SvRoleMock) Add(item *Model) (int, error) {
	sv.dataInit()
	sv.list = append(sv.list, item)
	return 1, nil
}

func (sv *SvRoleMock) Edit(item *Model) (int, error) {
	i := 0
	for _, v := range sv.list {
		if v.ID == item.ID {
			sv.list[i] = item
		}
		i++
	}
	return 1, nil
}
func (sv *SvRoleMock) UpdatePower(item *Model) (int, error) {
	i := 0
	for _, v := range sv.list {
		if v.ID == item.ID {
			sv.list[i].PowerIDS = item.PowerIDS
		}
		i++
	}
	return 1, nil
}

func (sv *SvRoleMock) Delete(item *Model) (int, error) {
	ls := make([]*Model, 0)
	for _, v := range sv.list {
		if v.ID != item.ID {
			ls = append(ls, v)
		}
	}
	sv.list = ls
	return 1, nil
}

func (sv *SvRoleMock) FindOne(roleID int64) (*Model, error) {
	for _, v := range sv.list {
		if v.ID == roleID {
			return v, nil
		}
	}
	return nil, errors.New("No data")
}
