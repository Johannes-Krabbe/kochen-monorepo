# Build stage
FROM golang:1.24-alpine AS builder

WORKDIR /app

# Copy go mod and sum files
COPY go.mod go.sum ./

# Download dependencies
RUN go mod download

# Copy source code
COPY . .

# Build the application
RUN CGO_ENABLED=0 GOOS=linux go build -a -installsuffix cgo -o main ./cmd/server

# Build the migrator
RUN CGO_ENABLED=0 GOOS=linux go build -a -installsuffix cgo -o migrator ./cmd/migrator

# Final stage
FROM alpine:latest

# Install ca-certificates for HTTPS requests
RUN apk --no-cache add ca-certificates

WORKDIR /root/

# Copy the binaries from builder stage
COPY --from=builder /app/main .
COPY --from=builder /app/migrator .

# Copy migration files
COPY --from=builder /app/internal/database/schema ./internal/database/schema

# Expose port 8080
EXPOSE 8080

# Run the binary
CMD ["./main"]
