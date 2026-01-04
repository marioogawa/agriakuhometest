# Monorepo CI/CD Kubernetes Backend Services

## Overview
This repository demonstrates a **monorepo-based backend system** consisting of two HTTP services:

- **Go Service** — simple HTTP API written in Go  
- **Node.js Service** — simple HTTP API written in Node.js  

Both services are:
- Containerized using **Docker**
- Built and published via **GitHub Actions**
- Deployed to a **Kubernetes cluster**

The CI/CD pipeline follows best practices by **separating build & push from deployment**, with deployment executed **manually** to ensure better control and safety.

---


---

## CI/CD Workflow

The CI/CD pipeline is implemented using **GitHub Actions**.

### Triggers
- **Automatic**: On every push to the `main` branch  
  → Build & push Docker images
- **Manual**: Deployment to Kubernetes  
  → Triggered via **GitHub Actions → Run workflow**

---

### Pipeline Stages

#### 1. Build & Push (Automatic)
Triggered on `push` to `main`.

Steps:
1. Checkout source code
2. Build Docker images for:
   - Go service
   - Node.js service
3. Push images to **Docker Hub**

Docker image tags:
- `latest`
- `<git-commit-sha>` (for traceability & rollback)

Example:
docker.io/mariopratama/go-service:latest
docker.io/mariopratama/go-service:<git-sha>


---

#### 2. Deploy to Kubernetes (Manual)
Deployment is **intentionally manual** to prevent accidental releases.

How it works:
- Executed via `workflow_dispatch`
- Requires successful completion of the build & push stage
- Uses `kubectl` with a provided kubeconfig

Deployment steps:
1. Apply Kubernetes manifests
2. Wait for rollout to complete
3. Validate deployment status

This approach reflects real-world production pipelines where deployment requires explicit approval or action.

---

## Kubernetes Deployment

Each service is deployed using standard Kubernetes resources:

- **Deployment** — manages application pods
- **Service** — internal cluster networking
- **Ingress** — external access via host-based routing

### Ingress Configuration
Example domains using `xip.io`:

- Go Service
http://go-service.127.0.0.1.xip.io
- Node.js Service
http://node-service.127.0.0.1.xip.io

> Note: These domains resolve automatically to `127.0.0.1` and are useful for local Kubernetes clusters such as Minikube.

---

## Running Locally (Without Kubernetes)

### Build and Run Go Service
```bash
docker build -t go-service services/go-service
docker run -p 8080:8080 go-service
http://localhost:8080
```

### Build and Run Go Service
```bash
docker build -t node-service services/node-service
docker run -p 3000:3000 node-service
http://localhost:3000
```