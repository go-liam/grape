# 用户登录权限-jwt认证

- 通过jwt,检查用户是否登录
- 可以自动刷新时间;
- 生成更新相关的 cookies;
- 有黑名单

## run

go run app/auth-jwt/main.go

## single 单点登录

- 给用户使用的
- redis 缓存
- 按客户端单点登录

## sv

- 服务器间的使用
- 直接使用 jwt 机制，不用redis;