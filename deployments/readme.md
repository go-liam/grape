# 发布目录

```shell script
cd deploy
make test-unit

# 单元测试

export PROJECT_ENV="unit" && go test $(go list ./... | grep -v /proto/ | grep -v /temp-demo/ | grep -v /frontend/ ) -coverprofile=coverage.data ./...;

# docker-compose
make deploy-dev-apply
make delete-dev-apply

```

## 注意

MakeDockerConfig.mk 没有发布到git, 
你可以把 base/MakeDockerConfig.mk.example 改为 MakeDockerConfig.mk

## 运行环境

PROJECT_ENV:
- unit 单元测试环境
- 本机开发 local 本地开发
- docker-compose dev 联合开发环境
- minikube test 内部测试环境
- ali k8s pre 预发布环境
- ali k8s prod  正式环境

## docker-compose中 调试

常用命令
```shell script
    # cd docker-compose.yml filePath
    # docker-compose up
    # docker-compose down
    # docker ps
    # docker container exec -it mn_gateway-api_1 /bin/sh
    # cd log
    # ls
    # cat d-2019-09-24-2.log -n
    
docker logs --since 30m mn_gateway-api_1
```

### 更改了没变化

docker-compose.yal 中，改下环境变量;
TEMP=test53 环境变量值设置下；
