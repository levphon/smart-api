FROM golang:1.21 AS builder

WORKDIR  /home/service/

COPY . .

RUN CGO_ENABLED=0 go build -o smart-api main.go

FROM alpine:3.18.4

WORKDIR /home/service/

EXPOSE 8000

RUN apk add bash

COPY --from=builder /home/service/smart-api .
COPY config  /home/service/config
CMD ["/home/service/smart-api","server","-c", "./config/settings.yml", "-a", "true"]

# 本地如果是arm架构，但是想构建为x86架构的镜像
# docker buildx build --platform linux/amd64 -t registry.cn-beijing.aliyuncs.com/sunwenbo/smart-api:latest . --load
# docker push registry.cn-beijing.aliyuncs.com/sunwenbo/smart-api:latest
# 启动后端服务
# docker run -it --rm   -p 8000:8000   -v /data/config/:/home/service/config/   --name smart-api   registry.cn-beijing.aliyuncs.com/sunwenbo/smart-api:latest



# 如果本地是arm架构，但是想将打好的镜像使用在x86的linux服务器，使用下面命令构建镜像即可兼容本地和linux x86个arm环境
# 在 docker buildx 中进行多平台构建时，默认的 Docker 驱动并不支持多平台构建。要解决这个问题，你可以切换到 docker-container 驱动，它支持多平台构建。
# 解决步骤：
# 创建并使用 docker-container 驱动：

# 1. 首先，创建一个新的 buildx 构建器，并设置为使用 docker-container 驱动。这个驱动支持跨平台构建。
# docker buildx create --name mybuilder --use
# docker buildx inspect mybuilder --bootstrap
# 以上命令会创建并切换到 mybuilder 构建器，并启动必要的基础环境（如 containerd）。

# 2. 运行多平台构建：
# 切换到 docker-container 驱动后，重新运行多平台构建命令：
# docker buildx build --platform linux/amd64 -t sunwenbo/smart-api:latest --load .
# 如果想要将镜像推送到 Docker Hub，可以加上 --push 参数：
# docker buildx build --platform linux/amd64,linux/arm64 -t sunwenbo/smart-api:latest --push .
