#!/bin/bash
echo "go build"
go mod tidy
go build -o smart-api main.go
chmod +x ./smart-api
echo "kill smart-api service"
killall smart-api # kill smart-api service
nohup ./smart-api server -c=config/settings.dev.yml >> access.log 2>&1 & #后台启动服务将日志写入access.log文件
echo "run smart-api success"
ps -aux | grep smart-api
