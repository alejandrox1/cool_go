FROM golang:1.10.5-stretch AS base

WORKDIR /go/src/app
COPY main.go .

RUN GOOS=linux GOARCH=amd64 CGO_ENABLED=0 go build -a -installsuffix cgo -o http-server


FROM scratch

COPY --from=base /go/src/app/http-server /bin/http-server

CMD ["/bin/http-server"]
