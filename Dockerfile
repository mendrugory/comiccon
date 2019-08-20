FROM golang:1.10.4 as builder

WORKDIR /go/src/comiccon
RUN go get -u github.com/golang/dep/cmd/dep

ADD . .

RUN dep ensure
RUN CGO_ENABLED=0 go build


FROM alpine
RUN apk update && apk add ca-certificates && rm -rf /var/cache/apk/*
COPY --from=builder /go/src/comiccon/comiccon /usr/local/bin
ENTRYPOINT [ "comiccon" ]


