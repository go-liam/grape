package site

import "github.com/go-liam/util/response"

type Site interface {
	Create(item *Model) (int64, error)
	FindOne(id int64) (*Model, error)
	FindMulti(page *response.Pagination, s *response.ListParameter) ([]*Model, error)
	Update(item *Model) (int64, error)
	UpdateStatusByIDs(status int, ls []int64) (int64, error)
	//cache
	CacheOne(id int64) (*Model, error)
}
