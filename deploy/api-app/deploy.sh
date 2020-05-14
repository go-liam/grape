#!/bin/bash
set -e

shell_dir=$(dirname $0)
cd ${shell_dir}

echo "deploy -------- -------- -------- -------- -------- start"

pwd

# utest
#echo "deploy  -------- -------- -------- -------- --------  unit"
make test-unit || (echo "[ERROR]test-unit $$?"; exit 1)

# docker

cd ${shell_dir}

echo "deploy  -------- -------- -------- -------- --------  build file"
make build-file || (echo "[ERROR]build-file $$?"; exit 1)
echo "deploy  -------- -------- -------- -------- --------  build image"
make build-image || (echo "[ERROR]build-image $$?"; exit 1)

# dev: docker-compose
#make deploy-dev-apply || (echo "[ERROR]build-image $$?"; exit 1)

# test: minikube
# make deploy-test-apply || (echo "[ERROR]deploy-test-apply $$?"; exit 1)

# prod :ali k8s
# make deploy-pre-apply || (echo "[ERROR]deploy-pre-apply $$?"; exit 1)

# prod :ali k8s
# make deploy-prod-apply || (echo "[ERROR]deploy-prod-apply $$?"; exit 1)

echo "deploy  -------- -------- -------- -------- --------  ok"
