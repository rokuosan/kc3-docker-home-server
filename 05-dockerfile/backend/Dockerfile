FROM golang:1.23.0-bullseye AS builder

WORKDIR /app

COPY go.mod go.sum ./

RUN go mod download

COPY . .

RUN go build -o /app/main .

FROM debian:bullseye-slim

COPY --from=builder /app/main /app/main

EXPOSE 8080

CMD ["/app/main"]
