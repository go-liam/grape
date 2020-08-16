package sylog

import "github.com/go-liam/util/response"

type Service interface {
	Create(item *Model) (int64, error)
	FindMulti(page *response.Pagination, s *response.ListParameter) ([]*Model, error)
	Delete(beforeTime int64) (int64, error)
	//other
	WriteLog(logID int64, msg string, level int32) error
	SetConfig(item *ConfigModel) error
}
