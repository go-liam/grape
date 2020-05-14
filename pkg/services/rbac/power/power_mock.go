package power

import "errors"

var ServerMock *SvPowerMock

func init() {
	ServerMock = new(SvPowerMock)
}

// SvPowerMock :
type SvPowerMock struct {
	list []*ModelPower
	hadData bool
}

func (sv *SvPowerMock) dataInit()  {
	if !sv.hadData {
		sv.hadData = true
		sv.list = make([]*ModelPower, 0)
		sv.list = append(sv.list, &ModelPower{PowerID: 1, PowerTag: "box-add"})
		sv.list = append(sv.list, &ModelPower{PowerID: 2, PowerTag: "box-edit"})
		sv.list = append(sv.list, &ModelPower{PowerID: 3, PowerTag: "box-delete"})
		sv.list = append(sv.list, &ModelPower{PowerID: 4, PowerTag: "box-update"})
	}
}
func (sv *SvPowerMock) List(powerType int8) ([]*ModelPower, error) {
	if powerType != 1 {
		return nil, nil
	}
	sv.dataInit()
	return sv.list, nil
}

func (sv *SvPowerMock) Add(item *ModelPower) (int, error) {
	sv.dataInit()
	sv.list = append(sv.list, item)
	return 1, nil
}

func (sv *SvPowerMock) Delete(item *ModelPower) (int, error) {
	sv.dataInit()
	ls := make([]*ModelPower, 0)
	for _, v := range sv.list {
		if v.PowerID != item.PowerID {
			ls = append(ls, v)
		}
	}
	sv.list = ls
	return 1, nil
}

func (sv *SvPowerMock) Edit(item *ModelPower) (int, error) {
	sv.dataInit()
	i := 0
	for _, v := range sv.list {
		if v.PowerID == item.PowerID {
			sv.list[i] = item
		}
		i++
	}
	return 1, nil
}


func (sv *SvPowerMock) FindOne(powerID int) (*ModelPower, error) {
	sv.dataInit()
	for _, v := range sv.list {
		if v.PowerID == powerID {
			return v, nil
		}
	}
	return nil, errors.New("No data")
}
