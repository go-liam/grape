package power

type InPower interface {
	List(powerType int8) ([]*ModelPower, error)
	Add(item *ModelPower) (int, error)
	Delete(item *ModelPower) (int, error)
	Edit(item *ModelPower) (int, error)
	FindOne(powerID int) (*ModelPower, error)
}
