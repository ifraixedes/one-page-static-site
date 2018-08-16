
DEP_VERSION := "0.5.0"

.PHONY: help
help:                                                    ## Show this help.
	@fgrep -h "##" $(MAKEFILE_LIST) | fgrep -v fgrep | sed -e 's/\\$$//' | sed -e 's/##//'

.PHONY: go-tools-install
go-tools-install:                                        ## Install Go tools.
	@make .gti-dep .gti-metalinter .gti-pegomock

.PHONY: lint
lint:                                                    ## Lint the code.
	@gometalinter --vendor --enable-all --line-length=120 --warn-unmatched-nolint --exclude=vendor --exclude=mock_ --deadline=5m ./...

.PHONY: gen-mocks
gen-mocks:                                               ## Generate the mocks used by the tests
	@go generate

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
	@go get github.com/petergtz/pegomock/...
