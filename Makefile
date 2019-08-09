CMD_NAME ?= "zimt"
BIN_PATH ?= "./bin/${CMD_NAME}"
PKGS = $(shell go list ./...)
TEST_PKGS=$(shell go list ./pkg/... 2> /dev/null)

PLATFORMS := \
	linux/amd64 \
	darwin/amd64

build: tidy
	@echo ">> Building ${BIN_PATH}"
	@go build -o $(BIN_PATH)
	@echo ">> Done"
	@echo ""

tidy:
	@echo ">> Sanitizing..."
	@go mod tidy

pkgs:
	@for p in ${PKGS}; do\
		echo $$p;\
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

sec:
	@echo ">> Running security audit..."
	@gosec -quiet ./...

test: tidy
	@echo ">> Running tests..."
	@go test -v -race ${TEST_PKGS}

ci: lint vet sec test

setup-ci:
	@go get -u golang.org/x/lint/golint
	@go get -u github.com/securego/gosec/cmd/gosec

dist:
	@ rm -rf ./dist
	@- $(foreach p,$(PLATFORMS), \
		build/scripts/dist.sh $(p); \
	)

.PHONY: build test dist
