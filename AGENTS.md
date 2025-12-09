# Repository Guidelines

## Project Structure & Module Organization
- Each design pattern lives in its own Go module directory (e.g., `factory-example`, `strategy-example`). Explore `domain/`, `usecase/`, and `adapter/` subfolders inside a pattern to see Clean Architecture layers in action.
- `book/` contains supporting documentation; treat it as read-only background material.
- `Makefile` collects the example subdirectories via `DIRS = $(sort $(wildcard *-example))` and ships a `run` target that iterates through all modules.

## Build, Test, and Development Commands
- `go run .` or `go run main.go` inside a pattern directory to execute that example and exercise its DI wiring.
- `make run` to sequentially execute every example module; useful when validating no example crashes when wired end-to-end.
- `go test ./...` run inside a pattern module (or from the repo root with `./...`) to exercise unit tests for that example.

## Coding Style & Naming Conventions
- Follow idiomatic Go: format code with `gofmt`/`go fmt`, keep package names lowercase without underscores, and keep exported identifiers capitalized.
- Domain models and use cases live in `domain`/`usecase` packages; the concrete implementations sit under `adapter`. Respect those folder boundaries when adding new code.
- Keep README and documentation bilingual (Japanese/English) concise; add code comments that clarify business intent rather than describe obvious mechanics.

## Testing Guidelines
- No centralized testing framework; rely on Go’s built-in `testing` package.
- Name test files with `_test.go` and tests with `Test...` to match Go tooling expectations.
- Run `go test ./...` before submitting to verify each module’s public interfaces and use cases behave.

## Commit & Pull Request Guidelines
- Recent commits follow `type: summary` style (e.g., `add: book`, `update: README`); keep the prefix lowercase and concise.
- PR descriptions should explain what patterns changed, link to any tracked issues, and note whether you ran `go test ./...` or `make run`.
- Include screenshots or logs only if the change affects output examples or documentation; otherwise focus on text descriptions.
