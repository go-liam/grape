package role

type InRole interface {
	List()([]*ModelRole,error)
	Add(item*ModelRole)(int,error)
	Edit(item*ModelRole)(int,error)
	UpdatePower (item*ModelRole)(int,error)
	Delete(item*ModelRole)(int,error)
	FindOne(roleID int) (*ModelRole, error)
}