package sylog

type Model struct {
	ID        int64  `gorm:"column:id;primary_key" json:"id"`  // id
	CreatedAt int64  `gorm:"column:created_at" json:"created"` // 创建时间戳
	Level     int32  `gorm:"column:level"  json:"level"`       // 级别
	Extended  string `json:"extended" gorm:"column:extended" ` // 扩展的
	Msg       string ` gorm:"column:msg" json:"msg"`           // 信息
}

func (sv *Model) TableName() string {
	return "lg_log"
}

type ConfigModel struct {
	URL string
}
