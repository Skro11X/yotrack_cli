# Repository Guidelines

## Project Structure & Module Organization

This repository is currently a minimal Go project. Source code lives at the
repository root, with `main.go` as the current entry point. As the project
grows, keep executable entry points in `cmd/<name>/`, reusable packages in
`internal/` or `pkg/`, and tests next to the code they cover using Go's
standard `*_test.go` naming.

Suggested future layout:

```text
main.go              # current entry point
internal/            # private application packages
cmd/<tool>/          # command entry points, if multiple tools are added
testdata/            # fixtures used by Go tests
```

## Build, Test, and Development Commands

Use standard Go tooling. If a `go.mod` file is added, run commands from the
repository root.

- `go run .` runs the current main package locally.
- `go test ./...` runs all Go tests in the repository.
- `go build ./...` verifies all packages compile.
- `gofmt -w .` formats Go source files in place.

If the repository remains a single-file experiment, prefer `go run main.go`
for quick local checks.

## Coding Style & Naming Conventions

Follow idiomatic Go. Use tabs as produced by `gofmt`, keep package names short
and lowercase, and prefer descriptive exported identifiers with comments when
they become part of a public API. File names should be lowercase and may use
underscores for clarity, for example `youtrack_client.go`.

Keep parsing, API, and CLI concerns separated as the codebase grows. Avoid
large `main.go` implementations once behavior becomes reusable or testable.

## Testing Guidelines

Use Go's built-in `testing` package unless a clear need for another framework
appears. Place tests beside the implementation and name them after behavior,
for example `TestParseIssueList`. Store fixtures under `testdata/` so Go tools
ignore them during builds.

Run `go test ./...` before submitting changes. Add table-driven tests for
parsing logic and edge cases.

## Commit & Pull Request Guidelines

No usable Git history is available in this checkout, so use clear conventional
commit messages such as `feat: add YouTrack parser` or `fix: handle empty
responses`.

Pull requests should include a short problem summary, the main implementation
changes, test results, and any relevant sample input or output. Link related
issues when available.

## Agent-Specific Instructions

Before editing, check whether generated guidance or configuration files already
exist and preserve user-created content. Keep changes narrow, run the relevant
Go commands when possible, and document any command that cannot be run.
