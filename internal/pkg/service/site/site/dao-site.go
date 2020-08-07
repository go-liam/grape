package site

import (
	"grape/internal/pkg/database/mysql"
	"time"
)

var Server = new(SrvSite)

type SrvSite struct {
}

func (e *SrvSite) Create(item *Model) (int64, error) {
	item.CreatedAt = time.Now().Unix()
	item.UpdatedAt = item.CreatedAt
	v := mysql.ServerAPI.Engine().Create(item)
	return v.RowsAffected, v.Error
}

func (e *SrvSite) FindOne(id int64) (*Model, error) {
	item := new(Model)
	sql := "select * from ws_site where id =?  and status < 44 limit 1 "
	v := mysql.ServerAPI.Engine().Raw(sql, id).Scan(item)
	return item, v.Error
}

func (e *SrvSite) FindMulti() ([]*Model, error) {
	var result []*Model
	sql := "select * from ws_site where status < 44 "
	v := mysql.ServerAPI.Engine().Raw(sql).Scan(&result)
	return result, v.Error
}

func (e *SrvSite) Update(item *Model) (int64, error) {
	item.UpdatedAt = time.Now().Unix()
	sql := "update ws_site set `title` = ?,email=?,description=?,extended=?,updated_at=? where `id` = ? "
	v := mysql.ServerAPI.Engine().Exec(sql, item.Title, item.Email, item.Description, item.Extended, item.UpdatedAt, item.ID)
	return v.RowsAffected, v.Error
}

func (e *SrvSite) UpdateState(item *Model) (int64, error) {
	item.UpdatedAt = time.Now().Unix()
	sql := "update ws_site set `status` = ?,updated_at=? where `id` = ? "
	v := mysql.ServerAPI.Engine().Exec(sql, item.Status, item.UpdatedAt, item.ID)
	return v.RowsAffected, v.Error
}
