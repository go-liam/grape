package account

import (
	"errors"
	"github.com/go-liam/util/encrypt/md5"
	"github.com/go-liam/util/uuid"
	"grape/internal/pkg/data/account/bind"
	"grape/internal/pkg/data/account/user"
	"grape/internal/pkg/data/region"
)

const (
	keyMd5 = "123@GciOiJIUzIInR5cCI6IkpXVCJ9#$%@liam"
)

type RegModel struct {
	UserName string
	Password string
	Status   int8
	Remark   string
	Extended string
	Email    string
	Phone    string
	NickName string
	RegionID int
}

func RegByNamePassword(i *RegModel) (int64, error) {
	// create uid
	ck, _ := bind.Server.FindOne(bind.TypeUserName, i.UserName)
	if ck.UserID > 0 {
		return 0, errors.New("[ERROR]RegByNamePassword: had name")
	}
	uid := uuid.AutoInt64ID()
	// add region
	region.Server.CheckAndCreate(&region.Model{UserID: uid, RegionID: i.RegionID})
	// add name
	bModel := &bind.Model{UserID: uid, Type: bind.TypeUserName, Name: i.UserName, ID: uid}
	got, _ := bind.Server.Create(bModel)
	if got <= 0 {
		return 0, errors.New("[ERROR]RegByNamePassword: add name error ")
	}
	// add psw
	psw := md5.GetMD5Hash(bind.KeyMd5 + i.Password)
	pModel := &bind.ModelPassword{ID: uid, Password: psw}
	bind.ServerPsd.CheckAndCreate(pModel)
	// add user
	uModel := &user.Model{
		ID:       uid,
		Status:   i.Status,
		Remark:   i.Remark,
		Extended: i.Extended,
		Email:    i.Email,
		Phone:    i.Phone,
		Name:     i.UserName,
		NickName: i.NickName,
	}
	user.Server.CheckAndCreate(uModel)
	return uid, nil
}
