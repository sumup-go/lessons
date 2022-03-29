# Builder image
FROM golang:1.17-alpine AS builder

COPY . /build
WORKDIR /build

RUN CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go build -o test . 

# Exec image
FROM alpine:latest

COPY --from=builder /build/test /test

ENTRYPOINT ["/test"]
