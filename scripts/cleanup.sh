#!/bin/bash
echo "Cleaning up all services and containers..."

if docker-compose ps > /dev/null 2>&1; then
  echo "Stopping containers with Docker Compose..."
  docker-compose down
fi

echo "Removing containers, images, and networks..."
docker container prune -f
docker image prune -af
docker network prune -f

echo "Stopping Kubernetes pods..."
kubectl delete -f deployments/kubernetes/ || echo "No active Kubernetes deployments."

echo "Cleaning Kubernetes logs..."
kubectl logs -l app=proxy --tail=0 || echo "No Kubernetes logs found."

echo "Cleanup completed."
