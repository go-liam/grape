package account

import (
	"errors"
	"github.com/go-liam/util/encrypt/md5"
	"grape/internal/pkg/data/account/bind"
)

func LoginByPassword(userName string, password string) (int64, error) {
	g1, _ := bind.Server.FindOne(bind.TypeUserName, userName)
	if g1.UserID <= 0 {
		return 0, errors.New("no user")
	}
	g2, _ := bind.ServerPsd.FindOne(g1.UserID)
	if g2.ID <= 0 {
		return 0, errors.New("no passsword")
	}
	psw := md5.GetMD5Hash(bind.KeyMd5 + password)
	if psw != g2.Password {
		return 0, errors.New("passsword error")
	}
	return g1.UserID, nil
}
