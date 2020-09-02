package user

type Model struct {
	ID        int64  `gorm:"column:id;primary_key" json:"id"`     // id
	CreatedAt int64  `gorm:"column:created_at" json:"created_at"` // 创建时间戳
	UpdatedAt int64  `gorm:"column:updated_at" json:"updated_at"` // 更新时间戳
	Status    int8   `gorm:"column:status" json:"status" `        // 44 删除, 1 启用, 4 禁用
	Remark    string `gorm:"column:remark"  json:"remark"`        // 备注
	Extended  string `gorm:"column:extended" json:"extended"`     // 扩展的
	Email     string `gorm:"column:email" json:"email"`           // 邮箱（冗余）
	Phone     string `gorm:"column:phone" json:"phone"`           // 手机（冗余）
	Name      string `gorm:"column:name" json:"name"`             //账号（冗余）
	NickName  string `gorm:"column:nick_name" json:"nick_name"`   //昵称
}

func (sv *Model) TableName() string {
	return "uc_user"
}
