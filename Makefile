COMMIT ?= $$(git rev-parse HEAD)
TAG ?= $$(git describe --tags --abbrev=0 2>/dev/null || echo dev)

init:
	go mod init poddy

deps:
	go mod tidy
	go mod vendor

pre-test:
	go install golang.org/x/tools/go/analysis/passes/shadow/cmd/shadow@latest
	go install github.com/golangci/golangci-lint/cmd/golangci-lint@latest

test: pre-test
	go vet -mod=vendor $(shell go list ./...)
	go vet -mod=vendor -vettool=$(shell which shadow) $(shell go list ./...)
	golangci-lint run -v

clean:
	rm -rf build
	mkdir build

.PHONY: build
build: clean
	CGO_ENABLED=0 go build -mod=vendor -ldflags "-X poddy/pkg/types.Version=${TAG} -X poddy/pkg/types.Commit=${COMMIT}" -o build/poddy main.go

linux:
	GOOS=linux GOARCH=amd64 CGO_ENABLED=0 go build -mod=vendor -ldflags "-X poddy/pkg/types.Version=${TAG} -X poddy/pkg/types.Commit=${COMMIT}" -o build/poddy.linux main.go

run:
	env DOCKER_SOCKET=/Users/ps/.local/share/containers/podman/machine/qemu/podman.sock go run main.go s

container:
	podman build -t psvmcc/poddy .
