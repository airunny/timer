FROM golang:1.20 AS builder

COPY . /src
WORKDIR /src

RUN make build

FROM debian:stable-slim

COPY --from=builder /src/bin/wiki_timer /app/app
COPY --from=builder /src/configs/config.yaml /data/conf/config.yaml

WORKDIR /app

EXPOSE 8080
EXPOSE 9090
VOLUME /data/conf

CMD ["./app", "-conf", "/data/conf"]
