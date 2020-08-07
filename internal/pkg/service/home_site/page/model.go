// 网站数据表配置
package page

import models "grape/internal/pkg/model"

// 页面
type Model struct {
	models.BaseModel
	SiteID      int    `json:"siteID" gorm:"column:site_id" ` // 网站ID
	MenuID      int    `json:"menuID" gorm:"column:menu_id" ` // 菜单ID
	Title       string `json:"title" gorm:"column:title" `
	Author      string `json:"author" gorm:"column:author" `
	Description string `json:"description" gorm:"column:description" `
	Extended    string `json:"extended" gorm:"column:extended" ` // 扩展的
	Content     string `json:"content" gorm:"column:content" `   // 内容
}

func (sv *Model) TableName() string {
	return "ws_page"
}
