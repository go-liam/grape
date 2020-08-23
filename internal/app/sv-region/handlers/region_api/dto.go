package region_api

import (
	"github.com/go-liam/util/conv"
	"grape/internal/pkg/data/region"
)

type ReqModel struct {
	UserID   string ` json:"userID"`
	RegionID int    `json:"regionID"  `
}

type RespModel struct {
	UserID    string ` json:"userID"`
	CreatedAt int64  ` json:"createdAt"` // 创建时间戳
	RegionID  int    `json:"regionID"  `
}

func GetModel(i *ReqModel) *region.Model {
	o := new(region.Model)
	o.UserID = conv.StringToInt64(i.UserID, 0)
	o.RegionID = i.RegionID
	return o
}

func GetRespModel(i *region.Model) *RespModel {
	o := new(RespModel)
	o.CreatedAt = i.CreatedAt
	o.RegionID = i.RegionID
	o.UserID = conv.Int64ToString(i.UserID)
	return o
}
