FROM golang:1.19-alpine AS builder

ENV GO111MODULE=on \
    GOPROXY=https://goproxy.cn,direct \
    CGO_ENABLED=0

COPY . /

WORKDIR /



RUN go build .



CMD chmod +x addressApi && ./addressApi

