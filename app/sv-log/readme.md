# 日志服务

提供http, grpc 服务；

[详情](../../doc/service/sv-log.md)

## run

go run app/sv-log/main.go
go run app/gateway-api/main.go

## test

```shell script

curl --header "tokenjwt: jsrsei**" \
  --request POST \
  --data '{
          "logID":1234,
          "msg":"msg--gin-123",
          "level":2
          }' \
  http://localhost:7001/sv-log/log

```
