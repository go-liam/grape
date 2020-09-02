package user

type Model struct {
	ID        int64  `gorm:"column:id;primary_key" json:"id"`     // id
	CreatedAt int64  `gorm:"column:created_at" json:"created_at"` // 创建时间戳
	UpdatedAt int64  `gorm:"column:updated_at" json:"updated_at"` // 更新时间戳
	Status    int8   `json:"status" gorm:"column:status" `        // 44 删除, 1 启用, 4 禁用
	Flag      int8   `gorm:"column:flag"  json:"flag"`            // 默认 0， 1 超级管理员
	Remark    string `gorm:"column:remark"  json:"remark"`        // 备注
	Extended  string `json:"extended" gorm:"column:extended" `    // 扩展的
	RoleIDs   string ` gorm:"column:role_ids" json:"roleIDs"`     // 角色IDs
}

func (sv *Model) TableName() string {
	return "rb_user"
}
