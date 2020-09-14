package server

import (
	"fmt"
	"github.com/go-liam/util/conv"
	"github.com/go-liam/util/response"
	"grape/internal/pkg/database/mysql"
	"time"
)

var Server = new(SrvServer)

type SrvServer struct {
}

func (e *SrvServer) Create(item *Model) (int64, error) {
	item.CreatedAt = time.Now().Unix()
	item.UpdatedAt = item.CreatedAt
	v := mysql.ServerAPI.Engine().Create(item)
	return v.RowsAffected, v.Error
}

func (e *SrvServer) FindOne(id int64) (*Model, error) {
	item := new(Model)
	sql := "select * from sa_server where id =?  and status < 44 limit 1 "
	v := mysql.ServerAPI.Engine().Raw(sql, id).Scan(item)
	return item, v.Error
}

func (e *SrvServer) FindMultiByNil() ([]*Model, error) {
	var result []*Model
	sql := "select * from sa_server where status < 44 "
	v := mysql.ServerAPI.Engine().Raw(sql).Scan(&result)
	return result, v.Error
}

func (e *SrvServer) FindMulti(p *response.Pagination, s *response.ListParameter) ([]*Model, error) {
	var result []*Model
	sqlLimit := fmt.Sprintf(" limit %d , %d  ", (p.Current-1)*p.PageSize, p.PageSize)
	sqlWhere := " status < 44 " + s.WhereSt
	sql := "select * from sa_server where " + sqlWhere + s.OrderSt + sqlLimit
	mysql.ServerAPI.Engine().Model(&Model{}).Where(sqlWhere).Count(&p.Total)
	v := mysql.ServerAPI.Engine().Raw(sql).Scan(&result)
	return result, v.Error
}

func (e *SrvServer) Update(item *Model) (int64, error) {
	item.UpdatedAt = time.Now().Unix()
	sql := "update sa_server set `title` = ?,description=?,extended=?,updated_at=?,url= ?  where `id` = ? "
	v := mysql.ServerAPI.Engine().Exec(sql, item.Title, item.Description, item.Extended, item.UpdatedAt, item.URL, item.ID)
	return v.RowsAffected, v.Error
}

func (e *SrvServer) UpdateStatus(item *Model) (int64, error) {
	item.UpdatedAt = time.Now().Unix()
	sql := "update sa_server set `status` = ?,updated_at=? where `id` = ? "
	v := mysql.ServerAPI.Engine().Exec(sql, item.Status, item.UpdatedAt, item.ID)
	return v.RowsAffected, v.Error
}

func (e *SrvServer) UpdateStatusByIDs(status int, ls []int64) (int64, error) {
	updatedAt := time.Now().Unix()
	ids := conv.ArrayToString(ls, ",")
	sql := fmt.Sprintf("update sa_server set `status` = ?,updated_at=? where `id` in (%s) ", ids)
	v := mysql.ServerAPI.Engine().Exec(sql, status, updatedAt)
	return v.RowsAffected, v.Error
}

func (e *SrvServer) UpdateHealth(item *Model) (int64, error) {
	item.UpdatedAt = time.Now().Unix()
	sql := "update sa_server set `run_status` = ?,health=?, updated_at=? where `id` = ? "
	v := mysql.ServerAPI.Engine().Exec(sql, item.RunStatus, item.Health, item.UpdatedAt, item.ID)
	return v.RowsAffected, v.Error
}
