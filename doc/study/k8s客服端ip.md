# 客户端IP

## 资源

https://www.ziji.work/kubernetes/kubernetes-externaltrafficpolicy-clientip.html

阿里云容器服务Ingress设置原IP透传:
https://yq.aliyun.com/articles/609162
"将下图对应的红色框的内容从Cluster改为Local"

## 实验1，直接用k8s

- host:
*.*.237.70  gateway-api-pre.mygs.com
备注： *.*.237.70  是 kube-system nginx-ingress-lb 绑定的ip;

- yaml关键点
  type: NodePort
  externalTrafficPolicy: Local
可以获取用户ip :
k8s: curl http://gateway-api-pre.mygs.com/api-www/ip
```json
{"IP":"113.88.13.21","Request":null,"Header":{"Accept":["text/html,application/xhtml+xml,application/xml;q=0.9,image/webp,image/apng,*/*;q=0.8,application/signed-exchange;v=b3;q=0.9"],"Accept-Encoding":["gzip, deflate"],"Accept-Language":["zh-CN,zh;q=0.9"],"Cache-Control":["max-age=0"],"Upgrade-Insecure-Requests":["1"],"User-Agent":["Mozilla/5.0 (Macintosh; Intel Mac OS X 10_15_3) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/80.0.3987.106 Safari/537.36"],"X-Forwarded-For":["113.88.13.21"],"X-Forwarded-Host":["gateway-api-pre.mygs.com"],"X-Forwarded-Port":["80"],"X-Forwarded-Proto":["http"],"X-Original-Uri":["/api-www/ip"],"X-Real-Ip":["113.88.13.21"],"X-Request-Id":["8011456cf7362867d58c29f7c19b1c72"],"X-Scheme":["http"]}}
```

- 内部访问：
liam18@liam api-www % kubectl get svc -n dev-cbox
NAME                TYPE        CLUSTER-IP      EXTERNAL-IP   PORT(S)          AGE
api-www             NodePort    172.21.5.43     <none>        7041:30556/TCP   41h
box2api             ClusterIP   172.21.4.167    <none>        80/TCP           56d
box2wx              ClusterIP   172.21.4.4      <none>        8080/TCP         56d
mbcontentplatform   ClusterIP   172.21.7.124    <none>        9080/TCP         39d
wechatservice       ClusterIP   172.21.14.206   <none>        3001/TCP         18d

liam18@liam api-www % kubectl get svc -n dev-liam
NAME      TYPE       CLUSTER-IP     EXTERNAL-IP   PORT(S)        AGE
api-www   NodePort   172.21.4.187   <none>        80:32429/TCP   34m

```shell script box2wx 去访问
bash-5.0# curl box2api
Hello,It works. a3194ca093ded51cf65df02d61a6568403508fd2bash-5.0# 
bash-5.0# curl api-www:7041
bash-5.0# curl api-www.dev-liam
Hello,It works.index bash-5.0# 
```
mbcontentplatform 去访问
```shell script
bash-5.0# curl box2api
Hello,It works. a3194ca093ded51cf65df02d61a6568403508fd2bash-5.0# 
bash-5.0# curl api-www:7041
Hello,It works.index bash-5.0# 
bash-5.0# curl api-www.dev-liam
Hello,It works.index bash-5.0#
```
api-www.dev-liam去访问
```shell script
bash-5.0# curl api-www.dev-cbox:7041
Hello,It works.index bash-5.0# 
bash-5.0# curl box2wx.dev-cbox:8080
{"code":0,"data":{"apiVersion":"bf7a8bc500a93f18893b98d42e893c838cc9a210","env":"dev"},"message":"OK"}
bash-5.0# 
```

下面环境不可以：
istio-injection: enabled : http://gateway-api-pre.mygs.com/api-www/ip

## 实验2，直接用k8s

- yaml关键点
  type: ClusterIP
  #externalTrafficPolicy: Local
