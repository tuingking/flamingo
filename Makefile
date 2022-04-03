.PHONY: build run

# version settings
NAMESPACE = flamingo
BUILD_TIME ?= $(shell date -u '+%Y-%m-%d_%H:%M:%S')
BUILD_VERSION ?= $(shell git describe --tag)
COMMIT_HASH ?= $(shell git rev-parse --short HEAD)

swag:
	@swag init --parseDependency --parseInternal --parseDepth 1 -g cmd/rest/main.go

build: swag
	@echo "📦 building binary..."
	@go build -ldflags "-X main.Namespace=${NAMESPACE} \
		-X main.BuildVersion=${BUILD_VERSION} \
		-X main.BuildTime=${BUILD_TIME} \
		-X main.CommitHash=${COMMIT_HASH}" \
		--race --tags=dynamic -o ./bin/${NAMESPACE}-rest cmd/rest/main.go

kill-process:
	@lsof -i :8080 | awk '$$1 ~ /app/ { print $$2 }' | xargs kill -9 || true

run: kill-process build
	@./bin/${NAMESPACE}-rest
	
run-dev: swag
	@air cmd/rest/main.go

migrate:
	@migrate -path script/migrations -database "mysql://root:password@tcp(localhost:3306)/playground" -verbose up

build-image:
	@ echo "📦 building docker image..."
	@ docker build -f Dockerfile -t flamingo .

docker-run:
	@ echo "📦 run container..."
	@ docker run --rm --name flamingo -p 8080:8080 flamingo