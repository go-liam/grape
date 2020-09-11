// 网站数据表配置
package server

const (
	// 服务器3种状态
	StatusOk    = 1
	StatusWarn  = 2
	StatusError = 3
)

type Model struct {
	ID          int64  `gorm:"column:id;primary_key" json:"id"`        // id
	CreatedAt   int64  `gorm:"column:created_at" json:"created_at"`    // 创建时间戳
	UpdatedAt   int64  `gorm:"column:updated_at" json:"updated_at"`    // 更新时间戳
	Status      int8   `json:"status" gorm:"column:status" `           // 状态 44 删除: 1 ok,2警告,3出错
	Title       string `json:"title" gorm:"column:title" `             //标题
	Description string `json:"description" gorm:"column:description" ` //描述
	Extended    string `json:"extended" gorm:"column:extended" `       // 扩展的
}

func (sv *Model) TableName() string {
	return "sa_server"
}
