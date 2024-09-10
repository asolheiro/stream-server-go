FROM golang:1.23 AS builder
LABEL authors="asolheiro"

WORKDIR /app-playback

COPY . .

RUN CGO_ENABLED=0 GOOS=linux GOARCH=amd64 \
    go build \
    -a -installsuffix cgo \
    -o playback-server \
    ./cmd/main.go

RUN chmod +x ./playback-server

FROM alpine:latest

WORKDIR /app

COPY --from=builder /app-playback/playback-server ./

RUN apk update && apk add curl

EXPOSE 8001

ENTRYPOINT ["./playback-server"]