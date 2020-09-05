package account

import (
	"github.com/go-liam/util/encrypt/md5"
	"grape/internal/pkg/data/account/bind"
)

func UpdatePassword(userID int64, password string) (int64, error) {
	psw := md5.GetMD5Hash(bind.KeyMd5 + password)
	pModel := &bind.ModelPassword{ID: userID, Password: psw}
	return bind.ServerPsd.Update(pModel)
}
