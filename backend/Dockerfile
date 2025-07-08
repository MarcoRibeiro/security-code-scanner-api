FROM golang:1.24.4-alpine AS builder

WORKDIR /app

COPY go.mod go.sum ./
RUN go mod download

COPY . .

RUN CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go build -ldflags="-s -w" -o api ./cmd/code_scanner

FROM alpine:latest

WORKDIR /app

COPY --from=builder /app .

EXPOSE 8080

CMD ["./api"]
