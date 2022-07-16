DOCKER ?= docker
DOCKER_IMAGE ?= edge-example
DOCKER_TAG ?= latest
DOCKER_CONTAINER ?= edge-example
DOCKER_CONTAINER_NAME ?= edge-example
DOCKER_CONTAINER_PORT ?= 8080
DOCKER_FILE_DEV ?= Dockerfile_dev

# Build the docker image for development
docker-dev-build:
	$(DOCKER) build -t $(DOCKER_IMAGE):$(DOCKER_TAG) . -f $(DOCKER_FILE_DEV)

# Run the docker image for development
docker-dev-run:
	$(DOCKER) run -p $(DOCKER_CONTAINER_PORT):8080 -t $(DOCKER_IMAGE):$(DOCKER_TAG)