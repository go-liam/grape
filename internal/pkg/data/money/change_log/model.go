package change_log

type Model struct {
	ID         int64  `gorm:"column:id;primary_key" json:"id"`      // id
	CreatedAt  int64  `gorm:"column:created_at" json:"created"`  // 创建时间戳
	UpdatedAt  int64  `gorm:"column:updated_at" json:"updated"`  // 更新时间戳
	Status     int8   `gorm:"column:status" json:"status" `         // 44 删除, 10 ok
	Remark     string `gorm:"column:remark"  json:"remark"`         // 备注
	Extended   string `gorm:"column:extended" json:"extended"`      // 扩展的
	UserID     int64  `gorm:"column:user_id" json:"userID"`         // uid
	Amount     int    `gorm:"column:amount" json:"amount"`          // 数额（分）人民币
	LastAmount int    `gorm:"column:last_amount" json:"lastAmount"` // 上次个人账户 数额
	Type  int8    `gorm:"column:type" json:"type"` // 支付类型｜充值类型
	OrderID  int64    `gorm:"column:order_id" json:"orderID"` // 订单ID
	MoneyType  int8    `gorm:"column:type" json:"type"` // 1 收入，2 消费
}

func (sv *Model) TableName() string {
	return "um_log"
}
