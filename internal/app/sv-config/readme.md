# 配置服务器

提供http, grpc 服务；

[详情](../../doc/service/sv-config.md)

## run

```shell script

go run app/sv-config/main.go
go run app/gateway-api/main.go

curl http://localhost:7001/sv-config
curl http://localhost:7001/sv-config/v1/client/list
curl http://localhost:7001/sv-config/v1/client/info/sv-ai

curl http://localhost:7001/sv-config/v1/cms/list
curl http://localhost:7001/sv-config/v1/cms/info/sv-ai

```

