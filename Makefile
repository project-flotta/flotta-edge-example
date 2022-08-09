run:
	cp .env.example .env
	docker-compose up -d

LINT_IMAGE=golangci/golangci-lint:v1.45.0
lint: ## Check if the go code is properly written, rules are in .golangci.yml
	docker run --rm -v $(CURDIR):$(CURDIR):z -w="$(CURDIR)" $(LINT_IMAGE) sh -c 'golangci-lint run'


build-prod-image:
	docker build -t ahmadateya/flotta-edge-example-workload:latest . -f Dockerfile

# make production build and push to docker hub
push-prod-image: build-prod-image
	docker push ahmadateya/flotta-edge-example-workload:latest

test:
	cd app/ && sudo go test -v ./...