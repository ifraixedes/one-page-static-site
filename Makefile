
DEP_VERSION := "0.5.0"

.PHONY: help
help:                                                    ## Show this help.
	@fgrep -h "##" $(MAKEFILE_LIST) | fgrep -v fgrep | sed -e 's/\\$$//' | sed -e 's/##//'

.PHONY: go-tools-install
go-tools-install: .gti-dep .gti-metalinter               ## Install Go tools.

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

.PHONY: lint
lint:                                                    ## Lint the code.
	@gometalinter --vendor --enable-all --line-length=120 --warn-unmatched-nolint --exclude=vendor --deadline=5m ./...
