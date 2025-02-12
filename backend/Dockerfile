# FROM golang:1.19-alpine

# RUN mkdir /app

# ADD . /app

# WORKDIR /app

# RUN go env -w  GOPROXY=https://goproxy.cn,direct
# RUN go env -w  GO111MODULE=on
# RUN go env
# RUN go build -o main ./cmd/main.go

# CMD ["/app/main"]

FROM golang:alpine AS builder

WORKDIR /build

ADD go.mod .
ADD go.sum .
COPY . .
RUN go env -w  GO111MODULE=on
RUN go env -w GOPROXY=https://mirrors.aliyun.com/goproxy/,direct

RUN go build -o main ./cmd/main.go
RUN ls -alh


FROM alpine

WORKDIR /build
COPY --from=builder /build/main /build/main
COPY ./.env .
COPY ./assets/dataset/*.json ./assets/dataset/

RUN chmod +x /build/main

CMD ["/build/main"]


EXPOSE 8080

