package user_api

import (
	"github.com/go-liam/util/conv"
	"grape/internal/pkg/data/account/user"
)

type ReqModel struct {
	ID       string      ` json:"id"` // id
	Name     string      ` json:"name"`
	Email    string      `json:"email"  ` //邮箱
	Phone    string      ` json:"phone"`
	Extended interface{} `json:"extended" ` // 扩展的
	Remark   string      ` json:"remark"`
}

type RespModel struct {
	ID        string      ` json:"id"`         // id
	CreatedAt int64       ` json:"created_at"` // 创建时间戳
	UpdatedAt int64       ` json:"updated_at"` // 更新时间戳
	Status    int8        `json:"status"  `    // 44 删除, 1 启用, 4 禁用
	Email     string      `json:"email"  `     //邮箱
	Extended  interface{} `json:"extended" `   // 扩展的
	Phone     string      ` json:"phone"`
	Remark    string      ` json:"remark"`
	Name      string      ` json:"name"`
}

func GetModel(i *ReqModel) *user.Model {
	o := new(user.Model)
	o.ID = conv.StringToInt64(i.ID, 0)
	o.Name = i.Name
	o.Email = i.Email
	o.Phone = i.Phone
	o.Remark = i.Remark
	o.Extended = conv.StructToJsonString(i.Extended)
	return o
}

func GetRespModel(i *user.Model) *RespModel {
	o := new(RespModel)
	o.Phone = i.Phone
	o.Status = i.Status
	o.ID = conv.Int64ToString(i.ID)
	o.Email = i.Email
	o.Name = i.Name
	o.Remark = i.Remark
	o.CreatedAt = i.CreatedAt
	o.UpdatedAt = i.UpdatedAt
	o.Extended = conv.StringToInterface(i.Extended)
	return o
}
