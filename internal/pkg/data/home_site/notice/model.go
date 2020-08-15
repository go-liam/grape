// 网站数据表配置
package notice

// 通知
type Model struct {
	ID         int64  `gorm:"column:id;primary_key" json:"id"`       // id
	CreatedAt  int64  `gorm:"column:created_at" json:"createdAt"`    // 创建时间戳
	UpdatedAt  int64  `gorm:"column:updated_at" json:"updatedAt"`    // 更新时间戳
	Status     int8   `json:"status" gorm:"column:status" `          // 44 删除, 1 启用, 4 禁用
	LanguageID int    `json:"languageID" gorm:"column:language_id" ` // 语言ID:1中文,2英语
	SiteID     int    `json:"siteID" gorm:"column:site_id" `         // 网站ID
	Title      string `json:"title" gorm:"column:title" `            // 标题
	Author     string `json:"author" gorm:"column:author" `          //作者
	Extended   string `json:"extended" gorm:"column:extended" `      // 扩展的
	Content    string `json:"content" gorm:"column:content" `        // 内容
}

func (sv *Model) TableName() string {
	return "ws_notice"
}
