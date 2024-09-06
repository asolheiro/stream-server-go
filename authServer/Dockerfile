FROM golang:1.23 AS builder
LABEL authors="asolheiro"

WORKDIR /app

COPY . .

RUN CGO_ENABLED=0 GOOS=linux GOARCH=amd64 \
    go build \
    -a -installsuffix cgo \
    -o stream-key-manager \
    ./cmd/main.go

RUN chmod +x ./stream-key-manager

FROM alpine:latest

WORKDIR /app

COPY --from=builder /app/stream-key-manager ./

RUN apk update && apk add curl

EXPOSE 8000

ENTRYPOINT ["./stream-key-manager"]