package page_cms

import (
	"github.com/go-liam/util/conv"
	"grape/internal/pkg/data/home_site/page"
)

type ReqModel struct {
	ID          string      ` json:"id"`           // id
	LanguageID  int         `json:"languageID"  `  // 语言ID:1中文,2英语
	Title       string      `json:"title"  `       //标题
	Description string      `json:"description"  ` //描述
	Extended    interface{} `json:"extended" `     // 扩展的
	MenuID      int         `json:"menuID"  `
	SiteID      int         `json:"siteID"  `
}

type RespModel struct {
	ID          string      ` json:"id"`           // id
	CreatedAt   int64       ` json:"created_at"`   // 创建时间戳
	UpdatedAt   int64       ` json:"updated_at"`   // 更新时间戳
	Status      int8        `json:"status"  `      // 44 删除, 1 启用, 4 禁用
	LanguageID  int         `json:"languageID"  `  // 语言ID:1中文,2英语
	Title       string      `json:"title"  `       //标题
	Description string      `json:"description"  ` //描述
	Extended    interface{} `json:"extended" `     // 扩展的
	MenuID      int         `json:"menuID"  `
	SiteID      int         `json:"siteID"  `
}

func GetModel(i *ReqModel) *page.Model {
	o := new(page.Model)
	o.ID = conv.StringToInt64(i.ID, 0)
	o.Title = i.Title
	o.SiteID = i.SiteID
	o.MenuID = i.MenuID
	o.Description = i.Description
	o.LanguageID = i.LanguageID
	o.Extended = conv.StructToJsonString(i.Extended)
	return o
}

func GetRespModel(i *page.Model) *RespModel {
	o := new(RespModel)
	o.Title = i.Title
	o.Status = i.Status
	o.ID = conv.Int64ToString(i.ID)
	o.Description = i.Description
	o.LanguageID = i.LanguageID
	o.Extended = conv.StringToInterface(i.Extended)
	o.SiteID = i.SiteID
	o.MenuID = i.MenuID
	return o
}
