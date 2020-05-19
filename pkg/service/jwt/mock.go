package jwt

import (
	"errors"
)

/*
模拟redis 存储数据
*/

type SvMock struct {
	list []*ModelUserToken
}

var ServerMock = new(SvMock)

func init() {
	ServerMock.list = make([]*ModelUserToken, 0)
}

func (sv *SvMock) UpdateOrAdd(item *ModelUserToken) (bool, error) {
	sv.DeleteByUserIDAndClient(item.UserID,item.ClientID)
	//for _, v := range sv.list {
	//	if v.UserID == item.UserID && item.ClientID == v.ClientID {
	//		v= item
	//		return true, nil
	//	}
	//}
	sv.list = append(sv.list, item)
	return true, nil
}

func (sv *SvMock) GET(userID int64, clientID int) (*ModelUserToken, error) {
	for _, v := range sv.list {
		if v.UserID == userID && clientID == v.ClientID {
			return v, nil
		}
	}
	return nil, errors.New("No data ")
}

func (sv *SvMock) DeleteAllByUserID(userID int64) (bool, error) {
	ls := make([]*ModelUserToken, 0)
	for _, v := range sv.list {
		if v.UserID != userID {
			ls = append(ls, v)
		}
	}
	sv.list = ls
	return true, nil
}

func (sv *SvMock) DeleteByUserIDAndClient(userID int64,clientID int) (bool, error) {
	ls := make([]*ModelUserToken, 0)
	for _, v := range sv.list {
		if v.UserID == userID && v.ClientID == clientID {
			continue
		}
		ls = append(ls, v)
	}
	sv.list = ls
	return true, nil
}