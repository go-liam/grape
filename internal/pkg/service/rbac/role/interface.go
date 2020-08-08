package role

type InRole interface {
	List() ([]*Model, error)
	Add(item *Model) (int, error)
	Edit(item *Model) (int, error)
	UpdatePower(item *Model) (int, error)
	Delete(item *Model) (int, error)
	FindOne(roleID int64) (*Model, error)
}
