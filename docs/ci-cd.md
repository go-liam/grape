# 发布ci/cd

## 说明

- 部署目录：deploy
- 环境分为 docker-compose, minikube , ali k8s 
- 发布：cd deploy/**/deploy

## 新项目

```
copy deploy/api-www 为 xxx
#1.整个文件夹批量更新
#2.修改项目名：
api-www -> xxx
#3.改端口: 
7041 --> 新端口
#4. docker-compose
加上新项目
#5. docker-compose
pkg/config 加上配置
```

## 快速部署集群

### docker-compose

```shell script
cd deploy/docker
docker-compose up -d
docker-compose down

curl http://localhost:6001
```
[调试](../deploy/docker/readme.md)

### minikube

```shell script

#config
cd ~/.kube 
cp -p config-minikube ./config 
# namespace
kubectl delete namespace test-mn
kubectl create namespace test-mn
# deploy 
kubectl apply -f ./deploy/gateway-api/minikube  -n test-mn
kubectl apply -f ./deploy/api-www/minikube  -n test-mn
kubectl apply -f ./deploy/api-app/minikube  -n test-mn
#... 其它..pod

## 快速
./deploy/deploy-minikube.sh
# 获取端口
kubectl get svc  -n test-mn
```
[调试](subject/minikube.md)

```log
liam18@liam grape % kubectl create namespace test-mn
liam18@liam grape % kubectl apply -f ./deploy/gateway-api/minikube  -n test-mn
liam18@liam grape % kubectl get svc  -n test-mn                               
NAME             TYPE       CLUSTER-IP      EXTERNAL-IP   PORT(S)          AGE
s1-gateway-api   NodePort   10.96.109.252   <none>        7001:32449/TCP   2m19s
liam18@liam grape % curl http://192.168.99.114:32449                          
3-index,ip=172.17.0.1%    
```

### k8s
