#!/bin/bash
kubectl create ns sre
cp ../../config/settings.full.yml settings.yml
kubectl create configmap settings --from-file=settings.yml -n sre
kubectl apply -f scripts/smart-api-deploy.yaml
