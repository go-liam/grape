# info

```shell script
cd deploy/api-cmsw
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

## docker 测试

docker build -t liam1803/api-app:latest -f Dockerfile .

docker run --rm liam1803/api-app

docker run liam1803/api-app
