package log_cms

import (
	"github.com/go-liam/util/conv"
	"grape/internal/pkg/data/sylog"
)

type ReqModel struct {
	ID       string      ` json:"id"`       // id
	Extended interface{} `json:"extended" ` // 扩展的
	Level    int32       ` json:"level"`
	Msg      string      ` json:"msg"`
}

type RespModel struct {
	ID        string      ` json:"id"`        // id
	CreatedAt int64       ` json:"createdAt"` // 创建时间戳
	Extended  interface{} `json:"extended" `  // 扩展的
	Level     int32       ` json:"level"`
	Msg       string      ` json:"msg"`
}

func GetModel(i *ReqModel) *sylog.Model {
	o := new(sylog.Model)
	o.ID = conv.StringToInt64(i.ID, 0)
	o.Level = i.Level
	o.Msg = i.Msg
	o.Extended = conv.StructToJsonString(i.Extended)
	return o
}

func GetRespModel(i *sylog.Model) *RespModel {
	o := new(RespModel)
	o.CreatedAt = i.CreatedAt
	o.Level = i.Level
	o.Msg = i.Msg
	o.ID = conv.Int64ToString(i.ID)
	o.Extended = conv.StringToInterface(i.Extended)
	return o
}
