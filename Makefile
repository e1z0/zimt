CMD_NAME ?= "zimt"
BIN_PATH ?= "./bin/${CMD_NAME}"
PKGS = $(shell go list ./...)
TEST_PKGS=$(shell go list ./pkg/... 2> /dev/null)

build: tidy
	@echo ">> Building ${BIN_PATH}"
	@go build -o $(BIN_PATH)
	@echo ">> Done"
	@echo ""

tidy:
	@echo ">> Sanitizing..."
	@go mod tidy

pkgs:
	@for word in ${PKGS}; do\
		echo $$word;\
	done

format:
	@echo ">> Formatting..."
	@go fmt $(PKGS)

vet:
	@echo ">> Vetting..."
	@go vet ${PKGS}

lint:
	@echo ">> Linting..."
	@golint $(PKGS)

test: tidy
	@echo ">> Running tests..."
	@go test -v -race ${TEST_PKGS}

setup-ci:
	@go get -u golang.org/x/lint/golint

.PHONY: build test
