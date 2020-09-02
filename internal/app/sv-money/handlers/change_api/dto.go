package change_api

import (
	"github.com/go-liam/util/conv"
	"grape/internal/pkg/data/money/change"
)

type ReqModel struct {
	ID         string      ` json:"id"`       // id
	Extended   interface{} `json:"extended" ` // 扩展的
	Remark     string      `  json:"remark"`
	UserID     int64       ` json:"userID"`
	Amount     int         ` json:"amount"`
	LastAmount int         ` json:"lastAmount"`
	Type       int8        ` json:"type"`
	OrderID    int64       ` json:"orderID"`
	MoneyType  int8        ` json:"moneyType"`
}

type RespModel struct {
	ID         string      ` json:"id"`         // id
	CreatedAt  int64       ` json:"created_at"` // 创建时间戳
	UpdatedAt  int64       ` json:"updated_at"` // 更新时间戳
	Status     int8        `json:"status"  `    // 44 删除, 1 启用, 4 禁用
	Extended   interface{} `json:"extended" `   // 扩展的
	Remark     string      `  json:"remark"`
	UserID     int64       ` json:"userID"`
	Amount     int         ` json:"amount"`
	LastAmount int         ` json:"lastAmount"`
	Type       int8        ` json:"type"`
	OrderID    int64       ` json:"orderID"`
	MoneyType  int8        ` json:"moneyType"`
}

func GetModel(i *ReqModel) *change.Model {
	o := new(change.Model)
	o.ID = conv.StringToInt64(i.ID, 0)
	o.Extended = conv.StructToJsonString(i.Extended)
	o.Remark = i.Remark
	o.UserID = i.UserID
	o.Amount = i.Amount
	o.LastAmount = i.LastAmount
	o.Type = i.Type
	o.OrderID = i.OrderID
	o.MoneyType = i.MoneyType
	return o
}

func GetRespModel(i *change.Model) *RespModel {
	o := new(RespModel)
	o.Status = i.Status
	o.ID = conv.Int64ToString(i.ID)
	o.CreatedAt = i.CreatedAt
	o.UpdatedAt = i.UpdatedAt
	o.Extended = conv.StringToInterface(i.Extended)
	o.Remark = i.Remark
	o.UserID = i.UserID
	o.Amount = i.Amount
	o.LastAmount = i.LastAmount
	o.Type = i.Type
	o.OrderID = i.OrderID
	o.MoneyType = i.MoneyType
	return o
}
