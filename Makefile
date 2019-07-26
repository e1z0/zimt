CMD_NAME ?= "zimt"
BIN_PATH ?= "./bin/${CMD_NAME}"
PKGS = $(shell go list ./...)

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

.PHONY: build
