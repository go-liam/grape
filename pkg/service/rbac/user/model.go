package user

const (
	ConstRoleRoot = 1 // 超级管理员
)

type ModelUser struct {
	UserID   int64  `json:"userID"`
	Name     string `json:"name"`     // 用户名
	PassWord string `json:"passWord"` // 密码
	RoleIDs  string `json:"roleIDs"`  // ids 角色IDS
	Status   int8   `json:"status"`   // default 0 ;  44 delete
	RoleFlag int8   `json:"roleFlag"` // 默认 0， 1 超级 管理员
	Remark   string `json:"remark"`   // 备注
	PowerIDs []int  `json:"powerIDs"` // 计算合并的权限(不是数据表)
}
