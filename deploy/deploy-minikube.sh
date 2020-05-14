#!/bin/bash
set -e

shell_dir=$(dirname $0)
cd ${shell_dir}

echo "deploy -------- -------- -------- -------- -------- start"

pwd

# deploy

cd ~/.kube
cp -p config-minikube ./config
## namespace
#kubectl delete namespace test-mn
kubectl create namespace test-mn
## deploy
cd ${shell_dir}
kubectl apply -f ./api-app/minikube  -n test-mn
kubectl apply -f ./api-cms/minikube  -n test-mn
kubectl apply -f ./api-www/minikube  -n test-mn
kubectl apply -f ./api-wx/minikube  -n test-mn
kubectl apply -f ./auth-jwt/minikube  -n test-mn
kubectl apply -f ./auth-rbac/minikube  -n test-mn
kubectl apply -f ./gateway-api/minikube  -n test-mn
kubectl apply -f ./sv-ai/minikube  -n test-mn
kubectl apply -f ./sv-config/minikube  -n test-mn
kubectl apply -f ./sv-log/minikube  -n test-mn
kubectl apply -f ./sv-money/minikube  -n test-mn
kubectl apply -f ./sv-mq/minikube  -n test-mn
kubectl apply -f ./sv-notice/minikube  -n test-mn
kubectl apply -f ./sv-oss/minikube  -n test-mn
kubectl apply -f ./sv-region/minikube  -n test-mn
kubectl apply -f ./sv-user/minikube  -n test-mn

echo "deploy  -------- -------- -------- -------- --------  ok"
