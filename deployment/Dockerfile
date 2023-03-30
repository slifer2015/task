FROM golang:1.20-alpine as builder
RUN apk add --update make git mercurial openssh

ENV GO111MODULE=on
WORKDIR /go/src/

COPY . .
RUN make

# Production image
FROM alpine:3.12

RUN apk add --no-cache ca-certificates make git \
    && rm -rf /var/cache/apk/*

WORKDIR /app/
COPY --from=builder /go/src/.build/server server

CMD /app/server