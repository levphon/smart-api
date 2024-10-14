FROM golang:1.21 as builder

WORKDIR  /home/service/

COPY . .

RUN CGO_ENABLED=0 go build -o smart-api main.go

FROM alpine:3.18.4

WORKDIR /home/service/

EXPOSE 8000

RUN apk add bash

COPY --from=builder /home/service/smart-api .

CMD ["/smart-api","server","-c", "/config/settings.yml"]




