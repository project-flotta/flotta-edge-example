# Edge Example: Workload Application

## Before you begin
you need to have:
- Docker installed
- Docker-compose installed

## Getting Started
to see the app up and running, run:
```
make app-up
```
browse the app at: http://localhost:3000

## For Developers
run the following commands to get started:
```
cp .env.example .env
```
- change the Dockerfile var value in .env to `dockerfile_dev`
```
docker-compose up
```