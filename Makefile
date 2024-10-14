PROJECT:=smart-api

.PHONY: build
build:
	CGO_ENABLED=0 go build -ldflags="-w -s" -a -installsuffix "" -o smart-api .

# make build-linux
build-linux:
	@docker build -t smart-api:latest .
	@echo "build successful"

build-sqlite:
	go build -tags sqlite3 -ldflags="-w -s" -a -installsuffix -o smart-api .

# make run
run:
    # delete smart-api-api container
	@if [ $(shell docker ps -aq --filter name=smart-api --filter publish=8000) ]; then docker rm -f smart-api; fi

    # 启动方法一 run smart-api-api container  docker-compose 启动方式
    # 进入到项目根目录 执行 make run 命令
	@docker-compose up -d

	# 启动方式二 docker run  这里注意-v挂载的宿主机的地址改为部署时的实际决对路径
    #@docker run --name=smart-api -p 8000:8000 -v /home/code/go/src/smart-api/smart-api/config:/smart-api-api/config  -v /home/code/go/src/smart-api/smart-api-api/static:/smart-api/static -v /home/code/go/src/smart-api/smart-api/temp:/smart-api-api/temp -d --restart=always smart-api:latest

	@echo "smart-api service is running..."

	# delete Tag=<none> 的镜像
	@docker image prune -f
	@docker ps -a | grep "smart-api"

stop:
    # delete smart-api-api container
	@if [ $(shell docker ps -aq --filter name=smart-api --filter publish=8000) ]; then docker-compose down; fi
	#@if [ $(shell docker ps -aq --filter name=smart-api --filter publish=8000) ]; then docker rm -f smart-api; fi
	#@echo "smart-api stop success"


#.PHONY: test
#test:
#	go test -v ./... -cover

#.PHONY: docker
#docker:
#	docker build . -t smart-api:latest

# make deploy
deploy:

	#@git checkout master
	#@git pull origin master
	make build-linux
	make run
