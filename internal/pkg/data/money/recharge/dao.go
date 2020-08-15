package recharge

import (
	"fmt"
	"github.com/go-liam/util/response"
	"grape/internal/pkg/database/mysql"
	"time"
)

var Server = new(SrvRecharge)

type SrvRecharge struct {
}

func (e *SrvRecharge) Create(item *Model) (int64, error) {
	item.CreatedAt = time.Now().Unix()
	item.UpdatedAt = item.CreatedAt
	v := mysql.ServerAPI.Engine().Create(item)
	return v.RowsAffected, v.Error
}

func (e *SrvRecharge) FindOne(id int64) (*Model, error) {
	item := new(Model)
	sql := "select * from um_recharge where id =?  and status < 44 limit 1 "
	v := mysql.ServerAPI.Engine().Raw(sql, id).Scan(item)
	return item, v.Error
}

// whereSt: " and id>1 " ,orderSt =" order by ID "
func (e *SrvRecharge) FindMulti(page *response.Pagination, s *response.ListParameter) ([]*Model, error) {
	var result []*Model
	sqlLimit := fmt.Sprintf(" limit %d , %d  ", (page.Current-1)*page.PageSize, page.PageSize)
	sqlWhere := " status < 44 " + s.WhereSt
	sql := "select * from um_recharge where " + sqlWhere + s.OrderSt + sqlLimit
	mysql.ServerAPI.Engine().Model(&Model{}).Where(sqlWhere).Count(&page.Total)
	v := mysql.ServerAPI.Engine().Raw(sql).Scan(&result)
	return result, v.Error
}

func (e *SrvRecharge) Update(item *Model) (int64, error) {
	item.UpdatedAt = time.Now().Unix()
	sql := "update um_recharge set extended=?,updated_at=?,remark=?  where `id` = ? "
	v := mysql.ServerAPI.Engine().Exec(sql, item.Extended, item.UpdatedAt, item.Remark, item.ID)
	return v.RowsAffected, v.Error
}

func (e *SrvRecharge) UpdateStatus(item *Model) (int64, error) {
	item.UpdatedAt = time.Now().Unix()
	sql := "update um_recharge set `status` = ?,updated_at=? where `id` = ? "
	v := mysql.ServerAPI.Engine().Exec(sql, item.Status, item.UpdatedAt, item.ID)
	return v.RowsAffected, v.Error
}
