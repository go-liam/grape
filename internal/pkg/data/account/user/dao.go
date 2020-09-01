package user

import (
	"errors"
	"fmt"
	"github.com/go-liam/util/conv"
	"github.com/go-liam/util/response"
	"grape/internal/pkg/database/mysql"
	"time"
)

var Server = new(SrvUser)

type SrvUser struct {
}

func (e *SrvUser) CheckAndCreate(item *Model) (int64, error) {
	ck,_:= e.FindOne(item.ID)
	if ck.ID > 0 {
		return 0,errors.New("had data")
	}
	return e.Create(item)
}

func (e *SrvUser) Create(item *Model) (int64, error) {
	item.CreatedAt = time.Now().Unix()
	item.UpdatedAt = item.CreatedAt
	v := mysql.ServerAPI.Engine().Create(item)
	return v.RowsAffected, v.Error
}

func (e *SrvUser) FindOne(id int64) (*Model, error) {
	item := new(Model)
	sql := "select * from uc_user where id =?  and status < 44 limit 1 "
	v := mysql.ServerAPI.Engine().Raw(sql, id).Scan(item)
	return item, v.Error
}

// whereSt: " and id>1 " ,orderSt =" order by ID "
func (e *SrvUser) FindMulti(p *response.Pagination, s *response.ListParameter) ([]*Model, error) {
	var result []*Model
	sqlLimit := fmt.Sprintf(" limit %d , %d  ", (p.Current-1)*p.PageSize, p.PageSize)
	sqlWhere := " status < 44 " + s.WhereSt
	sql := "select * from uc_user where " + sqlWhere + s.OrderSt + sqlLimit
	mysql.ServerAPI.Engine().Model(&Model{}).Where(sqlWhere).Count(&p.Total)
	v := mysql.ServerAPI.Engine().Raw(sql).Scan(&result)
	return result, v.Error
}

func (e *SrvUser) Update(item *Model) (int64, error) {
	item.UpdatedAt = time.Now().Unix()
	sql := "update uc_user set extended=?,updated_at=?,remark=?,email= ?,phone = ?  where `id` = ? "
	v := mysql.ServerAPI.Engine().Exec(sql, item.Extended, item.UpdatedAt, item.Remark, item.Email, item.Phone,
		item.ID)
	return v.RowsAffected, v.Error
}

func (e *SrvUser) UpdateStatus(item *Model) (int64, error) {
	item.UpdatedAt = time.Now().Unix()
	sql := "update uc_user set `status` = ?,updated_at=? where `id` = ? "
	v := mysql.ServerAPI.Engine().Exec(sql, item.Status, item.UpdatedAt, item.ID)
	return v.RowsAffected, v.Error
}

func (e *SrvUser) UpdateStatusByIDs(status int, ls []int64) (int64, error) {
	updatedAt := time.Now().Unix()
	ids := conv.ArrayToString(ls, ",")
	sql := fmt.Sprintf("update uc_user set `status` = ?,updated_at=? where `id` in (%s) ", ids)
	v := mysql.ServerAPI.Engine().Exec(sql, status, updatedAt)
	return v.RowsAffected, v.Error
}
