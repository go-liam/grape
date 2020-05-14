# 区域选择算法

## 新注册用户/未登录用户

可以根据用户IP，分配给他区域ID与域名；

## 登录用户

根据userID 查找所在区域，返回区域ID，与域名；

## 接口测试

[参考sv-region.html](../../frontend/sv-region.html)

## test

```shell script
go run app/gateway-api/main.go
go run app/sv-region/main.go

curl http://localhost:7001/sv-region
```

## 用户ID区域数据表
