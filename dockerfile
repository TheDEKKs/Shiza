FROM golang:1.25.5-alpine AS builder

WORKDIR /app

COPY . .
RUN go mod download


CMD ["go", "run", "./cmd/main.go"]
