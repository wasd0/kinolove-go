FROM golang:1.22.4 AS builder

ARG CGO_ENABLED=0
WORKDIR /app

COPY go.mod go.sum ./
RUN go mod download
COPY /. .
COPY .env .

RUN go build -o main.out cmd/kinolove/main.go

FROM scratch

COPY --from=builder /app/main.out /main.out
COPY --from=builder /app/.env /.env
COPY --from=builder /app/config /config

ENTRYPOINT ["/main.out"]