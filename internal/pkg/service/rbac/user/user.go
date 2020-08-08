package user

import (
	"github.com/go-liam/util/conv"
	"grape/internal/pkg/service/rbac/role"
)

// UserInfo : 获取用户的所有权限 (要缓存)
func Info(userID int64) (*ModelUser, error) {
	info, err1 := Server.FindOne(userID)
	if info == nil {
		return info, err1
	}
	if info.RoleFlag == ConstRoleRoot {
		return info, nil
	}
	roleIDs := conv.StringToInt64Array(info.RoleIDs)
	if roleIDs == nil || len(roleIDs) <= 0 {
		return info, nil
	}
	// role-> powers
	lsRole, _ := role.Server.List()
	info.PowerIDs = role.ListPowerIDsByIDs(roleIDs, lsRole)
	return info, nil
}
