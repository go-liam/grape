#!/bin/bash
set -e

docker build -t liam1803/alpine:latest -f Dockerfile .
#docker login --username=liam1803 -p 你的密码
docker push liam1803/alpine:latest