FROM golang:latest
WORKDIR /go/src/app

RUN go mod init github.com/go-echarts/go-echarts
RUN go mod tidy

COPY main.go .
