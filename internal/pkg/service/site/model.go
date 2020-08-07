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

// 菜单
type MenuModel struct {
	models.BaseModel
	SiteID   int    `json:"siteID" gorm:"column:site_id" `     // 网站ID
	Name     string `json:"name" gorm:"column:name" `          // 菜单名
	ParentID int    `json:"parentID" gorm:"column:parent_id" ` // 父亲ID
	Level    int    `json:"level" gorm:"column:level" `        //级别
}

func (sv *MenuModel) TableName() string {
	return "ws_menu"
}

// 页面
type PageModel struct {
	models.BaseModel
	SiteID      int    `json:"siteID" gorm:"column:site_id" ` // 网站ID
	MenuID      int    `json:"menuID" gorm:"column:menu_id" ` // 菜单ID
	Title       string `json:"title" gorm:"column:title" `
	Author      string `json:"author" gorm:"column:author" `
	Description string `json:"description" gorm:"column:description" `
	Extended    string `json:"extended" gorm:"column:extended" ` // 扩展的
	Content     string `json:"content" gorm:"column:content" `   // 内容
}

func (sv *PageModel) TableName() string {
	return "ws_page"
}

// 通知
type NoticeModel struct {
	models.BaseModel
	SiteID   int    `json:"siteID" gorm:"column:site_id" ` // 网站ID
	Title    string `json:"title" gorm:"column:title" `
	Author   string `json:"author" gorm:"column:author" `
	Extended string `json:"extended" gorm:"column:extended" ` // 扩展的
	Content  string `json:"content" gorm:"column:content" `   // 内容
}

func (sv *NoticeModel) TableName() string {
	return "ws_notice"
}
