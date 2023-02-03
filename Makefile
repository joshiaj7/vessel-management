GO_PACKAGES ?= $(shell go list ./... | grep -v 'examples\|qtest\|mock')

bin:
	@mkdir -p bin

check: tool-basic
	@go fmt ./...
	@goimports -w $(shell find . -type f -name '*.go' -not -path './vendor/*')
	./bin/golangci-lint run -v --timeout 3m0s

coverage:
	@go test -race -cover -coverprofile=coverage.out ${GO_PACKAGES}
	@go tool cover -func=coverage.out

test:
	@go generate ./...
	@go test -race -v ${GO_PACKAGES}

tool-basic: bin
	@test -e ./bin/golangci-lint || curl -sSfL https://raw.githubusercontent.com/golangci/golangci-lint/master/install.sh | sh -s -- -b ./bin v${LINTER_VERSION}
	@go install golang.org/x/tools/cmd/goimports@latest
	@go install github.com/golang/mock/mockgen@latest
