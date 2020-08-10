package config_sql

type Model struct {
	ID        int64  `gorm:"column:id;primary_key" json:"id"`     // id
	CreatedAt int64  `gorm:"column:created_at" json:"created_at"` // 创建时间戳
	UpdatedAt int64  `gorm:"column:updated_at" json:"updated_at"` // 更新时间戳
	Status    int8   `json:"status" gorm:"column:status" `        // 44 删除, 1 启用, 4 禁用
	Version   int    `gorm:"column:version"  json:"version"`      //版本
	Content   string `gorm:"column:content"  json:"content"`      // 内容
	Extended  string `json:"extended" gorm:"column:extended" `    // 扩展的
}

func (sv *Model) TableName() string {
	return "cf_configs"
}