可以获取用户ip :
k8s: curl http://gateway-api-pre.mygs.com/api-www/ip
```json
{"IP":"113.88.13.21","Request":null,"Header":{"Accept":["text/html,application/xhtml+xml,application/xml;q=0.9,image/webp,image/apng,*/*;q=0.8,application/signed-exchange;v=b3;q=0.9"],"Accept-Encoding":["gzip, deflate"],"Accept-Language":["zh-CN,zh;q=0.9"],"Cache-Control":["max-age=0"],"Upgrade-Insecure-Requests":["1"],"User-Agent":["Mozilla/5.0 (Macintosh; Intel Mac OS X 10_15_3) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/80.0.3987.106 Safari/537.36"],"X-Forwarded-For":["113.88.13.21"],"X-Forwarded-Host":["gateway-api-pre.mygs.com"],"X-Forwarded-Port":["80"],"X-Forwarded-Proto":["http"],"X-Original-Uri":["/api-www/ip"],"X-Real-Ip":["113.88.13.21"],"X-Request-Id":["dc470b0d599bfa6cdba02160927df975"],"X-Scheme":["http"]}}
```

liam18@liam api-www % kubectl get svc  -n dev-liam                
NAME      TYPE        CLUSTER-IP    EXTERNAL-IP   PORT(S)   AGE
api-www   ClusterIP   172.21.9.75   <none>        80/TCP    8s

小结：
在k8s 中，无论 type: ClusterIP ，NodePort 都可以获取客户端ip。问题出在 istio ;

## 实验3 istio 用 ali-k8s-liam

kubectl delete namespace dev-liam
kubectl create namespace dev-liam 
kubectl label namespace dev-liam istio-injection=enabled

kubectl apply -f ./ali-k8s-liam  -n dev-liam

http://gateway-api-pre.mygs.com/api-www/ip

```shell script
liam18@liam api-www % kubectl create namespace dev-liam 
namespace/dev-liam created
liam18@liam api-www % kubectl label namespace dev-liam istio-injection=enabled
namespace/dev-liam labeled
liam18@liam api-www % kubectl apply -f ./ali-k8s-liam  -n dev-liam
service/api-www created
deployment.apps/api-www-v1 created
ingress.extensions/api-www-ingress created
liam18@liam api-www % kubectl get svc  -n dev-liam
NAME      TYPE        CLUSTER-IP     EXTERNAL-IP   PORT(S)   AGE
api-www   ClusterIP   172.21.7.176   <none>        80/TCP    70s
```
json
```json
{"IP":"113.88.13.21","Request":null,"Header":{"Accept":["text/html,application/xhtml+xml,application/xml;q=0.9,image/webp,image/apng,*/*;q=0.8,application/signed-exchange;v=b3;q=0.9"],"Accept-Encoding":["gzip, deflate"],"Accept-Language":["zh-CN,zh;q=0.9"],"Cache-Control":["no-cache"],"Content-Length":["0"],"Pragma":["no-cache"],"Upgrade-Insecure-Requests":["1"],"User-Agent":["Mozilla/5.0 (Macintosh; Intel Mac OS X 10_15_3) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/80.0.3987.106 Safari/537.36"],"X-Forwarded-For":["113.88.13.21"],"X-Forwarded-Host":["gateway-api-pre.mygs.com"],"X-Forwarded-Port":["80"],"X-Forwarded-Proto":["http"],"X-Original-Uri":["/api-www/ip"],"X-Real-Ip":["113.88.13.21"],"X-Request-Id":["215a4a91c34874b474dd95c217b250ee"],"X-Scheme":["http"]}}
```
结果： 居然可以获取IP,因为用的是 Ingress，

## 实验4 用ali-k8s

- yaml相关
istio-ingressgateway：相关记录
ip: *.107.232.*
externalTrafficPolicy: Cluster

- 操作
kubectl delete namespace dev-liam
kubectl create namespace dev-liam 
kubectl label namespace dev-liam istio-injection=enabled

kubectl apply -f ./ali-k8s  -n dev-liam

