package sylog

import (
	"fmt"
	"github.com/go-liam/util/response"
	"grape/internal/pkg/database/mysql"
	"log"
	"time"
)

var Server = new(SrvLog)

type SrvLog struct {
}

func (e *SrvLog) SetConfig(item *ConfigModel) error {
	log.Printf("[INFO]log-config-%+v\n", item)
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
func (e *SrvLog) FindMulti(p *response.Pagination, s *response.ListParameter) ([]*Model, error) {
	var result []*Model
	sqlLimit := fmt.Sprintf(" limit %d , %d  ", (p.Current-1)*p.PageSize, p.PageSize)
	sqlWhere := " 1=1 " + s.WhereSt
	sql := "select * from lg_log where " + sqlWhere + s.OrderSt + sqlLimit
	mysql.ServerAPI.Engine().Model(&Model{}).Where(sqlWhere).Count(&p.Total)
	v := mysql.ServerAPI.Engine().Raw(sql).Scan(&result)
	return result, v.Error
}

// 删除 time 之前的数据
func (e *SrvLog) Delete(beforeTime int64) (int64, error) {
	sql := "delete from lg_log where `created_at` <=  ? limit 10000 "
	v := mysql.ServerAPI.Engine().Exec(sql, beforeTime)
	return v.RowsAffected, v.Error
}
