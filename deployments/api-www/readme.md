# info

```shell script
cd deploy/api-www
make get-config
make test-unit
make build-file
make build-image
make deploy-local-apply

```

## 发布

```shell script
# 设置权限
ls -l deploy.sh
chmod +x deploy.sh

# 发布
./deploy.sh

```

## minikube
```shell script
make deploy-test-config
kubectl get node
kubectl get namespace
kubectl get svc  -n dev-qps
kubectl delete svc api-www -n dev-qps
kubectl get pod  -n dev-qps

make deploy-test-apply
```

view:
```shell script
kubectl get svc  -n dev-qps
```

- IP端口是查出来的，不固定
- curl http://192.168.99.114:31481


## ali k8s

```shell script
make deploy-pre-config
kubectl get node
kubectl get namespace
kubectl get svc  -n dev-liam
kubectl get svc  -n dev-cbox

make deploy-pre-apply

kubectl apply -f ./ali-k8s-liam  -n dev-liam
kubectl delete svc api-www -n dev-liam

kubectl port-forward api-www-v1-7b9db584dd-8jkmr 8044:7041 -n dev-liam

```

## minikube-istio

kubectl delete namespace dev-ist
kubectl create namespace dev-ist 
kubectl label namespace dev-ist istio-injection=enabled
kubectl apply -f ./minikube-istio  -n dev-ist


## minikube-k8s

kubectl delete namespace dev-k8s
kubectl create namespace dev-k8s
kubectl apply -f ./minikube  -n dev-k8s

## ali istio

- 换环境
make deploy-pre-config

kubectl create namespace dev-liam
kubectl delete namespace dev-liam

kubectl label namespace dev-liam istio-injection=enabled

kubectl apply -f ./ali-k8s-istio  -n dev-liam

http://liam-test.mygs.com/api-www/ip
