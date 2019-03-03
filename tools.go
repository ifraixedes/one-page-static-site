// +build tools

// Why does this file exists? see https://github.com/golang/go/wiki/Modules#how-can-i-track-tool-dependencies-for-a-module
package tools

import (
	_ "github.com/golangci/golangci-lint/cmd/golangci-lint"
	_ "github.com/mitchellh/gox"
	_ "github.com/petergtz/pegomock/pegomock"
)
