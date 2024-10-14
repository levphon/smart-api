#!/bin/bash
killall smart-api # kill go-admin service
echo "stop smart-api success"
ps -aux | grep smart-api
