export LINTER_VERSION ?= 1.50.1

GO_PACKAGES ?= $(shell go list ./... | grep -v 'examples\|qtest\|mock')
TMP_DIR     := $(shell mktemp -d)
DOCKERFILES  = $(shell cd build && find */ -name 'Dockerfile' -print)
ODIR        := build/_output
MODULES      = $(shell cd module && ls -d */)
UNAME       := $(shell uname)

bin:
	@mkdir -p bin

coverage:
	@go test -race -cover -coverprofile=coverage.out ${GO_PACKAGES}
	@go tool cover -func=coverage.out

docker-up:
	@docker-compose -f docker-compose.yaml --env-file .env up -d

docker-down:
	@docker-compose -f docker-compose.yaml --env-file .env down --remove-orphans

test:
	@go generate ./...
	@go test -race -v ${GO_PACKAGES}


tool-migrate: bin
ifeq ($(UNAME), Linux)
	@curl -sSfL https://github.com/golang-migrate/migrate/releases/download/v4.15.1/migrate.linux-amd64.tar.gz | tar zxf - --directory /tmp \
	&& cp /tmp/migrate bin/
else ifeq ($(UNAME), Darwin)
	@curl -sSfL https://github.com/golang-migrate/migrate/releases/download/v4.15.1/migrate.darwin-amd64.tar.gz | tar zxf - --directory /tmp \
	&& cp /tmp/migrate bin/
else
	@echo "Your OS is not supported."
endif

migrate-up: tool-migrate
	@$(foreach module, $(MODULES), cp module/$(module)/db/migrate/*.sql $(TMP_DIR) 2>/dev/null;)
	@bin/migrate -source file://$(TMP_DIR) -database "mysql://$(SERVICE_DB_USERNAME):$(SERVICE_DB_PASSWORD)@tcp($(SERVICE_DB_HOST):$(SERVICE_DB_PORT))/$(SERVICE_DB_DATABASE)" up
