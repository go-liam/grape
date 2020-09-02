package site_cms

import (
	"github.com/go-liam/util/conv"
	"grape/internal/pkg/data/home_site/site"
)

type ReqModel struct {
	ID          string      ` json:"id"`           // id
	LanguageID  int         `json:"languageID"  `  // 语言ID:1中文,2英语
	Title       string      `json:"title"  `       //标题
	Email       string      `json:"email"  `       //邮箱
	Description string      `json:"description"  ` //描述
	Extended    interface{} `json:"extended" `     // 扩展的
}

type RespModel struct {
	ID          string      ` json:"id"`           // id
	CreatedAt   int64       ` json:"created_at"`   // 创建时间戳
	UpdatedAt   int64       ` json:"updated_at"`   // 更新时间戳
	Status      int8        `json:"status"  `      // 44 删除, 1 启用, 4 禁用
	LanguageID  int         `json:"languageID"  `  // 语言ID:1中文,2英语
	Title       string      `json:"title"  `       //标题
	Email       string      `json:"email"  `       //邮箱
	Description string      `json:"description"  ` //描述
	Extended    interface{} `json:"extended" `     // 扩展的
}

func GetModel(i *ReqModel) *site.Model {
	o := new(site.Model)
	o.ID = conv.StringToInt64(i.ID, 0)
	o.Title = i.Title
	//o.Status = 1
	o.Email = i.Email
	o.Description = i.Description
	o.LanguageID = i.LanguageID
	o.Extended = conv.StructToJsonString(i.Extended)
	return o
}

func GetRespModel(i *site.Model) *RespModel {
	o := new(RespModel)
	o.Title = i.Title
	o.Status = i.Status
	o.ID = conv.Int64ToString(i.ID)
	o.Email = i.Email
	o.Description = i.Description
	o.LanguageID = i.LanguageID
	o.Extended = conv.StringToInterface(i.Extended)
	return o
}
