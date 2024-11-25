FROM golang:1.22-alpine AS builder

WORKDIR /app

COPY go.* .
RUN go mod download

COPY . .

RUN go build -o main ./cmd/what_to_watch/main.go

# Runtime stage
FROM alpine:latest

COPY .env .
COPY templates/ templates/
COPY --from=builder /app/main main

ENTRYPOINT ["/main"]