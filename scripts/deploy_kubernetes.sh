#!/bin/bash
echo "Deploying services to Kubernetes..."

kubectl apply -f deployments/kubernetes/backend1-deployment.yaml
kubectl apply -f deployments/kubernetes/backend2-deployment.yaml
kubectl apply -f deployments/kubernetes/proxy-deployment.yaml

echo "Kubernetes deployments applied successfully."
kubectl get pods
