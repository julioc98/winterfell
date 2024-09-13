# Variables
DOCKER_USERNAME := julioc98
JOB_IMAGE := $(DOCKER_USERNAME)/winterfell-job-image:latest
API_IMAGE := $(DOCKER_USERNAME)/winterfell-api-image:latest

# Enable Buildx
BUILDX := docker buildx

# Create a new builder instance
# docker buildx create --name mybuilder --use

# Initialize the builder instance
#docker buildx inspect mybuilder --bootstrap

# Build the job image for multiple architectures
build-job:
	$(BUILDX) build --platform linux/amd64,linux/arm64/v8 -f Dockerfile.job -t $(JOB_IMAGE) --push .

# Build the API image for multiple architectures
build-api:
	$(BUILDX) build --platform linux/amd64,linux/arm64/v8 -f Dockerfile.api -t $(API_IMAGE) --push .

# Build the images
build: build-job build-api

# Clean up local images
clean:
	docker rmi $(JOB_IMAGE) $(API_IMAGE)

# Aplly secrets
k8s-apply-secrets:
	kubectl apply -f ./.k8s/secret.yaml

# Aplly cronjob
k8s-apply-cronjob: 
	kubectl apply -f ./.k8s/cronjob.yaml

# Aplly api
k8s-apply-api: 
	kubectl apply -f ./.k8s/api.yaml

# Apply the Kubernetes resources
k8s-apply: k8s-apply-secrets k8s-apply-cronjob k8s-apply-api

# Delete cronjob
k8s-delete-cronjob:
	kubectl delete cronjob invoices-cronjob

# Delete api
k8s-delete-api:
	kubectl delete deploy  api-deployment
	kubectl delete svc api-service
	kubectl delete ingress api-ingress

# Delete the Kubernetes resources
k8s-delete: k8s-delete-cronjob k8s-delete-api

# Get the Kubernetes resources
k8s-get:
	kubectl get cronjob,pods,svc,deploy,ingress

# Create cluster
k8s-create-cluster:
	kind create cluster --name winterfell-cluster --config ./.k8s/kind-config.yaml

# Delete cluster
k8s-delete-cluster:
	kind delete cluster --name winterfell-cluster

# Start
start: k8s-create-cluster k8s-apply
