package region

type Model struct {
	UserID    int64 `gorm:"column:user_id" json:"userID"`        // 用户ID
	CreatedAt int64 `gorm:"column:created_at" json:"created_at"` // 时间戳
	//UpdatedAt  int64 `gorm:"column:updated_at" json:"updated_at"` //
	RegionID int `gorm:"column:region_id" json:"regionID"` // 区域ID
}

func (sv *Model) TableName() string {
	return "ar_user_region"
}
