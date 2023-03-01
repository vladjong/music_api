FROM golang:1.20-alpine AS builder

WORKDIR /app

COPY ./cmd ./cmd
COPY ./internal ./internal
COPY ./go.mod ./go.mod
COPY ./go.sum ./go.sum
COPY ./.env ./.env
COPY ./migration ./migration

RUN go build -o music_api ./cmd/music_api/main.go
CMD ["/app/music_api"]
