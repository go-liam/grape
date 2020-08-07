// 网站数据表配置
package site

import models "grape/internal/pkg/model"

// ConfigModel :网站配置
type Model struct {
	models.BaseModel
	Title       string `json:"title" gorm:"column:title" `
	Email       string `json:"email" gorm:"column:email" `
	Description string `json:"description" gorm:"column:description" `
	Extended    string `json:"extended" gorm:"column:extended" ` // 扩展的
}

func (sv *Model) TableName() string {
	return "ws_site"
}
