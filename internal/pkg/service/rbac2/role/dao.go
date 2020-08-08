package role

import (
	"grape/internal/pkg/database/mysql"
	"time"
)

//var Server2 = new(SrvRole)

type SrvRole struct {
}

func (e *SrvRole) Create(item *Model) (int64, error) {
	item.CreatedAt = time.Now().Unix()
	item.UpdatedAt = item.CreatedAt
	v := mysql.ServerAPI.Engine().Create(item)
	return v.RowsAffected, v.Error
}

func (e *SrvRole) FindOne(id int64) (*Model, error) {
	item := new(Model)
	sql := "select * from rb_role where id =?  and status < 44 limit 1 "
	v := mysql.ServerAPI.Engine().Raw(sql, id).Scan(item)
	return item, v.Error
}

func (e *SrvRole) FindMulti() ([]*Model, error) {
	var result []*Model
	sql := "select * from rb_role where status < 44 "
	v := mysql.ServerAPI.Engine().Raw(sql).Scan(&result)
	return result, v.Error
}

func (e *SrvRole) Update(item *Model) (int64, error) {
	item.UpdatedAt = time.Now().Unix()
	sql := "update rb_role set `name` = ?,extended=?,updated_at=?,power_ids=? where `id` = ? "
	v := mysql.ServerAPI.Engine().Exec(sql, item.Name, item.Extended, item.UpdatedAt, item.PowerIDS, item.ID)
	return v.RowsAffected, v.Error
}

func (e *SrvRole) UpdateStatus(item *Model) (int64, error) {
	item.UpdatedAt = time.Now().Unix()
	sql := "update rb_role set `status` = ?,updated_at=? where `id` = ? "
	v := mysql.ServerAPI.Engine().Exec(sql, item.Status, item.UpdatedAt, item.ID)
	return v.RowsAffected, v.Error
}
