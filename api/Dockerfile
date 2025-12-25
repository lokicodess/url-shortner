# Build stage
FROM golang:1.24-alpine AS builder

ENV CGO_ENABLED=0

WORKDIR /app

COPY go.mod go.sum ./

RUN go mod tidy && go mod download

COPY . .

RUN go build -o /app/app ./cmd/web/

# Runner Stage 
FROM alpine:3.20
WORKDIR /app
COPY --from=builder /app/app .
EXPOSE 8080
CMD ["./app"]
