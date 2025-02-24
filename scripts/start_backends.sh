#!/bin/bash

echo "Starting backend servers..."

PORT=8081 go run internal/backend/backend.go &
BACKEND1_PID=$!
echo "Backend1 running on port 8081 with PID $BACKEND1_PID"

PORT=8082 go run internal/backend/backend.go &
BACKEND2_PID=$!
echo "Backend2 running on port 8082 with PID $BACKEND2_PID"

# Wait for processes to exit
wait $BACKEND1_PID $BACKEND2_PID
