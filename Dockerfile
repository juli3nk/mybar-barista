FROM golang:alpine AS builder

RUN apk --update add \
		ca-certificates \
		gcc \
		git \
		musl-dev

COPY . /go/src/github.com/juli3nk/mybar/
WORKDIR /go/src/github.com/juli3nk/mybar

ENV GO111MODULE off
ENV GOSUMDB=off

RUN go get
RUN go build -ldflags "-linkmode external -extldflags -static -s -w" -o /tmp/mybar


FROM scratch

COPY --from=builder /tmp/mybar /mybar
