package models

// APIResponse :
type APIResponse struct {
	Code    int         `json:"code" binding:"required"`
	Data    interface{} `json:"data"`
	Message string      `json:"message"`
}

type BaseModel struct {
	ID         int64 `gorm:"column:id;primary_key" json:"id"`       // id
	CreatedAt  int64 `gorm:"column:created_at" json:"createdAt"`    // 创建时间戳
	UpdatedAt  int64 `gorm:"column:updated_at" json:"updatedAt"`    // 更新时间戳
	Status     int8  `json:"status" gorm:"column:status" `          // 44 删除, 1 启用, 4 禁用
	LanguageID int   `json:"languageID" gorm:"column:language_id" ` // 语言ID:1中文,2英语
}

type ListParameter struct {
	OrderSt string
	WhereSt string
}
