# Build stage
FROM golang:1.24-alpine AS builder

WORKDIR /app

RUN apk add nodejs npm make
RUN go install github.com/a-h/templ/cmd/templ@latest

# Copy go mod and sum files
COPY go.mod go.sum ./

# Download dependencies
RUN go mod download

# Copy source code
COPY . .

# Install npm dependencies for tailwindcss
RUN cd internal/ui && npm install

RUN make build_server
RUN make build_migrator

FROM alpine:latest

WORKDIR /root/

COPY --from=builder /app/bin .

EXPOSE 8080

# Run the binary
CMD ["./server/server"]
