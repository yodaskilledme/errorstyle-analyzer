# errstyle — GolangCI-Lint Plugin

`errstyle` is a plugin for [`golangci-lint`](https://golangci-lint.run) that enforces a consistent error style based on a custom `domain.Error` type.

## 🔍 What it checks

- The `error` return value **must be the last** in the function signature.
- Errors constructed as `domain.Error` must:
  - Include an `Op` field — typically a string constant equal to the function name.
  - Use a **named string constant** for the `Op` field.
  - Include either the `Err` or `Message` field (or both).

## 🛠️ How to use

1. Add the plugin to the `.golangci.yml` config:

```yaml
plugins:
  - module: github.com/yodaskilledme/errorstyle-analyzer
    version: v0.1.0
```

2. Enable the linter:

```yaml
linters:
  enable:
    - errstyle
```

3. Optionally configure the analyzer:

```yaml
linters-settings:
  errstyle:
    errType: domain.Error     # full type name
    op_name: op               # expected constant name
```

4. Include the plugin module in your project (e.g. in `tools.go`) to prevent removal by `go mod tidy`:

```go
//go:build tools
package tools

import (
    _ "github.com/yodaskilledme/errorstyle-analyzer/cmd/error_style"
)
```

## 🧪 Testing

To run tests for the analyzer:

```bash
go test ./pkg/analyzer
```

Test data is located under `pkg/analyzer/testdata/src`.

## 📁 Project Structure

- `pkg/analyzer`: Core analyzer implementation
- `cmd/error_style/main.go`: Plugin entrypoint for golangci-lint
- `golangci-lint-plugin.yml`: Plugin manifest

## 📜 License

MIT
