NAME="github.com/odpf/siren"
LAST_COMMIT := $(shell git rev-parse --short HEAD)
LAST_TAG := "$(shell git rev-list --tags --max-count=1)"
APP_VERSION := "$(shell git describe --tags ${LAST_TAG})-next"
PROTON_COMMIT := "675f38f48c3cf83918fd71ab85ad49ac568bd945"

.PHONY: all build test clean dist vet proto install

all: build

build: ## Build the siren binary
	@echo " > building siren version ${APP_VERSION}"
	go build -ldflags "-X main.Version=${APP_VERSION}" ${NAME}
	@echo " - build complete"

test: ## Run the tests
	go test ./... -race -covermode=atomic -coverprofile=coverage.out

e2e-test: ## Run all e2e tests
	go test -v -race ./test/e2e_test/... -coverprofile=coverage.out

coverage: ## Print code coverage
	go test -race -coverprofile coverage.out -covermode=atomic ./... && go tool cover -html=coverage.out

generate: ## run all go generate in the code base (including generating mock files)
	go generate ./...

lint: ## lint checker
	golangci-lint run

proto: ## Generate the protobuf files
	@echo " > generating protobuf from odpf/proton"
	@echo " > [info] make sure correct version of dependencies are installed using 'make install'"
	@buf generate https://github.com/odpf/proton/archive/${PROTON_COMMIT}.zip#strip_components=1 --template buf.gen.yaml --path odpf/siren
	@echo " > protobuf compilation finished"

clean: ## Clean the build artifacts
	rm -rf siren dist/

install: ## install required dependencies
	@echo "> installing dependencies"
	go mod tidy
	go get -d github.com/vektra/mockery/v2@v2.13.1
	go get -d google.golang.org/protobuf/cmd/protoc-gen-go@v1.28.1
	go get google.golang.org/protobuf/proto@v1.28.1
	go get google.golang.org/grpc@v1.49.0
	go get -d google.golang.org/grpc/cmd/protoc-gen-go-grpc@v1.2.0
	go get -d github.com/grpc-ecosystem/grpc-gateway/v2/protoc-gen-grpc-gateway@v2.11.3
	go get -d github.com/grpc-ecosystem/grpc-gateway/v2/protoc-gen-openapiv2@v2.11.3
	go get -d github.com/bufbuild/buf/cmd/buf@v1.7.0
	go get github.com/envoyproxy/protoc-gen-validate@v0.6.7

help: ## Display this help message
	@cat $(MAKEFILE_LIST) | grep -e "^[a-zA-Z_\-]*: *.*## *" | awk 'BEGIN {FS = ":.*?## "}; {printf "\033[36m%-30s\033[0m %s\n", $$1, $$2}'