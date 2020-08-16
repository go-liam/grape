package bind

import (
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
