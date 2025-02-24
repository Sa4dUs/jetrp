# JetRP

This repository provides a reverse proxy with load balancing, backend services, and benchmarking capabilities. The reverse proxy supports round-robin load balancing across backend services.

This guide will walk you through setting up Minikube, deploying the reverse proxy along with the backends, and running benchmarks to test performance.

## Prerequisites

1. **Minikube** (for Kubernetes deployment)
    - [Install Minikube](https://minikube.sigs.k8s.io/docs/installation/) on your local machine.
2. **Docker** (for building Docker images)

    - [Install Docker](https://docs.docker.com/get-docker/) on your local machine.

3. **Kubectl** (for interacting with the Kubernetes cluster)

    - [Install kubectl](https://kubernetes.io/docs/tasks/tools/install-kubectl/) on your local machine.

4. **Go** (if you want to build from source)

    - [Install Go](https://golang.org/doc/install) on your local machine.

5. **Minikube Addons** (like ingress controller)
    - You might need to enable a specific ingress controller if you want to test with external access to the proxy.

---

## Steps to Deploy the Proxy

### 1. Start Minikube

You need a Kubernetes cluster running locally using Minikube. You can start Minikube by running:

```bash
minikube start
```

This will start a Minikube cluster with a local Kubernetes setup.

### 2. Build the Docker Images

Once Minikube is running, you need to build the Docker images and push them to your local Docker registry or a container registry (like Docker Hub).

- **Build the reverse proxy image**:

```bash
sudo docker build -t <your_dockerhub_username>/jetrp-proxy:latest -f docker/Dockerfile .
```

- **Build the backend image**:

```bash
sudo docker build -t <your_dockerhub_username>/jetrp-backend:latest -f docker/backend.Dockerfile .
```

- **Push these images to Docker Hub (if necessary)**:

```bash
sudo docker push <your_dockerhub_username>/jetrp-proxy:latest
sudo docker push <your_dockerhub_username>/jetrp-backend:latest
```

### 3. Deploy Services on Minikube

Now, let’s deploy the proxy and backend services to Minikube.

- **Apply the Kubernetes deployments for backend services and the reverse proxy:**

```bash
kubectl apply -f deployments/kubernetes/backend1-deployment.yaml
kubectl apply -f deployments/kubernetes/backend2-deployment.yaml
kubectl apply -f deployments/kubernetes/proxy-deployment.yaml
```

These files define the deployments of your backend services and the reverse proxy service.

### 4. Expose the Reverse Proxy Service

Once the services are deployed, you need to expose the reverse proxy service to access it from outside the Kubernetes cluster.

- **Expose the proxy service using `kubectl expose`**:

```bash
kubectl expose deployment jetrp-proxy --type=LoadBalancer --name=jetrp-proxy-service --port=8080 --target-port=8080
```

- **Get the external IP address of the reverse proxy**:

```bash
kubectl get svc jetrp-proxy-service
```

If you're using Minikube, you can also get the external IP by running:

```bash
minikube service jetrp-proxy-service --url
```

This will provide you with the URL for accessing the reverse proxy.

### 5. Run Benchmark Tests

To evaluate the performance of the reverse proxy and backends, you can run some benchmark tests using wrk or hey (depending on which tool you prefer).

**Running Benchmarks with `wrk`**

- **Install wrk** on your machine.
- Run the benchmark script for `wrk`:

```bash
    ./benchmarks/run_wrk.sh
```

This will run the `wrk` benchmark against your reverse proxy service and print the results.

**Running Benchmarks with `hey`**

- **Install hey** on your machine.
- Run the benchmark script for `hey`:

```bash
    ./benchmarks/run_hey.sh
```

This will run the `hey` benchmark against your reverse proxy service and print the results.

### 6. Verify the Deployment

To verify that everything is working as expected:

- Access the proxy through the Minikube service URL or directly by using `kubectl port-forward` for internal access.

For example, to access the proxy on port `8080`:

```bash
kubectl port-forward svc/jetrp-proxy-service 8080:8080
```

Now you can try to send requests to `http://localhost:8080` and the reverse proxy will route them to the backend services.

### 7. Cleanup

Once you're done with the setup and testing, you can clean up the resources:

- **Stop the Minikube cluster**:

```bash
minikube stop
```

- **Delete the deployed Kubernetes resources**:

```bash
kubectl delete -f deployments/kubernetes/backend1-deployment.yaml
kubectl delete -f deployments/kubernetes/backend2-deployment.yaml
kubectl delete -f deployments/kubernetes/proxy-deployment.yaml
```

- **Clean up Docker images if necessary**:

```bash
sudo docker rmi <your_dockerhub_username>/jetrp-proxy:latest
sudo docker rmi <your_dockerhub_username>/jetrp-backend:latest
```

## Troubleshooting

- **"Connection refused" error**: If you see connection issues (e.g., connection refused), make sure that the backend services are running correctly by checking the status of the pods in Kubernetes:

```bash
kubectl get pods
```

You can also check logs of the proxy and backend services:

```bash
kubectl logs <proxy_pod_name>
kubectl logs <backend_pod_name>
```

- **Proxy not forwarding requests**: Ensure that the proxy is configured with the correct backend service URLs and that the services are exposed correctly in Kubernetes.
