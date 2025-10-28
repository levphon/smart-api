#!/bin/bash
# 创建k8s sre 命名空间
kubectl create ns sre

# 拷贝配置文件
cp ../../config/settings.full.yml settings.yml

# 创建secret
kubectl create secret generic smart-settings \
  --from-file=settings.yml=./settings.yml \
  -n sre

# 部署服务，并挂载secret
kubectl apply -f scripts/smart-api-deploy.yaml