http://gateway-api-pre.mygs.com/api-www/ip
```json
{"IP":"172.20.1.1","Request":null,"Header":{"Accept":["text/html,application/xhtml+xml,application/xml;q=0.9,image/webp,image/apng,*/*;q=0.8,application/signed-exchange;v=b3;q=0.9"],"Accept-Encoding":["gzip, deflate"],"Accept-Language":["zh-CN,zh;q=0.9"],"Cache-Control":["no-cache"],"Content-Length":["0"],"Pragma":["no-cache"],"Upgrade-Insecure-Requests":["1"],"User-Agent":["Mozilla/5.0 (Macintosh; Intel Mac OS X 10_15_3) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/80.0.3987.106 Safari/537.36"],"X-Envoy-Internal":["true"],"X-Forwarded-For":["172.20.1.1"],"X-Forwarded-Proto":["http"],"X-Request-Id":["4055bd81-7302-41e2-bba7-a68331d3782f"]}}
```

```shell script
liam18@liam api-www % kubectl get svc  -n dev-liam             
NAME      TYPE        CLUSTER-IP    EXTERNAL-IP   PORT(S)    AGE
api-www   ClusterIP   172.21.2.78   <none>        7041/TCP   79s

```

- 加一个gateway
cd deploy/gateway
kubectl apply -f ./ali-k8s-liam  -n dev-liam
http://gateway-api-pre.mygs.com/api-www/ip

```json
{"IP":"*.88.13.*","Request":null,"Header":{"Accept":["text/html,application/xhtml+xml,application/xml;q=0.9,image/webp,image/apng,*/*;q=0.8,application/signed-exchange;v=b3;q=0.9"],"Accept-Encoding":["gzip, deflate"],"Accept-Language":["zh-CN,zh;q=0.9"],"Cache-Control":["max-age=0"],"Content-Length":["0"],"Upgrade-Insecure-Requests":["1"],"User-Agent":["Mozilla/5.0 (Macintosh; Intel Mac OS X 10_15_3) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/80.0.3987.106 Safari/537.36"],"X-Forwarded-For":["*.88.13.*, 127.0.0.1"],"X-Forwarded-Host":[""],"X-Forwarded-Port":["80"],"X-Forwarded-Proto":["http"],"X-Original-Uri":["/api-www/ip"],"X-Real-Ip":["*.88.13.*"],"X-Request-Id":["866dfebfbf496722c042e8081422cd73"],"X-Scheme":["http"],"X-Test-Host":["*.123.0.1"]}}
```

方案G1: k8s :弄个gateway 转发到各个 server中；
http://gateway-api-pre.mygs.com/api-www/ip

方案2: istio
http://xring.info/2019/istio-envoy-get-client-real-ip.html
istio-ingressgateway 的 Service，类型为 LoadBalancer，可以将服务暴露到公网；
这个 Service 的 spec.externalTrafficPolicy 默认值是 Cluster，我们只要将这个值更新为 Local，就可以在 X-Forwarded-For 请求头获取到客户端真实 IP 地址了。

https://kubernetes.io/zh/docs/tasks/access-application-cluster/create-external-load-balancer/
service.spec.externalTrafficPolicy - 表示此服务是否希望将外部流量路由到节点本地或集群范围的端点。有两个可用选项：Cluster（默认）和 Local。
Cluster 隐藏了客户端源 IP，可能导致第二跳到另一个节点，但具有良好的整体负载分布。
Local 保留客户端源 IP 并避免 LoadBalancer 和 NodePort 类型服务的第二跳，但存在潜在的不均衡流量传播风险。

获取客户端真实IP：
方案A: 弄1个网关代理服务,不经过istioGateway，代理转到需要的服务；
前提条件：nginx-ingress-lb 中 externalTrafficPolicy: Local 

方案B：
istio-ingressgateway 中 externalTrafficPolicy: Local 就可以透出传；
externalTrafficPolicy=Cluster 找不到方法；
