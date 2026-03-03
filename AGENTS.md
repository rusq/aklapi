# AGENTS.md

## Project Snapshot
- Module: `github.com/rusq/aklapi`
- Language: Go
- Purpose: unofficial Auckland Council API wrapper/service for address lookup and rubbish/recycling collection dates.

## Repository Map
- `cmd/`: executable entrypoints.
- `addr.go`: address lookup logic.
- `rubbish.go`: rubbish/recycling API logic and response shaping.
- `caches.go`: cache helpers.
- `time.go`: date/time helpers.
- `*_test.go`: unit tests.
- `test_assets/`: fixtures used by tests.

## Development Commands
- Run tests: `go test ./...`
- Run focused tests: `go test ./... -run <Name>`
- Tidy dependencies: `go mod tidy`
- Build all packages: `go build ./...`

## Working Conventions
- Prefer small, targeted changes over broad refactors.
- Keep API behavior backward compatible unless explicitly requested.
- Add or update tests for behavioral changes.
- Keep exported identifiers and package-level docs concise.

## Validation Checklist
Before finishing a code change, run:
1. `go test ./...`
2. `go build ./...`

If a change only affects docs or comments, note that tests/build were not required.
