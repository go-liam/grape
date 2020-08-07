// 网站数据表配置
package notice

import models "grape/internal/pkg/model"

// 通知
type Model struct {
	models.BaseModel
	SiteID   int    `json:"siteID" gorm:"column:site_id" ` // 网站ID
	Title    string `json:"title" gorm:"column:title" `
	Author   string `json:"author" gorm:"column:author" `
	Extended string `json:"extended" gorm:"column:extended" ` // 扩展的
	Content  string `json:"content" gorm:"column:content" `   // 内容
}

func (sv *Model) TableName() string {
	return "ws_notice"
}
