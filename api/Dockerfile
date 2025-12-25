# Build stage
FROM golang:1.24-alpine AS builder

WORKDIR /app

# Copy go.mod and go.sum first
COPY go.mod go.sum ./
RUN go mod download

# Copy all source code
COPY . .

# Build binary from cmd/ folder

RUN go build -o /app/app ./cmd/web/

# Runtime stage
FROM alpine:3.18
WORKDIR /app
RUN apk add --no-cache ca-certificates

# Copy binary
COPY --from=builder /app/app .

EXPOSE 8080

CMD ["./app"]
