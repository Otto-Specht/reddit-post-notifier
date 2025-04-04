# ---- Build Stage ----
    FROM golang:1.23-alpine AS builder

    WORKDIR /app

    RUN apk add --no-cache git
    
    COPY go.mod go.sum ./
    RUN go mod download
    
    COPY . .
    
    RUN go build -o main ./cmd/main.go
    
# ---- Runtime Stage ----
    FROM alpine:3.18

    ENV GIN_MODE=release
    
    COPY --from=builder /app/main /app/main
    
    ENTRYPOINT ["/app/main"]
    