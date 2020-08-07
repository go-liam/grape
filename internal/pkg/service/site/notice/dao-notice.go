package notice

import (
	"fmt"
	"github.com/go-liam/util/response"
	"grape/internal/pkg/database/mysql"
	models "grape/internal/pkg/model"
	"time"
)

var ServerNotice = new(SrvNotice)

type SrvNotice struct {
}

func (e *SrvNotice) Create(item *NoticeModel) (int64, error) {
	item.CreatedAt = time.Now().Unix()
	item.UpdatedAt = item.CreatedAt
	v := mysql.ServerAPI.Engine().Create(item)
	return v.RowsAffected, v.Error
}

func (e *SrvNotice) FindOne(id int64) (*NoticeModel, error) {
	item := new(NoticeModel)
	sql := "select * from ws_notice where id =?  and status < 44 limit 1 "
	v := mysql.ServerAPI.Engine().Raw(sql, id).Scan(item)
	return item, v.Error
}

// whereSt: " and id>1 " ,orderSt =" order by ID "
func (e *SrvNotice) FindMulti(page *response.Pagination, s *models.ListParameter) ([]*NoticeModel, error) {
	var result []*NoticeModel
	sqlLimit := fmt.Sprintf(" limit %d , %d  ", (page.Current-1)*page.PageSize, page.PageSize)
	sqlWhere := " status < 44 " + s.WhereSt
	sql := "select * from ws_notice where " + sqlWhere + s.OrderSt + sqlLimit
	mysql.ServerAPI.Engine().Model(&NoticeModel{}).Where(sqlWhere).Count(&page.Total)
	v := mysql.ServerAPI.Engine().Raw(sql).Scan(&result)
	return result, v.Error
}

func (e *SrvNotice) Update(item *NoticeModel) (int64, error) {
	item.UpdatedAt = time.Now().Unix()
	sql := "update ws_notice set extended=?,updated_at=?,title=?,author=?,content=? where `id` = ? "
	v := mysql.ServerAPI.Engine().Exec(sql, item.Extended, item.UpdatedAt, item.Title, item.Author,
		item.Content, item.ID)
	return v.RowsAffected, v.Error
}

func (e *SrvNotice) UpdateState(item *NoticeModel) (int64, error) {
	item.UpdatedAt = time.Now().Unix()
	sql := "update ws_notice set `status` = ?,updated_at=? where `id` = ? "
	v := mysql.ServerAPI.Engine().Exec(sql, item.Status, item.UpdatedAt, item.ID)
	return v.RowsAffected, v.Error
}
