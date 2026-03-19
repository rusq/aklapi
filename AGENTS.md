# AGENTS.md — Coding Agent Instructions for `aklapi`

This document provides guidance for agentic coding assistants operating in this repository.

---

## Project Overview

`aklapi` is a Go library and HTTP server that exposes Auckland Council APIs
(rubbish collection schedules, property address lookup) as a simple REST service.

- **Module:** `github.com/rusq/aklapi` (`go 1.24`)
- **Library package:** root (`aklapi`)
- **Binary:** `cmd/aklapi/` — standard HTTP server on port 8080
- **Language:** Go only — no TypeScript, JavaScript, or Node tooling

---

## Build, Run & Test Commands

```sh
# Build the server binary
go build -o server ./cmd/aklapi

# Run the server (port defaults to 8080)
./server

# Run all tests
go test ./...

# Run all tests with verbose output
go test -v ./...

# Run a single test by name (supports regex)
go test -v -run TestFunctionName ./...

# Run a single test in a specific package
go test -v -run TestCollectionDayDetail ./cmd/aklapi/

# Run tests with race detector
go test -race ./...

# Build all packages (verify compilation)
go build -v ./...

# Format code (use goimports, not gofmt)
goimports -w .

# Lint (golangci-lint with default config)
golangci-lint run ./...

# Docker build
docker build -t aklapi .

# Make targets
make server   # go build -o server ./cmd/aklapi
make docker   # docker build -t aklapi .
```

> **To run a single test:** use `go test -v -run <TestName> <./package/path>`
> Example: `go test -v -run TestNextRubbish .`

---

## Code Style Guidelines

### Formatting

- Use **`goimports`** (not plain `gofmt`) — it manages imports automatically.
- Indentation: **tabs** (Go standard).
- VS Code devcontainer is configured with `"editor.formatOnSave": true` using `goimports`.
- No trailing whitespace; no blank lines at end of file.

### Imports

Group imports in two blocks separated by a blank line:
1. Standard library
2. Third-party packages

```go
import (
    "context"
    "encoding/json"
    "net/http"

    "github.com/PuerkitoBio/goquery"
)
```

- Use blank imports only where required: `_ "time/tzdata"`, `_ "embed"`.
- Never use dot imports (`.`).
- Alias imports only when disambiguation is genuinely needed.

### Naming Conventions

| Element | Convention | Example |
|---|---|---|
| Exported types | PascalCase | `AddrRequest`, `RubbishCollection` |
| Unexported types | camelCase | `refuseParser`, `lruCache` |
| Exported functions | PascalCase | `AddressLookup`, `CollectionDayDetail` |
| Unexported functions | camelCase | `fetchandparse`, `oneAddress` |
| Receiver names | Short (1–2 chars) | `(r *RubbishCollection)`, `(c *lruCache[K,V])` |
| Package-level vars | camelCase | `addrCache`, `defaultLoc` |
| Unexported constants | camelCase | `defCacheSz`, `dateLayout` |
| Acronyms | Go convention | `addrURI` (not `addrUrl`), `ID` (not `Id`) |

### Types & Structs

- Add JSON struct tags to all exported response types: `json:"field,omitempty"`.
- Prefer pointer receivers for types that may mutate state or are large.
- Use **generics** for reusable containers (see `lruCache[K comparable, V any]`).
- Use a stateful parser type (struct with fields for state, error, and results) when
  parsing multi-step data (see `refuseParser`).

### Error Handling

- Always check errors: `if err != nil { return nil, err }`.
- No `panic` in production code.
- Use `errors.New("...")` for static error messages.
- Use string concatenation (not `fmt.Sprintf`) for simple dynamic error strings:
  ```go
  errors.New("address API returned status code: " + strconv.Itoa(resp.StatusCode))
  ```
- Prefer `fmt.Errorf("context: %w", err)` for wrapping errors that need context.
- Use package-level sentinel errors for flow control:
  ```go
  var errSkip = errors.New("skip this date")
  ```
- Use `errors.Is` for sentinel error comparisons.
- HTTP handlers: use `http.Error(w, msg, code)` or a typed `respond(w, body, code)` helper.

### HTTP & Networking

- Use **standard library `net/http` only** — no external router (no Gin, Echo, Chi).
- Register routes with `http.HandleFunc` on the default mux.
- Always pass context to outgoing HTTP requests:
  ```go
  req, err := http.NewRequestWithContext(ctx, http.MethodGet, uri, nil)
  ```
- Always `defer resp.Body.Close()` immediately after a successful response.
- Decode JSON responses with `json.NewDecoder(resp.Body).Decode(&v)`.

### Logging

- Use **`log/slog`** for all logging — not `log.Printf`, `fmt.Println`, etc.
- Prefer context-aware variants: `slog.DebugContext(ctx, ...)`, `slog.InfoContext(ctx, ...)`.
- Add structured key-value pairs for observability:
  ```go
  start := time.Now()
  // ... operation ...
  slog.DebugContext(ctx, "fetched addresses", "count", len(results), "duration", time.Since(start))
  ```

### Dependency Injection & Testability

- Declare external URLs as **package-level `var`** (not `const`) so tests can override them:
  ```go
  var addrURI = `https://example.com/api/addresses`
  ```
- Inject time via a replaceable variable: `var now = time.Now`.
- Injectable function-type variables enable handler testing without real upstream calls:
  ```go
  var addressLookup = aklapi.AddressLookup
  ```
- Restore overridden vars with `defer`:
  ```go
  old := addrURI
  addrURI = ts.URL
  defer func() { addrURI = old }()
  ```

---

## Testing Guidelines

### Style

- Use **table-driven tests** for all non-trivial functions.
- Table entry struct fields: `name string`, `args`, `want`, `wantErr bool`.
- Field names may be omitted for the `name` field in composite literals.
- Prefer `github.com/stretchr/testify/assert` for assertions in new tests (avoid raw
  `reflect.DeepEqual` + `t.Errorf` patterns from older tests).
- Use `t.Context()` (Go 1.24+) for context in subtests.
- Use `t.Cleanup(func() {...})` for teardown instead of `defer` in the test function body
  when working with subtests.

### HTTP Testing

- Use `net/http/httptest.NewServer` to mock upstream APIs.
- Use `httptest.NewRequest` + `httptest.NewRecorder` for handler unit tests.

### Test Fixtures

- Embed HTML fixture files with `//go:embed`:
  ```go
  //go:embed test_assets/some-page.html
  var fixtureHTML []byte
  ```
- Fixtures are refreshed by `//go:generate` directives that `curl` the live page.

### Subtests

- Always run subtests with `t.Run(tt.name, func(t *testing.T) { ... })`.
- Use `t.Helper()` in assertion helper functions.

---

## Repository Conventions

- **One concern per file:** `addr.go`, `rubbish.go`, `caches.go`, `time.go`.
- **Library in root, binary in `cmd/`:** follows standard Go project layout.
- CI runs on `push` and `pull_request` to `master` (see `.github/workflows/go.yml`):
  `go build -v ./...` then `go test -v ./...`.
- Docker images are published to `ffffuuu/aklapi` on GitHub Release events.
