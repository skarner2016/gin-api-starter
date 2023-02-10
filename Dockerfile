ARG GO_VERSION=1.18

# build
FROM golang:${GO_VERSION}-alpine AS builder
ENV GO111MODULE=on \
    GOPROXY=https://goproxy.cn,direct
RUN apk update && apk add alpine-sdk git && rm -rf /var/cache/apk/*
RUN mkdir -p /workspace
WORKDIR /workspace
COPY . .
RUN go mod download
RUN go build -o ./main.app

# image
FROM alpine:latest
RUN apk update && apk add ca-certificates && rm -rf /var/cache/apk/*
RUN mkdir -p /workspace
WORKDIR /workspace
COPY --from=builder /workspace/main.app .
EXPOSE 8008

ENTRYPOINT ["./main.app"]