package role

type Model struct {
	ID        int64  `json:"id" gorm:"column:id;primary_key" `   //ID
	CreatedAt int64  `gorm:"column:created_at" json:"createdAt"` // 创建时间戳
	UpdatedAt int64  `gorm:"column:updated_at" json:"updatedAt"` // 更新时间戳
	Status    int8   `json:"status" gorm:"column:status"`        // 44 删除, 1 启用, 4 禁用
	Name      string `json:"name" gorm:"column:name"`            //名称
	PowerIDS  string `json:"powerIDS" gorm:"column:power_ids" `  // 权限 集合
	Extended  string `json:"extended" gorm:"column:extended" `   // 扩展的
}

func (sv *Model) TableName() string {
	return "rb_role"
}
