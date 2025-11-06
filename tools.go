//go:build tools

package tools

import (
    _ "github.com/golangci/golangci-lint/cmd/golangci-lint"
    _ "golang.org/x/tools/cmd/goimports"
    _ "honnef.co/go/tools/cmd/staticcheck" // 另一個常見的 linter
)