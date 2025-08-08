FROM golang:1.24.0 AS builder

WORKDIR /app

COPY go.mod go.sum ./
RUN go mod download

COPY . .

ENV CGO_ENABLED=0 GOOS=linux GOARCH=amd64

RUN go build -o main ./cmd/server

FROM alpine:latest

WORKDIR /app

COPY --from=builder /app/main .

RUN adduser -D appuser
USER appuser

LABEL org.opencontainers.image.source="https://github.com/israelviana/meta-cloud-api"
LABEL org.opencontainers.image.description="Meta gRPC Server"
LABEL org.opencontainers.image.version="1.0"

EXPOSE 50051

CMD ["./main"]