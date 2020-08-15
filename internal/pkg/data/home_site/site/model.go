// 网站数据表配置
package site

import models "grape/internal/pkg/model"

// ConfigModel :网站配置
type Model struct {
	models.BaseModel
	Title       string `json:"title" gorm:"column:title" `             //标题
	Email       string `json:"email" gorm:"column:email" `             //邮箱
	Description string `json:"description" gorm:"column:description" ` //描述
	Extended    string `json:"extended" gorm:"column:extended" `       // 扩展的
}

func (sv *Model) TableName() string {
	return "ws_site"
}
