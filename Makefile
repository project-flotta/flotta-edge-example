run:
	cp .env.example .env
	docker-compose up -d

build-prod-image:
	docker build -t ahmadateya/flotta-edge-example-workload:latest . -f Dockerfile

# make production build and push to docker hub
push-prod-image: build-prod-image
	docker push ahmadateya/flotta-edge-example-workload:latest

test:
	cd app/ && sudo go test -v ./...