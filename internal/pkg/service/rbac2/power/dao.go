package power

import (
	"grape/internal/pkg/database/mysql"
	"time"
)

//var Server2 = new(SrvPower)

type SrvPower struct {
}

func (e *SrvPower) Create(item *Model) (int64, error) {
	item.CreatedAt = time.Now().Unix()
	item.UpdatedAt = item.CreatedAt
	v := mysql.ServerAPI.Engine().Create(item)
	return v.RowsAffected, v.Error
}

func (e *SrvPower) FindOne(id int64) (*Model, error) {
	item := new(Model)
	sql := "select * from rb_power where id =?  and status < 44 limit 1 "
	v := mysql.ServerAPI.Engine().Raw(sql, id).Scan(item)
	return item, v.Error
}

func (e *SrvPower) FindMulti() ([]*Model, error) {
	var result []*Model
	sql := "select * from rb_power where status < 44 "
	v := mysql.ServerAPI.Engine().Raw(sql).Scan(&result)
	return result, v.Error
}

func (e *SrvPower) Update(item *Model) (int64, error) {
	item.UpdatedAt = time.Now().Unix()
	sql := "update rb_power set `name` = ?,extended=?,updated_at=?,tag=? where `id` = ? "
	v := mysql.ServerAPI.Engine().Exec(sql, item.Name, item.Extended, item.UpdatedAt, item.Tag, item.ID)
	return v.RowsAffected, v.Error
}

func (e *SrvPower) UpdateStatus(item *Model) (int64, error) {
	item.UpdatedAt = time.Now().Unix()
	sql := "update rb_power set `status` = ?,updated_at=? where `id` = ? "
	v := mysql.ServerAPI.Engine().Exec(sql, item.Status, item.UpdatedAt, item.ID)
	return v.RowsAffected, v.Error
}
