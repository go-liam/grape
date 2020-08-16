package bind

type Service interface {
	Create(item *Model) (int64, error)
	FindOne(id int64) (*Model, error)
}
