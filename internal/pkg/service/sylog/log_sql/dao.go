package log_sql

import (
	"fmt"
	"github.com/go-liam/util/response"
	"grape/internal/pkg/database/mysql"
	models "grape/internal/pkg/model"
	"time"
)

var Server = new(SrvLog)

type SrvLog struct {
}

func (e *SrvLog) Init(config string) error {
	println(config)
	return nil
}

func (e *SrvLog) WriteLog(logID int64, msg string, level int32) error {
	item := new(Model)
	item.Msg = msg
	item.ID = logID
	item.Level = level
	_, err := new(SrvLog).Create(item)
	return err
}

func (e *SrvLog) Create(item *Model) (int64, error) {
	item.CreatedAt = time.Now().Unix()
	v := mysql.ServerAPI.Engine().Create(item)
	return v.RowsAffected, v.Error
}

// whereSt: " and id>1 " ,orderSt =" order by ID "
func (e *SrvLog) FindMulti(page *response.Pagination, s *models.ListParameter) ([]*Model, error) {
	var result []*Model
	sqlLimit := fmt.Sprintf(" limit %d , %d  ", (page.Current-1)*page.PageSize, page.PageSize)
	sqlWhere := " 1=1 "+ s.WhereSt
	sql := "select * from lg_log where " + sqlWhere + s.OrderSt + sqlLimit
	mysql.ServerAPI.Engine().Model(&Model{}).Where(sqlWhere).Count(&page.Total)
	v := mysql.ServerAPI.Engine().Raw(sql).Scan(&result)
	return result, v.Error
}

// 删除 time 之前的数据
func (e *SrvLog) Delete(beforeTime int64) (int64, error) {
	sql := "delete from lg_log where `created_at` <=  ? limit 10000 "
	v := mysql.ServerAPI.Engine().Exec(sql, beforeTime)
	return v.RowsAffected, v.Error
}
