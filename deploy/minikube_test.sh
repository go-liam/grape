#!/bin/bash
set -e

echo "deploy -------- -------- -------- -------- -------- start"

echo -e "\n gateway \n"
curl http://mn-gateway-api.mygs.com

echo -e "\n app \n"
curl http://mn-gateway-api.mygs.com/api-app
echo -e "\n cms \n"
curl http://mn-gateway-api.mygs.com/api-cms
echo -e "\n www \n"
curl http://mn-gateway-api.mygs.com/api-www
echo -e "\n wx \n"
curl http://mn-gateway-api.mygs.com/api-wx
echo -e "\n jwt \n"
curl http://mn-gateway-api.mygs.com/auth-jwt
echo -e "\n rbac \n"
curl http://mn-gateway-api.mygs.com/auth-rbac
echo -e "\n ai \n"
curl http://mn-gateway-api.mygs.com/sv-ai
echo -e "\n config \n"
curl http://mn-gateway-api.mygs.com/sv-config
echo -e "\n log \n"
curl http://mn-gateway-api.mygs.com/sv-log
echo -e "\n money \n"
curl http://mn-gateway-api.mygs.com/sv-money
echo -e "\n mq \n"
curl http://mn-gateway-api.mygs.com/sv-mq
echo -e "\n notice \n"
curl http://mn-gateway-api.mygs.com/sv-notice
echo -e "\n oss \n"
curl http://mn-gateway-api.mygs.com/sv-oss
echo -e "\n region \n"
curl http://mn-gateway-api.mygs.com/sv-region
echo -e "\n sv-user \n"
curl http://mn-gateway-api.mygs.com/sv-user

echo -e "\n deploy  -------- -------- -------- -------- --------  ok"
