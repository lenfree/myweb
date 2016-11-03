FROM golang:1.7.1
MAINTAINER lenfree.yeung@gmail.com

ENV GOPATH=/go

WORKDIR /go/src/github.com/lenfree/myweb

RUN mkdir -p /go/src/github.com/lenfree/myweb
COPY . /go/src/github.com/lenfree/myweb

CMD ["go run *.go"]
