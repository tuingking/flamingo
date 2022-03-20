.PHONY: build run

# version settings
NAMESPACE = flamingo
BUILD_TIME ?= $(shell date -u '+%Y-%m-%d_%H:%M:%S')
BUILD_VERSION ?= $(shell git describe --tag)
COMMIT_HASH ?= $(shell git rev-parse --short HEAD)

build:
	@echo "😎 build app"
	@go build -ldflags "-X main.Namespace=${NAMESPACE} \
		-X main.BuildVersion=${BUILD_VERSION} \
		-X main.BuildTime=${BUILD_TIME} \
		-X main.CommitHash=${COMMIT_HASH}" \
		--race --tags=dynamic -o ./bin/${NAMESPACE}-rest-server cmd/rest-server/main.go

run: build
	@./bin/${NAMESPACE}-rest-server
	