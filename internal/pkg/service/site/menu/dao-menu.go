package menu

import (
	"grape/internal/pkg/database/mysql"
	"time"
)

var Server = new(SrvMenu)

type SrvMenu struct {
}

func (e *SrvMenu) Create(item *Model) (int64, error) {
	item.CreatedAt = time.Now().Unix()
	item.UpdatedAt = item.CreatedAt
	v := mysql.ServerAPI.Engine().Create(item)
	return v.RowsAffected, v.Error
}

func (e *SrvMenu) FindOne(id int64) (*Model, error) {
	item := new(Model)
	sql := "select * from ws_menu where id =?  and status < 44 limit 1 "
	v := mysql.ServerAPI.Engine().Raw(sql, id).Scan(item)
	return item, v.Error
}

func (e *SrvMenu) FindMulti() ([]*Model, error) {
	var result []*Model
	sql := "select * from ws_menu where status < 44 "
	v := mysql.ServerAPI.Engine().Raw(sql).Scan(&result)
	return result, v.Error
}

func (e *SrvMenu) Update(item *Model) (int64, error) {
	item.UpdatedAt = time.Now().Unix()
	sql := "update ws_menu set `Name` = ?,extended=?,updated_at=? where `id` = ? "
	v := mysql.ServerAPI.Engine().Exec(sql, item.Name, item.Extended, item.UpdatedAt, item.ID)
	return v.RowsAffected, v.Error
}

func (e *SrvMenu) UpdateStatus(item *Model) (int64, error) {
	item.UpdatedAt = time.Now().Unix()
	sql := "update ws_menu set `status` = ?,updated_at=? where `id` = ? "
	v := mysql.ServerAPI.Engine().Exec(sql, item.Status, item.UpdatedAt, item.ID)
	return v.RowsAffected, v.Error
}
