
DEP_VERSION := "0.5.0"

.PHONY: help
help: ## Show this help
	@echo "Execute one of this targets: "
	@echo
	@fgrep -h "##" $(MAKEFILE_LIST) | fgrep -v fgrep | sed -e 's/:.*##/:##/' | column -t -s '##'

.PHONY: go-tools-install
go-tools-install: .gti-dep .gti-metalinter .gti-pegomock .gti-gox ## Install Go tools

.PHONY: lint
lint: ## Lint the code
	@gometalinter --vendor --enable-all --line-length=120 --warn-unmatched-nolint --exclude=mock_ --deadline=5m ./...

.PHONY: gen-mocks
gen-mocks: ## Generate the mocks used by the tests
	@go generate

.PHONY: build-bins
build-bins: ## Build the binaries for linux, OSX, Windows and for the arch 386 and amd64
	@mkdir -p build
	@cd build && \
		gox -osarch="windows/386 windows/amd64 linux/386 linux/amd64 darwin/386 darwin/amd64" \
			-output "opss_{{.OS}}_{{.Arch}}" ../cmd

.PHONY: .go-tools-install-ci
.go-tools-install-ci: .gti-dep .gti-metalinter

.PHONY: .gti-dep
.gti-dep:
# Download the dep binary to bin folder in $GOPATH
	@curl -L -s https://github.com/golang/dep/releases/download/v$(DEP_VERSION)/dep-linux-amd64 -o $(GOPATH)/bin/dep
# Make dep binary executable
	@chmod +x $(GOPATH)/bin/dep

.PHONY: .gti-metalinter
.gti-metalinter:
	@go get -u github.com/alecthomas/gometalinter
# gometalinter --install is deprecated but for now it's the best solution:
# https://github.com/alecthomas/gometalinter/issues/418
	@gometalinter --install --update --force --debug

.PHONY: .gti-gmock
.gti-pegomock:
	@go get -u github.com/petergtz/pegomock/...

.PHONY: .gti-gox
.gti-gox:
	@go get -u github.com/mitchellh/gox
