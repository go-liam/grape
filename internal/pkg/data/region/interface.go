package region

import "github.com/go-liam/util/response"

type Service interface {
	Create(item *Model) (int64, error)
	FindOne(id int64) (*Model, error)
	FindMulti(page *response.Pagination, s *response.ListParameter) ([]*Model, error)
	//cache
	CacheOne(id int64) (*Model, error)
}
