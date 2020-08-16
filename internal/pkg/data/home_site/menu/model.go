// 网站数据表配置
package menu

import models "grape/internal/pkg/model"

// 菜单
type Model struct {
	models.BaseModel
	SiteID   int64  `json:"siteID" gorm:"column:site_id" `     // 网站ID
	Name     string `json:"name" gorm:"column:name" `          // 菜单名
	ParentID int64  `json:"parentID" gorm:"column:parent_id" ` // 父亲ID
	Level    int    `json:"level" gorm:"column:level" `        //级别
	Extended string `json:"extended" gorm:"column:extended" `  // 扩展的
}

func (sv *Model) TableName() string {
	return "ws_menu"
}
