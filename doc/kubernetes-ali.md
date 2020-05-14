# dev k8s

```shell script
# 换环境
make deploy-pre-config
# get
kubectl get namespace
kubectl get node
kubectl get svc  -n dev-liam
kubectl get pod  -n dev-cbox
kubectl get gateway -n dev-liam
#namespace
kubectl create namespace dev-liam
kubectl delete namespace dev-liam

kubectl port-forward api-www-v1-7b9db584dd-8jkmr 8044:7041 -n dev-liam
kubectl logs api-www-v1-7b9db584dd-8jkmr -n dev-liam
kubectl delete svc api-www -n dev-liam

# istio:
kubectl label namespace dev-liam istio-injection=enabled
```

## 测试发布

api-www 项目测试：

```shell script
make deploy-pre-apply
```

kubectl apply -f ./ali-k8s-liam  -n dev-liam

http://gateway-api-pre.mygs.com/api-www/ip


## 客户端ip
virtualservice.networking.istio.io/api-www unchanged
The Service "api-www" is invalid: spec.externalTrafficPolicy: Invalid value: "Local": ExternalTrafficPolicy can only be set on NodePort and LoadBalancer service
deploy-pre finish

