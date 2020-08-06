# info

```shell script
cd deploy/gateway
make get-config
make build-file
make build-image
make deploy-dev-apply

```

## 发布

```shell script
# 设置权限
ls -l deploy.sh
chmod +x deploy.sh

# 发布
./deploy.sh

```

## ali

kubectl apply -f ./ali-k8s  -n dev-liam
