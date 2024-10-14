#!/bin/bash
kubectl create ns sre
kubectl create configmap settings-admin --from-file=../../config/settings.yml -n sre
