# istio

## minikube

```shell script
# 切换环境
make deploy-test-config
```

### 安装istio 1.5

https://istio.io/docs/setup/getting-started/

```shell script
curl -L https://istio.io/downloadIstio | sh -
cd istio-1.5.2
cd bin
# 学习版本
./istioctl manifest apply --set profile=demo
# 简单版本
$ istioctl manifest apply 
```
使用：
$ kubectl label namespace default istio-injection=enabled

查看：
```shell script
minikube dashboard
# 访问
http://127.0.0.1:57576/api/v1/namespaces/kubernetes-dashboard/services/http:kubernetes-dashboard:/proxy/
```

### Uninstall istio

cd /Users/liampro/Downloads/temp/tool/istio/istio-1.5.2/bin
./istioctl manifest generate --set profile=demo | kubectl delete -f -

## example

kubectl delete namespace dev-liam
kubectl create namespace dev-liam

```shell script
kubectl label namespace dev-liam istio-injection=enabled
cd /Users/liampro/Downloads/temp/tool/istio/istio-1.5.2
kubectl apply -f samples/bookinfo/platform/kube/bookinfo.yaml -n dev-liam

kubectl get services -n dev-liam
kubectl get pods -n dev-liam

kubectl exec -it $(kubectl get pod -n dev-liam -l app=ratings -o jsonpath='{.items[0].metadata.name}') -n dev-liam -c ratings -- curl productpage:9080/productpage | grep -o "<title>.*</title>" 
# or
kubectl exec -it ratings-v1-6c9dbf6b45-kbmf4 -n dev-liam -c ratings -- curl productpage:9080/productpage | grep -o "<title>.*</title>" 
# gateway-api
kubectl apply -f samples/bookinfo/networking/bookinfo-gateway.yaml -n dev-liam
kubectl get gateway -n dev-liam
kubectl delete gateway bookinfo-gateway  -n dev-liam

# port
$ export INGRESS_PORT=$(kubectl -n istio-system get service istio-ingressgateway -o jsonpath='{.spec.ports[?(@.name=="http2")].nodePort}')
$ export SECURE_INGRESS_PORT=$(kubectl -n istio-system get service istio-ingressgateway -o jsonpath='{.spec.ports[?(@.name=="https")].nodePort}')

echo $INGRESS_PORT
30430
echo $SECURE_INGRESS_PORT
31938
export INGRESS_HOST=$(minikube ip)
echo $INGRESS_HOST
192.168.99.114

$ minikube tunnel
#打印
Status: 
        machine: minikube
        pid: 86628
        route: 10.96.0.0/12 -> 192.168.99.114
        minikube: Running
        services: [istio-ingressgateway]

$ export GATEWAY_URL=$INGRESS_HOST:$INGRESS_PORT
echo $GATEWAY_URL
192.168.99.114:30430
echo http://$GATEWAY_URL/productpage
http://192.168.99.114:30430/productpage

./bin/istioctl dashboard kiali
Access the Kiali dashboard. The default user name is admin and default password is admin.

kubectl get svc istio-ingressgateway -n istio-system

```

liam18@liam istio-1.5.2 % minikube ip
192.168.99.114



### 账号

./bin/istioctl dashboard kiali
Access the Kiali dashboard. The default user name is admin and default password is admin.

### minikube 实战

1. gateway 搭建
设置本地的gateway
/Users/liampro/Downloads/mblock/github/k8s-yaml/aliyun/mb-sz-test/gateway/deploy.sh

2。部署例子



