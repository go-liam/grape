#!/bin/bash
set -e

shell_dir=$(dirname $0)
cd ${shell_dir}

echo "deploy -------- -------- -------- -------- -------- start"

pwd

# deploy
echo "deploy -------- -------- -------- -------- -------- [docker-compose]"
cd ./docker/mn
docker-compose up -d
#docker-compose down

echo "deploy  -------- -------- -------- -------- --------  ok"
