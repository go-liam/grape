package jwt

const (
	ConstClientDefault = 0
	ConstClientWeb     = 1
	ConstClientAndroid = 2
	ConstClientIPad    = 3
	ConstClientIPhone  = 4
)

type ModelUserToken struct {
	UserID   int64  `json:"userID"`
	Token    string `json:"token"`
	ClientID int    `json:"client"`
}
