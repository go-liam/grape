package jwt

type InJWTUser interface {
	UpdateOrAdd(item *ModelUserToken) (bool, error)
	GET(userID int64, clientID int) (*ModelUserToken, error)
	DeleteAllByUserID(userID int64) (bool, error)
}
