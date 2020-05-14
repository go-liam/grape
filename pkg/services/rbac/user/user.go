package user

import (
	"grape/pkg/services/rbac/role"
	"grape/pkg/util/conv"
)

// UserInfo : 获取用户的所有权限 (要缓存)
func Info(userID int64) (*ModelUser,error) {
	info ,err1:= Server.FindOne(userID)
	if info == nil {
		return info,err1
	}
	if info.RoleFlag == ConstRoleRoot {
		return info,nil
	}
	roleIDs := conv.StringToIntArray(info.RoleIDs)
	if roleIDs == nil || len(roleIDs) <=0 {
		return info,nil
	}
	// role-> powers
	lsRole,_ := role.Server.List()
	info.PowerIDs = role.ListPowerIDsByIDs(roleIDs,lsRole)
	return info,nil
}

