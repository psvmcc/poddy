#COMMIT ?= $$(git rev-parse HEAD)
#TAG ?= $$(git describe --tags --abbrev=0 2>/dev/null || echo dev)

TAG = 1.2.3.4
COMMIT = qwerty12334


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
	$(shell which golangci-lint) run main.go

clean:
	rm -rf build
	mkdir build

.PHONY: build
build: clean
	CGO_ENABLED=0 go build -mod=vendor -ldflags "-X poddy/pkg/types.Version=${TAG} -X poddy/pkg/types.Commit=${COMMIT}" -o build/poddy main.go

linux:
	env GOOS=linux GOARCH=amd64 go build -mod=vendor -ldflags "-X main.version=${TAG} -X main.commit=${COMMIT}" -o build/poddy.linux main.go

run:
	go run main.go s
