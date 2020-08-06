package power

type ModelPower struct {
	PowerID    int    `json:"powerID"`
	PowerName  string `json:"powerName"`
	Status     int8   `json:"status"` // default0 ;  44 delete
	UpdateTime int64  `json:"updateTime"`
	CreateTime int64  `json:"createTime"`
	PowerTag   string `json:"powerTag"` //`权限标签`,唯一，eg: box-edit
	Type       int8   `json:"type"`     //产品 类型, 默认 1
}
