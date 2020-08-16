// 网站数据表配置
package menu

// 菜单
type Model struct {
	ID         int64 `gorm:"column:id;primary_key" json:"id"`       // id
	CreatedAt  int64 `gorm:"column:created_at" json:"createdAt"`    // 创建时间戳
	UpdatedAt  int64 `gorm:"column:updated_at" json:"updatedAt"`    // 更新时间戳
	Status     int8  `json:"status" gorm:"column:status" `          // 44 删除, 1 启用, 4 禁用
	LanguageID int   `json:"languageID" gorm:"column:language_id" ` // 语言ID:1中文,2英语

	SiteID   int64  `json:"siteID" gorm:"column:site_id" `     // 网站ID
	Name     string `json:"name" gorm:"column:name" `          // 菜单名
	ParentID int64  `json:"parentID" gorm:"column:parent_id" ` // 父亲ID
	Level    int    `json:"level" gorm:"column:level" `        //级别
	Extended string `json:"extended" gorm:"column:extended" `  // 扩展的
}

func (sv *Model) TableName() string {
	return "ws_menu"
}
