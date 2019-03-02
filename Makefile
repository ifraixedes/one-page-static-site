
BIN_DIR=.bin

.PHONY: help
help: ## Show this help
	@echo "Execute one of this targets: "
	@echo
	@fgrep -h "##" $(MAKEFILE_LIST) | fgrep -v fgrep | sed -e 's/:.*##/:##/' | column -t -s '##'

.PHONY: go-tools-install
go-tools-install: ## Install Go tools
	@mkdir -p .bin
	@GOBIN=$(realpath $(BIN_DIR)) go install github.com/golangci/golangci-lint/cmd/golangci-lint
	@GOBIN=$(realpath $(BIN_DIR)) go install github.com/petergtz/pegomock/pegomock
	@GOBIN=$(realpath $(BIN_DIR)) go install github.com/mitchellh/gox


.PHONY: lint
lint: ## Lint the sources of all the packages contained in this repo
	@PATH=$(realpath $(BIN_DIR)):$$PATH golangci-lint run --enable-all --exclude-use-default=false

.PHONY: gen-mocks
gen-mocks: ## Generate the mocks used by the tests
	@PATH=$(realpath $(BIN_DIR)):$$PATH go generate

.PHONY: build-bins
build-bins: ## Build the binaries for linux, OSX, Windows and for the arch 386 and amd64
	@mkdir -p build
	@ cd build && \
		PATH=$(realpath $(BIN_DIR)):$$PATH \
		gox -osarch="windows/386 windows/amd64 linux/386 linux/amd64 darwin/386 darwin/amd64" \
			-output "opss_{{.OS}}_{{.Arch}}" ../cmd/opss

.PHONY: test ## Execute all the tests
test:
	@go test -race $(TARGS) ./...

.PHONY: ci
ci: ## Contains the set of checks that the CI runs (can be used for executing in local previous to push code)
	@if [ "$$LINT" = true ]; then make lint; fi
	@if [ "$$COVERAGE" = true ]; then make test TARGS="-v -covermode=atomic -coverprofile=profile.cov"; else make test TARGS="-v";  fi

.PHONY: .go-tools-install-ci
.go-tools-install-ci:
	@go install github.com/golangci/golangci-lint/cmd/golangci-lint
