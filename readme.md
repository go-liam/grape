# Go边缘微服集群架构

> - 用了1年（2019年）时间思考，怎么做出【适用、简单、演化】的微服架构，【可读、可维护、可扩展】的微服。
通向罗马的道路千万条，参考了大量大牛的思想，我选择了一条，进行尝试，演化。
为了快速演示，架构采用了原型方式，mock数据，减少复杂度。程序员喜欢造轮子，忍不住也想造个轮子。
> - 随着数据量的增大，单一数据中心维护的成本成指数增长，架构的复杂度，开发的团队人数也增长迅猛；
有没一种方式，既能可扩容支持大数据，开发、维护成本尽可能没那么高？有，就是边缘微服集群。这个项目就是为它而生。
> - 企业的组织架构影响技术架构。此微服架构适合10人以内的开发团队。项目的产出是一个个Docker镜像。
> - 通过此项目练手，加深对微服架构的理解。

- [一、项目架构](#一项目架构)
  - [1、项目初衷](#1项目初衷)
  - [2、架构图](#2架构图)
  - [3、项目技术栈](#3项目技术栈)
    
- [二、服务监控](#二服务监控)
  - [1、prometheus](#1prometheus)
  - [2、grafana](#2grafana)
  - [3、全链路跟踪](#3全链路跟踪)
  
- [三、发布CI/CD](#三发布cicd)
  - [1、gitlab CI/CD](#1gitlab-cicd)
  
- [四、测试](四测试)
  - [1、本地开发](#1本地开发)
  - [2、docker-compose环境](#2docker-compose环境)
  - [3、minikube环境](#3minikube环境)
  - [4、阿里云Kubernetes环境](#4阿里云Kubernetes环境)
  - [5、Istio环境](#5istio环境)
  
- [五、代码开发](#五代码开发)
  - [1、Monorepo代码管理方式](#1Monorepo代码管理方式)
  - [2、代码规范](#2代码规范)
  - [3、代码测试](#3代码测试)
  - [4、性能要求](#4性能要求)
  - [5、代码编译构建工具Bazel](#5代码编译构建工具bazel)
  
- [六、安全](六安全)
  - [1、数据库存储](#1数据库存储)
  - [2、用户敏感数据](#2用户敏感数据) 
  - [3、运营数据](#2运营数据) 
  
- [七、其它可供参考微服务案例项目](#七其它可供参考微服务案例项目)

- [八、更新日志](#八更新日志)

## 一、项目架构

### 1、项目初衷

弄一个未来的微服架构。同时也符合未来的serverless。

### 2、架构图

[详情:doc/项目架构.md](doc/项目架构.md)

#### 区域信息(sv-region) 

能定位用户是属于哪个区域服务集群。含各区域全量数据。

#### 认证服务([auth-jwt](./doc/service/auth-jwt.md))

用户通过各种方式登录后，付给它Token令牌去访问资源，支持分端单点登录。

#### 鉴权服务([auth-rbac](./doc/service/auth-rbac.md))

正规的rbac鉴权。分角色，用户，权限。

#### 网关(gateway-api)

统一对用户身份认证处理。统一入口。

#### 配置中心(sv-config)

提供配置的数据接口；

#### 日志服务(sv-log)

记录日志，统一记录；

#### AI相关接口服务器(sv-ai)

#### 钱包(sv-money)

#### 用户中心(sv-user)

#### 通知服务(sv-notice)

#### MQ服务(sv-mq)

#### OSS文件管理(sv-oss)

#### 微信/小程序接口(api-wx)

#### APP接口(api-app)

#### CMS接口(api-cms)

#### 网站接口(api-www)

### 3、项目技术栈

golang,mysql,redis,mongo,jwt,docker,
kubernetes, istio, mq, devOps 。

## 二、服务监控

### 1、prometheus

[prometheus](doc/subject/prometheus.md)

### 2、grafana

[grafana](doc/subject/grafana.md)

### 3、全链路跟踪

[全链路跟踪](./doc/全链路跟踪.md)

## 三、发布CI/CD

### 1、gitlab CI/CD

[ci-cd文档](./doc/ci-cd.md)

## 四、测试

### 1、本地开发

```shell script
## 开发的go版本
go1.13.3
## 入口服务
go run app/gateway-api/main.go
curl http://localhost:7001
# host配置
# 查看 doc/host.md 
```

### 2、docker-compose环境

```shell script
cd deploy
./deploy-docker.sh
# 详情配置
deploy/docker/mn/docker-compose.yml
```

### 3、minikube环境

```shell script
## 先要把当前的k8s config 设置为minikube 。
cd deploy
./deploy-minikube.sh
```

### 4、阿里云Kubernetes环境

### 5、Istio环境

## 五、代码开发

### 1、Monorepo代码管理方式
Monorepo 是一种单库存储代码的方式。

[monorepo详细](doc/subject/monorepo.md)

### 2、代码规范

[代码规范详细](./doc/代码规范.md)

### 3、代码测试

[代码测试](./doc/代码测试.md)

### 4、性能要求

[性能要求](./doc/性能要求.md)

### 5、代码编译构建工具Bazel

Bazel 是一个很厉害的开源构建工具，我还没有领略它的奥秘，还在学习中。现在还是通过shell 脚本构建。

[查看doc/Bazel.md文档](doc/subject/bazel.md)

## 六、安全

[详情](./doc/安全.md)

### 1、数据库存储

### 2、用户敏感数据

### 3、运营数据

## 七、可供参考微服务案例项目

- [eShopOnContainers 微软支持](https://github.com/dotnet-architecture/eShopOnContainers)
- [microservices-demo 谷歌支持](https://github.com/GoogleCloudPlatform/microservices-demo)
- [杨波《Kubernetes 云原生微服务实践》](https://github.com/mdemo19/staffjoy)
- [StaffjoyV2](https://github.com/mdemo19/StaffjoyV2)

## 八、更新日志

- 2020-5-3 编写文档目录

[更多日志](./doc/更新日志.md)

## 九、异常/注意

- 要配置自己的docker账号（deploy/base/MakeDockerConfig.mk）
