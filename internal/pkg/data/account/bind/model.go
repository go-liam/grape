package bind

type Model struct {
	ID        int64 `gorm:"column:id;primary_key" json:"id"`    //log id
	CreatedAt int64 `gorm:"column:created_at" json:"createdAt"` // 创建时间戳
	UserID    int64 `gorm:"column:user_id" json:"userID"`       // uid
	//类型：1：gzhOpenID，2: wxUnionID，3：weappOpenId ，4：手机 , 10: 账号 11:密码, 12 邮箱 ;
	Type int8   `gorm:"column:type"  json:"type"`
	Name string `gorm:"column:name"  json:"name"` // 绑定名称
}

func (sv *Model) TableName() string {
	return "uc_user_bind"
}
