# 区域服务


## 功能

- 注册

```shell script
# local
curl http://localhost:7409/sv-region/v1/reg?ip=127.0.0.1
# test
curl http://localhost:7409/sv-region/v1/reg?ip=192.168.0.1
# default
curl http://localhost:7409/sv-region/v1/reg?ip=123.456.0.1
```

- 获取

```shell script
#local
curl http://localhost:7409/sv-region/v1/info/10
#test
curl http://localhost:7409/sv-region/v1/info/100
#default
curl http://localhost:7409/sv-region/v1/info/1000

```
