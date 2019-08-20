FROM golang:1.10.4 as builder

WORKDIR /go/src/comiccon
RUN go get -u github.com/golang/dep/cmd/dep

ADD . .

RUN dep ensure
RUN go build

FROM alpine

WORKDIR /app
COPY --from=builder /go/src/comiccon/comiccon .

ENTRYPOINT [ "./comiccon" ]


