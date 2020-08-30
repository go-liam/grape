// 网站数据表配置
package site

// ConfigModel :网站配置
type Model struct {
	ID          int64  `gorm:"column:id;primary_key" json:"id"`        // id
	CreatedAt   int64  `gorm:"column:created_at" json:"createdAt"`     // 创建时间戳
	UpdatedAt   int64  `gorm:"column:updated_at" json:"updatedAt"`     // 更新时间戳
	Status      int8   `json:"status" gorm:"column:status" `           // 状态 44 删除, 1 启用, 4 禁用
	LanguageID  int    `json:"languageID" gorm:"column:language_id" `  // 语言ID:1中文,2英语
	Title       string `json:"title" gorm:"column:title" `             //标题
	Email       string `json:"email" gorm:"column:email" `             //邮箱
	Description string `json:"description" gorm:"column:description" ` //描述
	Extended    string `json:"extended" gorm:"column:extended" `       // 扩展的
}

func (sv *Model) TableName() string {
	return "ws_site"
}
