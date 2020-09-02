package role_api

import (
	"github.com/go-liam/util/conv"
	"grape/internal/pkg/data/rbac/role"
)

type ReqModel struct {
	ID       string      ` json:"id"`       // id
	Extended interface{} `json:"extended" ` // 扩展的
	Name     string      `json:"name" `
	PowerIDS string      `json:"powerIDS"  `
}

type RespModel struct {
	ID        string      ` json:"id"`        // id
	CreatedAt int64       ` json:"createdAt"` // 创建时间戳
	UpdatedAt int64       ` json:"updatedAt"` // 更新时间戳
	Status    int8        `json:"status"  `   // 44 删除, 1 启用, 4 禁用
	Extended  interface{} `json:"extended" `  // 扩展的
	Name      string      `json:"name" `
	PowerIDS  string      `json:"powerIDS"  `
}

func GetModel(i *ReqModel) *role.Model {
	o := new(role.Model)
	o.ID = conv.StringToInt64(i.ID, 0)
	o.Name = i.Name
	o.PowerIDS = i.PowerIDS
	o.Extended = conv.StructToJsonString(i.Extended)
	return o
}

func GetRespModel(i *role.Model) *RespModel {
	o := new(RespModel)
	o.Status = i.Status
	o.ID = conv.Int64ToString(i.ID)
	o.Extended = conv.StringToInterface(i.Extended)
	o.Name = i.Name
	o.PowerIDS = i.PowerIDS
	return o
}