package rbac

import (
	"grape/pkg/services/rbac/role"
	"grape/pkg/services/rbac/user"
	"grape/pkg/util/conv"
)

// UserInfo : 获取用户的所有权限 (要缓存)
func UserInfo(userID int64) (*user.ModelUser,error) {
	info ,err1:= user.Server.FindOne(userID)
	if info == nil {
		return info,err1
	}
	if info.RoleFlag == 1{
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

