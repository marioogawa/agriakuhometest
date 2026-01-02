# Monorepo CI/CD Kubernetes Backend Services

## Overview
This repository demonstrates a monorepo-based setup containing two backend services:
- A Go-based HTTP service
- A Node.js-based HTTP service

Both services are containerized using Docker, built and published via a CI/CD pipeline, and deployed to a Kubernetes cluster.

---

## CI/CD Workflow
The CI/CD pipeline is implemented using **GitHub Actions** and is triggered on every push to the `main` branch.

### Pipeline Steps
1. Checkout source code
2. Build Docker images for Go and Node.js services
3. Push images to Docker Hub
4. Deploy services to Kubernetes using `kubectl`
5. Verify deployment rollout status

Container images are tagged with:
- `latest`
- Git commit SHA (for traceability and rollback)

---

## Kubernetes Deployment
Each service is deployed using:
- **Deployment** for application pods
- **Service** for internal networking
- **Ingress** for external access using host-based routing

Example domains (using xip.io):
- `http://go-service.127.0.0.1.xip.io`
- `http://node-service.127.0.0.1.xip.io`

---

## Running Locally

### Build and run Go service
```bash
docker build -t go-service services/go-service
docker run -p 8080:8080 go-service
