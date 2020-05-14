package role

type ModelRole struct {
	RoleID   int    `json:"roleID"`
	RoleName string `json:"roleName"`
	Status   int8   `json:"status"`    // default0 ;  44 delete
	PowerIDS string `json:"powerIDS"` // 权限 集合
}
