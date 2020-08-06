# rbac 鉴权

- 用户管理
- 角色管理
- 权限管理

![路径](./doc/rs/rbac.webp)


## run

go run app/auth-rbac/main.go

## test
- 查看 userID=2 信息
curl /auth-rbac/v1/user/2 

- 查看 userID=2 是否有 "box-add" 权限
curl /auth-rbac/v1/check/2?power=box-add
