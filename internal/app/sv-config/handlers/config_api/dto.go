package config_api

import (
	"github.com/go-liam/util/conv"
	"grape/internal/pkg/data/config"
)

type ReqModel struct {
	ID       string      ` json:"id"` // id
	Content  string      `json:"content"  `
	Version  int         `  json:"version"`
	Extended interface{} `json:"extended" ` // 扩展的
}

type RespModel struct {
	ID        string      ` json:"id"`         // id
	CreatedAt int64       ` json:"created_at"` // 创建时间戳
	UpdatedAt int64       ` json:"updated_at"` // 更新时间戳
	Status    int8        `json:"status"  `    // 44 删除, 1 启用, 4 禁用
	Extended  interface{} `json:"extended" `   // 扩展的
	Content   string      `json:"content"  `
	Version   int         `  json:"version"`
}

func GetModel(i *ReqModel) *config.Model {
	o := new(config.Model)
	o.ID = conv.StringToInt64(i.ID, 0)
	o.Content = i.Content
	o.Version = i.Version
	o.Extended = conv.StructToJsonString(i.Extended)
	return o
}

func GetRespModel(i *config.Model) *RespModel {
	o := new(RespModel)
	o.Content = i.Content
	o.Status = i.Status
	o.ID = conv.Int64ToString(i.ID)
	o.Version = i.Version
	o.Extended = conv.StringToInterface(i.Extended)
	return o
}
