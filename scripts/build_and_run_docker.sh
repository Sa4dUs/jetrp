#!/bin/bash

echo "Building and running Docker containers..."

docker build -t jetrp -f docker/Dockerfile .
docker build -t backend1 -f docker/backend.Dockerfile .
docker build -t backend2 -f docker/backend.Dockerfile .

docker push sa4dus/jetrp:latest
docker push sa4dus/backend1:latest
docker push sa4dus/backend1:latest

docker-compose up --build -d

echo "Docker services are up and running."

docker ps
