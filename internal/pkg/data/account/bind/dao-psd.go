package bind

import (
	"errors"
	"grape/internal/pkg/database/mysql"
	"time"
)

var ServerPsd = new(SrvPsd)

type SrvPsd struct {
}

func (e *SrvPsd) Create(item *ModelPassword) (int64, error) {
	item.CreatedAt = time.Now().Unix()
	item.UpdatedAt = item.CreatedAt
	v := mysql.ServerAPI.Engine().Create(item)
	return v.RowsAffected, v.Error
}

func (e *SrvPsd) FindOne(userID int64) (*ModelPassword, error) {
	item := new(ModelPassword)
	sql := "select * from uc_user_psd where id = ? limit 1 "
	v := mysql.ServerAPI.Engine().Raw(sql, userID).Scan(item)
	return item, v.Error
}

func (e *SrvPsd) CheckAndCreate(item *ModelPassword) (int64, error) {
	ck, _ := e.FindOne(item.ID)
	if ck.ID > 0 {
		return 0, errors.New("had data")
	}
	return e.Create(item)
}

func (e *SrvPsd) Update(item *ModelPassword) (int64, error) {
	item.UpdatedAt = time.Now().Unix()
	sql := "update uc_user_psd set password=?,updated_at=?  where `id` = ? "
	v := mysql.ServerAPI.Engine().Exec(sql, item.Password, item.UpdatedAt, item.ID)
	return v.RowsAffected, v.Error
}
