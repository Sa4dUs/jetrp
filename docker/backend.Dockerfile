# Use Golang base image
FROM golang:1.24-alpine

# Set the working directory
WORKDIR /app

# Copy the backend server source code
COPY internal/backend .

# Build the backend
RUN go build backend.go

# Expose the backend port
EXPOSE 8081

# Run the backend server
CMD ["./backend"]
