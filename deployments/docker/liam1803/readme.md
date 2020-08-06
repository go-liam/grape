# 定制docker镜像

```shell script

docker build -t liam1803/alpine:latest -f Dockerfile .
#docker login --username=liam1803 -p 你的密码
docker push liam1803/alpine:latest
# 查看
docker images
## liam1803/alpine     latest              f70734b6a266        11 days ago         5.61MB

docker ps -a
docker rm mn_gateway_1

```