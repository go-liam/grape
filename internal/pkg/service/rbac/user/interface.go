package user

type InUser interface {
	List() ([]*ModelUser, error)
	Add(item *ModelUser) (int, error)
	Edit(item *ModelUser) (int, error)
	UpdateRole(item *ModelUser) (int, error)
	Delete(item *ModelUser) (int, error)
	FindOne(userID int64) (*ModelUser, error)
}
