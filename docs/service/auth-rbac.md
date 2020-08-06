# 鉴权

## 认证 (authentication) 和授权 (authorization) 的区别

- 你要登机，你需要出示你的 passport 和 ticket，passport 是为了证明你张三确实是你张三，这就是 authentication；
- 而机票是为了证明你张三确实买了票可以上飞机，这就是 authorization。

### 在 computer science 领域再举个例子：

- 你要登陆论坛，输入用户名张三，密码1234，密码正确，证明你张三确实是张三，这就是 authentication；
- 再一check用户张三是个版主，所以有权限加精删别人帖，这就是 authorization。

## 测试

```shell script
go run app/gateway-api/main.go
go run app/auth-rbac/main.go

curl http://localhost:7001/auth-rbac

#查看 userID=2 的权限
curl http://localhost:7001/auth-rbac/v1/user/2
# 查看 userID=2 是否有 box-add 标签页面的权限
curl http://localhost:7001/auth-rbac/v1/check/2?power=box-add
```

## 功能

- 用户的管理
- 角色管理
- 权限管理


