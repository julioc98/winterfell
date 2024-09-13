# Winterfell Project

## Overview

Winterfell is a Go-based project designed to integrate with the Stark Bank platform for generating and managing invoices. It provides two main components: a REST API and a job service for batch invoice creation. The project is containerized with Docker, supports multi-architecture builds, and is ready for Kubernetes deployment.

## Features

- Written in Go (version 1.22.6)
- Stark Bank integration using `sdk-go` and `core-go`
- REST API built with `go-chi` for managing transfers
- Job service for generating multiple invoices using `gofakeit`
- Configuration management with `github.com/ardanlabs/conf`
- Docker and Kubernetes support with automated tasks via Makefile

## Dependencies

- **Go version**: 1.22.6
- **Go Modules**:
  - `github.com/ardanlabs/conf/v3 v3.1.8`
  - `github.com/brianvoe/gofakeit/v6 v6.28.0`
  - `github.com/go-chi/chi/v5 v5.1.0`
  - `github.com/starkinfra/core-go v0.2.3`
  - `github.com/starkbank/sdk-go v0.5.0`

## Project Structure

```bash
.
├── Dockerfile.api           # Docker configuration for API service
├── Dockerfile.job           # Docker configuration for Job service
├── Makefile                 # Build and automation tasks
├── cmd/                     # Application entry points
│   ├── api/                 # API service
│   └── job/                 # Job service
├── internal/                # Internal libraries and business logic
│   ├── app/                 # Use cases
│   ├── domain/              # Domain models (e.g., transfers, webhooks)
│   └── infra/               # Infrastructure layer (API and gateways)
├── script/                  # Helper scripts for balance and transfers
├── .k8s/                    # Kubernetes deployment files
│   ├── secret.example.yaml  # Example secret configuration for Kubernetes
│   ├── cronjob.yaml         # CronJob configuration for the invoice generation job
│   ├── api.yaml             # API deployment, service, and ingress configuration
│   └── kind-config.yaml     # Kind cluster configuration
└── go.mod                   # Go module configuration
```

## Prerequisites

- [Go](https://golang.org/dl/) 1.22.6 or later
- [Docker](https://www.docker.com/)
- [Kubectl](https://kubernetes.io/docs/tasks/tools/) for Kubernetes management
- [Kind](https://kind.sigs.k8s.io/) for Kubernetes clusters

## Setup and Installation

### Step 1: Clone the repository

```bash
git clone https://github.com/julioc98/winterfell.git
cd winterfell
```

### Step 2: Setup the Environment Variables

This project requires the Stark Bank private key to be provided through environment variables. You can create a `.env` file to store these values.

Create a `.env` file in the project root:

```bash
touch .env
```

Then add the following content to the `.env` file, replacing the private key placeholder with your actual key:

```bash
export PROJECT_PRIVATE_KEY="-----BEGIN EC PRIVATE KEY-----\\n...\\n-----END EC PRIVATE KEY-----"
```

You can load the environment variables by running:

```bash
source .env
```

### Step 3: Build the Docker images

For the API service:

```bash
docker build -f Dockerfile.api -t winterfell-api .
```

For the job service:

```bash
docker build -f Dockerfile.job -t winterfell-job .
```

Alternatively, you can use the provided Makefile to build both images for multiple architectures (linux/amd64, linux/arm64/v8):

```bash
make build
```

### Step 4: Running the services

To run the API service locally:

```bash
docker run --rm -it -p 3000:3000 --env-file .env winterfell-api
```

To run the job service:

```bash
docker run --rm -it --env-file .env winterfell-job
```

Alternatively, run the services directly using Go:

For the API:

```bash
go run ./cmd/api
```

For the job:

```bash
go run ./cmd/job
```

## Kubernetes Deployment

This project is designed for Kubernetes deployment with automated tasks using the Makefile. 

### Step 1: Create a Kubernetes cluster using Kind

The `kind-config.yaml` file configures a local Kind cluster with exposed ports for the API service.

```bash
make k8s-create-cluster
```

### Step 2: Apply the Kubernetes resources

1. Replace the example secret file with your private key. The `PROJECT_PRIVATE_KEY` value should be base64 encoded.

```bash
cp .k8s/secret.example.yaml .k8s/secret.yaml
# Update PROJECT_PRIVATE_KEY with your base64-encoded private key
```

2. Apply the Kubernetes resources:

```bash
make k8s-apply
```

This will apply the secret, API deployment, service, ingress, and the CronJob for generating invoices.

### Step 3: Get the current status of Kubernetes resources

```bash
make k8s-get
```

### Cleanup

To delete the cluster:

```bash
make k8s-delete-cluster
```

To delete specific Kubernetes resources:

```bash
make k8s-delete
```

## Makefile Commands

- **Build Images**:
  - `make build`: Builds both the API and job Docker images for multiple architectures
  - `make clean`: Removes local Docker images
- **Kubernetes**:
  - `make k8s-apply`: Applies secrets, cronjobs, and API resources to the cluster
  - `make k8s-delete`: Deletes the Kubernetes cronjob and API deployment
  - `make k8s-create-cluster`: Creates a Kubernetes cluster using Kind
  - `make k8s-delete-cluster`: Deletes the Kind cluster

## Docker Hub

Prebuilt Docker images are available on Docker Hub for ARM64 architecture:

- API: `julioc98/winterfell-api-image:latest`
- Job: `julioc98/winterfell-job-image:latest`

Pull the latest images:

```bash
docker pull julioc98/winterfell-api-image:latest
docker pull julioc98/winterfell-job-image:latest
```

## Contributing

Contributions are welcome! Please open issues or pull requests with any improvements.
