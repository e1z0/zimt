CMD_NAME ?= "zimt"
BIN_PATH ?= "./bin/${CMD_NAME}"
PKGS = $(shell go list ./...)
TEST_PKGS=$(shell go list ./pkg/... 2> /dev/null)

build:
	@go mod tidy
	@echo ">> Building ${BIN_PATH}"
	@go build -o $(BIN_PATH)
	@echo ">> Done"
	@echo ""

pkgs:
	@for word in ${PKGS}; do\
		echo $$word;\
	done

format:
	@echo ">> Formatting codebase..."
	@go fmt $(PKGS)

lint:
	@echo ">> Linting codebase..."
	@golint $(PKGS)

test:
	@echo ">> Running tests..."
	@go test -v -race ${TEST_PKGS}

.PHONY: build test
