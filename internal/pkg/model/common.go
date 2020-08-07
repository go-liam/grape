package models

// APIResponse :
type APIResponse struct {
	Code    int         `json:"code" binding:"required"`
	Data    interface{} `json:"data"`
	Message string      `json:"message"`
}

type BaseModel struct {
	ID         int64 `gorm:"column:id;primary_key" json:"id"`
	CreatedAt  int64 `gorm:"column:created_at" json:"created_at"`
	UpdatedAt  int64 `gorm:"column:updated_at" json:"updated_at"`
	Status     int8  `json:"status" gorm:"column:status" `          // 44 删除, 1 启用, 4 禁用
	LanguageID int   `json:"languageID" gorm:"column:language_id" ` // 语言ID:1中文,2英语
}
