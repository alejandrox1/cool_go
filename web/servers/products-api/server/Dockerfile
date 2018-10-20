FROM golang:1.11.0-stretch AS base                                              

WORKDIR /go/src/github.com/server 

ADD . /go/src/github.com/server 

RUN go get -d github.com/lib/pq github.com/gorilla/mux \
    && go build -o server


CMD ["./server"]
