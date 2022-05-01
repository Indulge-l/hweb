FROM sunmi-docker-images-registry.cn-hangzhou.cr.aliyuncs.com/public/golang:1.16.2 AS builder
WORKDIR /build
COPY . /build

RUN git config --global url."https://flow:9f6e6800cfae7749eb6c486619254b9c@gomod.sunmi.com".insteadOf "https://gomod.sunmi.com"
RUN go env -w GO111MODULE=on && go env -w GOPROXY=https://goproxy.cn,direct && go env -w GOPRIVATE=gomod.sunmi.com && go env -w GOSUMDB=off && go mod tidy
RUN CGO_ENABLED=0 GOOS=linux go build -o main main.go

FROM sunmi-docker-images-registry.cn-hangzhou.cr.aliyuncs.com/public/alpine:3.14
WORKDIR /app

RUN sed -i 's/dl-cdn.alpinelinux.org/mirrors.tuna.tsinghua.edu.cn/g' /etc/apk/repositories && apk add --no-cache curl net-tools iproute2
COPY --from=builder /build/main  /usr/bin/main