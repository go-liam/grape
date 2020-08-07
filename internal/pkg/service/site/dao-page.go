package site

import (
	"fmt"
	"github.com/go-liam/util/response"
	"grape/internal/pkg/database/mysql"
	models "grape/internal/pkg/model"
	"time"
)

var ServerPage = new(SrvPage)

type SrvPage struct {
}

func (e *SrvPage) Create(item *PageModel) (int64, error) {
	item.CreatedAt = time.Now().Unix()
	item.UpdatedAt = item.CreatedAt
	v := mysql.ServerAPI.Engine().Create(item)
	return v.RowsAffected, v.Error
}

func (e *SrvPage) FindOne(id int64) (*PageModel, error) {
	item := new(PageModel)
	sql := "select * from ws_page where id =?  and status < 44 limit 1 "
	v := mysql.ServerAPI.Engine().Raw(sql, id).Scan(item)
	return item, v.Error
}

// whereSt: " and id>1 " ,orderSt =" order by ID "
func (e *SrvPage) FindMulti(page *response.Pagination, s *models.ListParameter) ([]*PageModel, error) {
	var result []*PageModel
	sqlLimit := fmt.Sprintf(" limit %d , %d  ", (page.Current-1)*page.PageSize, page.PageSize)
	sqlWhere := " status < 44 " + s.WhereSt + " "
	sql := "select * from ws_page where " + sqlWhere + s.OrderSt + sqlLimit
	mysql.ServerAPI.Engine().Model(&PageModel{}).Where(sqlWhere).Count(&page.Total)
	v := mysql.ServerAPI.Engine().Raw(sql).Scan(&result)
	return result, v.Error
}

func (e *SrvPage) Update(item *PageModel) (int64, error) {
	item.UpdatedAt = time.Now().Unix()
	sql := "update ws_page set extended=?,updated_at=?,title=?,author=?,description=?,content=? where `id` = ? "
	v := mysql.ServerAPI.Engine().Exec(sql, item.Extended, item.UpdatedAt, item.Title, item.Author,
		item.Description, item.Content, item.ID)
	return v.RowsAffected, v.Error
}

func (e *SrvPage) UpdateState(item *PageModel) (int64, error) {
	item.UpdatedAt = time.Now().Unix()
	sql := "update ws_page set `status` = ?,updated_at=? where `id` = ? "
	v := mysql.ServerAPI.Engine().Exec(sql, item.Status, item.UpdatedAt, item.ID)
	return v.RowsAffected, v.Error
}
