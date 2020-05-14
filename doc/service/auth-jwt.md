# 认证

## 认证 (authentication) 和授权 (authorization) 的区别

- 你要登机，你需要出示你的 passport 和 ticket，passport 是为了证明你张三确实是你张三，这就是 authentication；
- 而机票是为了证明你张三确实买了票可以上飞机，这就是 authorization。

## 服务间简单的JWT的认证

- 登录
- 退出
- 刷新 token

### 演示

[auth-jwt-sv.html](../../frontend/auth-jwt-sv.html)

## 网站用户的JWT的认证

- 分端单点登录
- 刷新 token
- 退出当前端
- 有退出所有端接口（修改密码后）

### 演示

[auth-jwt.html](../../frontend/auth-jwt.html)

## test

```shell script
go run app/gateway-api/main.go
go run app/auth-jwt/main.go

```
用浏览器直接打开 auth-jwt.html 或 auth-jwt-sv.html
(不用node环境)
