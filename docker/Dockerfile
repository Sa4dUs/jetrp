# Use Golang base image
FROM golang:1.24-alpine AS builder

# Set the working directory
WORKDIR /app

# Copy go.mod and go.sum files
COPY go.mod ./

# Download dependencies
RUN go mod download

# Copy the entire source code
COPY . .

# Build the application
RUN go build -o proxy ./cmd/proxy.go

# Use a minimal image for runtime
FROM alpine:latest

# Copy the built binary from the builder image
COPY --from=builder /app/proxy /usr/local/bin/proxy

# Expose the proxy port
EXPOSE 8080

# Run the proxy application
CMD ["proxy"]
