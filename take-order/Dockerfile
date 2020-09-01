FROM golang:latest

ENV GOPROXY https://goproxy.cn,direct

#这里指定了后面COPY相对于哪个路径
WORKDIR $GOPATH/src/take-order

#这里的. 就是当前目录下的所有复制到 容器里的 $GOPATH/src 即 /go/src下
COPY ./ $GOPATH/src/take-order

# 然后在容器里的/go/src下执行gobuild
RUN go build take_order.go

RUN apt-get update
RUN apt-get install -y vim

EXPOSE 31226
# 在容器里的/go/src里执行./take_order
ENTRYPOINT ["./take_order"]