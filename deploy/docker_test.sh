#!/bin/bash
set -e


echo "deploy -------- -------- -------- -------- -------- start"

echo -e "\n gateway \n"
curl http://localhost:6001

echo -e "\n app \n"
curl http://localhost:6001/api-app
echo -e "\n cms \n"
curl http://localhost:6001/api-cms
echo -e "\n www \n"
curl http://localhost:6001/api-www
echo -e "\n wx \n"
curl http://localhost:6001/api-wx
echo -e "\n jwt \n"
curl http://localhost:6001/auth-jwt
echo -e "\n rbac \n"
curl http://localhost:6001/auth-rbac
echo -e "\n ai \n"
curl http://localhost:6001/sv-ai
echo -e "\n config \n"
curl http://localhost:6001/sv-config
echo -e "\n log \n"
curl http://localhost:6001/sv-log
echo -e "\n money \n"
curl http://localhost:6001/sv-money
echo -e "\n mq \n"
curl http://localhost:6001/sv-mq
echo -e "\n notice \n"
curl http://localhost:6001/sv-notice
echo -e "\n oss \n"
curl http://localhost:6001/sv-oss
echo -e "\n region \n"
curl http://localhost:6001/sv-region
echo -e "\n sv-user \n"
curl http://localhost:6001/sv-user

echo -e "\n deploy  -------- -------- -------- -------- --------  ok"
