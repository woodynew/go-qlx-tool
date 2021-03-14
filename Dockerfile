FROM library/golang:1.16
  
WORKDIR /go/src/app
COPY . /go/src/app

RUN go env -w GO111MODULE=on
RUN go env -w GOPROXY=https://goproxy.cn,direct

RUN export GO111MODULE=on
RUN export GOPROXY=https://goproxy.cn

RUN chmod 777 ./go-qlx-tool

EXPOSE 9011

ENTRYPOINT ["./go-qlx-tool"]