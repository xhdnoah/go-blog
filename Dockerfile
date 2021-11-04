FROM golang:latest

ENV GOPROXY https://goproxy.cn, direct
WORKDIR $GOPATH/src/github.com/xhdnoah/go-blog
COPY . $GOPATH/src/github.com/xhdnoah/go-blog
RUN go build .

EXPOSE 8000
ENTRYPOINT ["./go-blog"]
