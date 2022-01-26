FROM golang:latest
WORKDIR /go/src/app
COPY . .
RUN go get -u github.com/go-echarts/go-echarts/v2/...
