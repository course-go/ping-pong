FROM golang:alpine AS builder

WORKDIR /usr/app

COPY . .
RUN go build -v -o /usr/app/bin/ping-pong main.go

FROM alpine:latest

WORKDIR /usr/app

COPY --from=builder /usr/app/bin/ping-pong /usr/app/ping-pong

CMD ["/usr/app/ping-pong"]
