package bind

const (
	TypeUserName = 10
	TypeEmail    = 12
	TypePhone    = 4
)

const (
	KeyMd5 = "123@GciOiJIUzIInliamR5cCI6IkpXVCJ9#$%@jl"
)

type Model struct {
	ID        int64 `gorm:"column:id;primary_key" json:"id"`     //log id
	CreatedAt int64 `gorm:"column:created_at" json:"created_at"` // 创建时间戳
	UserID    int64 `gorm:"column:user_id" json:"userID"`        // uid
	//类型：1：gzhOpenID，2: wxUnionID，3：weappOpenId ，4：手机 , 10: 账号  12 邮箱 ;
	Type int8   `gorm:"column:type"  json:"type"`
	Name string `gorm:"column:name"  json:"name"` // 绑定名称
}

func (sv *Model) TableName() string {
	return "uc_user_bind"
}

type ModelPassword struct {
	ID        int64  `gorm:"column:id;primary_key" json:"id"`     //user id
	CreatedAt int64  `gorm:"column:created_at" json:"created_at"` // 创建时间戳
	UpdatedAt int64  `gorm:"column:updated_at" json:"updated_at"` // 更新时间戳
	Password  string `gorm:"column:password"  json:"password"`    // 密码
}

func (sv *ModelPassword) TableName() string {
	return "uc_user_psd"
}
