package menu_api

import (
	"github.com/go-liam/util/conv"
	"grape/internal/pkg/data/home_site/menu"
)

type ReqModel struct {
	ID         string      ` json:"id"`          // id
	LanguageID int         `json:"languageID"  ` // 语言ID:1中文,2英语
	Extended   interface{} `json:"extended" `    // 扩展的
	SiteID     int64       `json:"siteID"  `     // 网站ID
	Name       string      `json:"name"  `       // 菜单名
	ParentID   int64       `json:"parentID"  `   // 父亲ID
	Level      int         `json:"level"  `      //级别
}

type RespModel struct {
	ID         string      ` json:"id"`          // id
	CreatedAt  int64       ` json:"created_at"`  // 创建时间戳
	UpdatedAt  int64       ` json:"updated_at"`  // 更新时间戳
	Status     int8        `json:"status"  `     // 44 删除, 1 启用, 4 禁用
	LanguageID int         `json:"languageID"  ` // 语言ID:1中文,2英语
	Extended   interface{} `json:"extended" `    // 扩展的
	SiteID     int64       `json:"siteID"  `     // 网站ID
	Name       string      `json:"name"  `       // 菜单名
	ParentID   int64       `json:"parentID"  `   // 父亲ID
	Level      int         `json:"level"  `      //级别
}

func GetModel(i *ReqModel) *menu.Model {
	o := new(menu.Model)
	o.ID = conv.StringToInt64(i.ID, 0)
	o.LanguageID = i.LanguageID
	o.Extended = conv.StructToJsonString(i.Extended)
	o.SiteID = i.SiteID
	o.Name = i.Name
	o.ParentID = i.ParentID
	o.Level = i.Level
	return o
}

func GetRespModel(i *menu.Model) *RespModel {
	o := new(RespModel)
	o.Status = i.Status
	o.ID = conv.Int64ToString(i.ID)
	o.LanguageID = i.LanguageID
	o.Extended = conv.StringToInterface(i.Extended)
	o.SiteID = i.SiteID
	o.Name = i.Name
	o.ParentID = i.ParentID
	o.Level = i.Level
	return o
}
