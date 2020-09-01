package account

import (
	"errors"
	"github.com/go-liam/util/uuid"
	"grape/internal/pkg/data/account/bind"
	"grape/internal/pkg/data/account/user"
	"grape/internal/pkg/data/region"
)

type RegModel struct {
	UserName string
	Password string
	Status    int8
	Remark    string
	Extended  string
	Email     string
	Phone     string
	Name      string
}

func RegByNamePassword(i *RegModel) (int64,error) {
	// create uid
	ck,_ := bind.Server.FindOne(bind.TypeUserName,i.UserName)
	if ck.UserID > 0 {
		return 0,errors.New("[ERROR]RegByNamePassword: had name")
	}
	uid :=uuid.AutoInt64ID()
	// add region
	got1 ,_:= region.Server.FindOne(uid)
	if got1.UserID == 0{
		region.Server.Create(&region.Model{UserID:uid,RegionID:1})
	}
	// add name
	bModel := &bind.Model{UserID:uid, Type:bind.TypeUserName,Name:i.UserName ,ID:uid}
	got,_:= bind.Server.Create(bModel)
	if got <=0 {
		return 0,errors.New("[ERROR]RegByNamePassword: add name error ")
	}
	// add psw
	pModel := &bind.Model{UserID:uid, Type:bind.TypeUserPsw,Name:i.Password ,ID:uuid.AutoInt64ID()}
	bind.Server.Create(pModel)
	// add user
	uModel := &user.Model{
		ID:        uid,
		Status:    i.Status,
		Remark:    i.Remark,
		Extended:  i.Extended,
		Email:     i.Email,
		Phone:     i.Phone,
		Name:      i.Name,
	}
	user.Server.CheckAndCreate(uModel)
	return uid,nil
}