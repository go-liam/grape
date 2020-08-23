package user_cms

import (
	"github.com/go-liam/util/conv"
	"grape/internal/pkg/data/rbac/user"
)

type ReqModel struct {
	ID       string      ` json:"id"`       // id
	Extended interface{} `json:"extended" ` // 扩展的
	Flag     int8        ` json:"flag"`
	Remark   string      ` json:"remark"`
	RoleIDs  string      ` json:"roleIDs"`
}

type RespModel struct {
	ID        string      ` json:"id"`        // id
	CreatedAt int64       ` json:"createdAt"` // 创建时间戳
	UpdatedAt int64       ` json:"updatedAt"` // 更新时间戳
	Status    int8        `json:"status"  `   // 44 删除, 1 启用, 4 禁用
	Extended  interface{} `json:"extended" `  // 扩展的
	Flag      int8        ` json:"flag"`
	Remark    string      ` json:"remark"`
	RoleIDs   string      ` json:"roleIDs"`
}

func GetModel(i *ReqModel) *user.Model {
	o := new(user.Model)
	o.ID = conv.StringToInt64(i.ID, 0)
	o.Flag = i.Flag
	o.Remark = i.Remark
	o.RoleIDs = i.RoleIDs
	o.Extended = conv.StructToJsonString(i.Extended)
	return o
}

func GetRespModel(i *user.Model) *RespModel {
	o := new(RespModel)
	o.Status = i.Status
	o.ID = conv.Int64ToString(i.ID)
	o.Extended = conv.StringToInterface(i.Extended)
	o.Flag = i.Flag
	o.Remark = i.Remark
	o.RoleIDs = i.RoleIDs
	return o
}
