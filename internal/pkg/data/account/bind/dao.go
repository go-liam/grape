package bind

import (
	"errors"
	"grape/internal/pkg/database/mysql"
	"time"
)

var Server = new(SrvBind)

type SrvBind struct {
}

func (e *SrvBind) Create(item *Model) (int64, error) {
	item.CreatedAt = time.Now().Unix()
	v := mysql.ServerAPI.Engine().Create(item)
	return v.RowsAffected, v.Error
}

func (e *SrvBind) FindOne(ty int8, name string) (*Model, error) {
	item := new(Model)
	sql := "select * from uc_user_bind where type =? and name = ? limit 1 "
	v := mysql.ServerAPI.Engine().Raw(sql, ty, name).Scan(item)
	return item, v.Error
}

func (e *SrvBind) FindOneByUserID(ty int8, userID int64) (*Model, error) {
	item := new(Model)
	sql := "select * from uc_user_bind where type =? and user_id = ? limit 1 "
	v := mysql.ServerAPI.Engine().Raw(sql, ty, userID).Scan(item)
	return item, v.Error
}

func (e *SrvBind) CheckAndCreate(item *Model) (int64, error) {
	ck, _ := e.FindOne(item.Type, item.Name)
	if ck.ID > 0 {
		return 0, errors.New("had data")
	}
	return e.Create(item)
}

func (e *SrvBind) CheckAndCreateByUserID(item *Model) (int64, error) {
	ck, _ := e.FindOneByUserID(item.Type, item.UserID)
	if ck.ID > 0 {
		return 0, errors.New("had data")
	}
	return e.Create(item)
}
