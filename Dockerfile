FROM golang:1.15.2
 
WORKDIR /go/src/routing-service
 
ADD . /go/src/routing-service

ENV GO111MODULE=on
ENV GOFLAGS=-mod=vendor

RUN go build

EXPOSE 8000

CMD ["./routing-service"]