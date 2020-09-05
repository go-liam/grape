package role

import (
	"fmt"
	"github.com/go-liam/util/conv"
	"github.com/go-liam/util/response"
	"grape/internal/pkg/database/mysql"
	"time"
)

var Server = new(SrvRole)

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

func (e *SrvRole) FindMultiByNil() ([]*Model, error) {
	var result []*Model
	sql := "select * from rb_role where status < 44 "
	v := mysql.ServerAPI.Engine().Raw(sql).Scan(&result)
	return result, v.Error
}

func (e *SrvRole) FindMulti(p *response.Pagination, s *response.ListParameter) ([]*Model, error) {
	var result []*Model
	sqlLimit := fmt.Sprintf(" limit %d , %d  ", (p.Current-1)*p.PageSize, p.PageSize)
	sqlWhere := " status < 44 " + s.WhereSt
	sql := "select * from rb_role where " + sqlWhere + s.OrderSt + sqlLimit
	mysql.ServerAPI.Engine().Model(&Model{}).Where(sqlWhere).Count(&p.Total)
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

func (e *SrvRole) UpdateStatusByIDs(status int, ls []int64) (int64, error) {
	updatedAt := time.Now().Unix()
	ids := conv.ArrayToString(ls, ",")
	sql := fmt.Sprintf("update rb_role set `status` = ?,updated_at=? where `id` in (%s) ", ids)
	v := mysql.ServerAPI.Engine().Exec(sql, status, updatedAt)
	return v.RowsAffected, v.Error
}

func (e *SrvRole) FindMoreByList(list []*Model, groupIDs []int64) []*Model {
	ls := make([]*Model, 0)
	for _, v := range list {
		if conv.ArrayInt64Contains(groupIDs, v.ID) {
			ls = append(ls, v)
		}
	}
	return ls
}

func (e *SrvRole) FindMultiByIDs(ids []int64) ([]*Model, error) {
	result := make([]*Model, 0)
	st := conv.ArrayToString(ids, ",")
	sql := fmt.Sprintf("select * from rb_role where id in (%v)  and status < 44  ", st)
	v := mysql.ServerAPI.Engine().Raw(sql).Scan(&result)
	return result, v.Error
}

func (e *SrvRole) GetPowerIDsByIDs(roleIDs []int64) []int64 {
	arr := make([]int64, 0)
	if len(roleIDs) == 0 {
		return arr
	}
	ls, _ := e.FindMultiByIDs(roleIDs)
	for _, v := range ls {
		ay := conv.StringToInt64Array(v.PowerIDS)
		arr = append(arr, ay...)
	}
	return conv.RemoveDuplicateArrayInt64(arr)
}
