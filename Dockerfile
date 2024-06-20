FROM golang:1.20 AS builder

COPY . /src
WORKDIR /src

RUN GOPROXY=https://goproxy.cn make build

FROM debian:stable-slim

COPY --from=builder /src/bin /app

WORKDIR /app

EXPOSE 8080
EXPOSE 9090
VOLUME /data/conf

CMD ["./server", "-conf", "/data/conf"]
