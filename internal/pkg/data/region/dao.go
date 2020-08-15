package region

import (
	"fmt"
	"github.com/go-liam/util/response"
	"grape/internal/pkg/database/mysql"
	models "grape/internal/pkg/model"
	"time"
)

var Server = new(SrvRegion)

type SrvRegion struct {
}

func (e *SrvRegion) Create(item *Model) (int64, error) {
	item.CreatedAt = time.Now().Unix()
	v := mysql.ServerAPI.Engine().Create(item)
	return v.RowsAffected, v.Error
}

func (e *SrvRegion) FindOne(id int64) (*Model, error) {
	item := new(Model)
	sql := "select * from ar_user_region where user_id = ? limit 1 "
	v := mysql.ServerAPI.Engine().Raw(sql, id).Scan(item)
	return item, v.Error
}

// whereSt: " and id>1 " ,orderSt =" order by ID "
func (e *SrvRegion) FindMulti(page *response.Pagination, s *response.ListParameter) ([]*Model, error) {
	var result []*Model
	sqlLimit := fmt.Sprintf(" limit %d , %d  ", (page.Current-1)*page.PageSize, page.PageSize)
	sqlWhere := " 1 = 1 " + s.WhereSt + " "
	sql := "select * from ar_user_region where " + sqlWhere + s.OrderSt + sqlLimit
	mysql.ServerAPI.Engine().Model(&Model{}).Where(sqlWhere).Count(&page.Total)
	v := mysql.ServerAPI.Engine().Raw(sql).Scan(&result)
	return result, v.Error
}

//func (e *SrvRegion) Update(item *Model) (int64, error) {
//	item.UpdatedAt = time.Now().Unix()
//	sql := "update ar_user_region set `region_id` = ? ,updated_at=?  where `user_id` = ? "
//	v := mysql.ServerAPI.Engine().Exec(sql, item.RegionID,item.UpdatedAt, item.UserID)
//	return v.RowsAffected, v.Error
//}
