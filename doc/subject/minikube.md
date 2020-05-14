# info
https://kubernetes.io/zh/docs/setup/learning-environment/minikube/
https://kubernetes.io/

## k8s-minikube处理

安装

```shell script
# 开翻墙
brew install minikube
minikube status
minikube stop
minikube start --vm-driver=virtualbox
minikube delete
```

安装dashboard

```shell script
minikube dashboard
# 访问
http://127.0.0.1:57576/api/v1/namespaces/kubernetes-dashboard/services/http:kubernetes-dashboard:/proxy/
```

部署应用

```shell script
cd deploy
# 切换环境
make deploy-test-config
make deploy-test-apply
# 创建空间名词
make deploy-dev-create
kubectl create namespace test-mn
kubectl delete namespace test-mn
# 部署
make deploy-dev-apply
```

检测

```shell script
kubectl get node
kubectl get svc  -n test-mn
kubectl get svc -l app=my-qps1k  -n dev-qps
kubectl get po -n dev-qps



```

检测部分结果

``` log
qps1k % kubectl get node
NAME       STATUS   ROLES    AGE   VERSION
minikube   Ready    master   31d   v1.17.0

qps1k % kubectl get svc  -n dev-qps
NAME        TYPE       CLUSTER-IP     EXTERNAL-IP   PORT(S)        AGE
my-qps1k   NodePort   10.96.36.121   <none>        80:32722/TCP   14s
```

测试

- IP端口是查出来的，不固定
- curl http://192.168.99.114:31481

清理

```shell script
make deploy-dev-delete
```

仪表板: minikube dashboard
取IP： minikube ip

进入服务器 minikube 主机：
minikube ssh 

查看Container的日志：
kubectl get pod -n test-mn
kubectl logs gateway-api-v1-fd489c7f9-5dgcm -n test-mn

kubectl get svc -n test-mn
kubectl delete svc gateway-api -n test-mn

kubectl port-forward gateway-api-v1-fd489c7f9-5dgcm  8044:7001 -n dev-canary

dashboard
kubectl apply -f https://raw.githubusercontent.com/kubernetes/dashboard/v2.0.0-beta6/aio/deploy/recommended.yaml
minikube dashboard

## 实例子 k8s

- TLS证书
```shell script
kubectl create secret tls mcom-secret --namespace=test-secret --cert=public/cert/m-com-cert.pem --key=public/cert/m-com-key.pem
kubectl create secret tls mcom-secret --namespace=dev-qps --cert=public/cert/m-com-cert.pem --key=public/cert/m-com-key.pem
kubectl describe secret mcom-secret --namespace=test-secret
```
- ingress
https://kubernetes.io/docs/concepts/services-networking/ingress/
kubectl apply -f ./minikube  -n dev-qps 

## pod 之间无法访问问题

- 检查日志，是否是新版本？
- 检查端口
- 是否有缓存？ 

## 常用命令
```shell script
# 查看集群信息
kubectl cluster-info

kubectl get namespace

kubectl get pods --all-namespaces 
kubectl get svc --all-namespaces 

kubectl logs gateway-api-v1-fd489c7f9-5dgcm -n test-mn
kubectl port-forward gateway-api-v1-fd489c7f9-5dgcm  8044:7001 -n dev-canary

```

### 部分的日志

```
liam18@liam grape % kubectl cluster-info
Kubernetes master is running at https://192.168.99.114:8443
KubeDNS is running at https://192.168.99.114:8443/api/v1/namespaces/kube-system/services/kube-dns:dns/proxy

192.168.99.114 就是本地绑定的ip
```

