# gateway

- 对外的防火墙；
- 节省域名;
- 路由代理;
- 通用处理Token转UserID;
- 黑名单处理;

## test run

go run app/gateway-api/main.go
curl --request GET -H "Token: 1234-web-token" http://localhost:7001

go run app/api-cms/main.go

curl --request GET -H "Token: 1234-web-token" http://gateway-api-local.mygs.com:7001/api-cms/aa
curl --request GET -H "Token: 1234-web-token" http://192.168.31.235:7001/api-cms/aa
curl --request GET -H "Token: 1234-web-token" http://localhost:7001/api-cms/aa

- bad ip:
curl --request GET -H "X-Real-Ip:192.168.31.235" http://192.168.31.235:7001/

- 换IP
curl --request GET -H "X-Real-Ip:192.168.31.200" http://192.168.31.235:7001/
