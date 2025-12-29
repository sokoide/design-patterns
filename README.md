# Go Design Patterns (Clean Architecture)

[日本語版はこちら](./README_ja.md)

This repository is a collection of Gang of Four design pattern samples implemented in **Go**.
Each example follows Clean Architecture ideas, separating business logic from implementation details.

- [Japanese README](./README_ja.md)
- [Detailed explanation](./book/00_introduction.md)

## Goals

- Understand the intent and use cases of GoF patterns through Go-style type/interface design
- Learn Clean Architecture layering (`domain` / `usecase` / `adapter`) and dependency injection (DI) by example

## Usage

### Run a single example

Each `*-example` directory is an independent Go module. Move into one and run it.

```bash
cd strategy-example
go run main.go
```

### Run all examples

From the repo root, this runs every module sequentially.

```bash
make run
```

### Tests

Run tests in a specific module or across the whole repo.

```bash
cd strategy-example
go test ./...

# From the repo root
go test ./...
```

## Common Directory Structure

Each example typically follows this Clean Architecture layout:

- `domain/`: Pure domain layer. Defines domain models and interfaces with no external dependencies.
- `usecase/`: Application layer. Implements use cases using `domain` interfaces.
- `adapter/`: Concrete implementations (external I/O, algorithms, repositories, etc.).
- `main.go`: Entry point. Wires dependencies and runs the app.

## Pattern List

### 1. Creational Patterns

Patterns that make object creation more flexible. Many of these are omitted because standard constructors are often sufficient in Go, or they are less idiomatic.

- **Builder** (`builder-example`): Separates complex construction steps.
  - **Note**: In Go, we usually use **Functional Options**, so this is omitted.
- **Abstract Factory** (`abstract-factory-example`): Switches between families of related products.
  - **Note**: Rarely used in practice. Omitted.
- **Factory Method** (`factory-example`): Delegates creation to implementations.
  - **Note**: Not very common in Go. `NewServer(cfg Config)` style constructors are usually sufficient.
- **Singleton** (`singleton-example`): Manages a single shared instance.
  - **Note**: **Avoid**. It hinders testability and can cause concurrency issues.
- **Prototype** (`prototype-example`): Creates new objects by cloning.
  - **Note**: Since copying structs is cheap (`b := a`), this is omitted.

### 2. Structural Patterns

Patterns for composing objects and classes into larger structures.

- [**Adapter**](./adapter-example) (`adapter-example`): Connects incompatible interfaces.
- [**Decorator**](./decorator-example) (`decorator-example`): Layers additional behavior dynamically.
- [**Composite**](./composite-example) (`composite-example`): Treats parts and wholes uniformly.
- [**Facade**](./facade-example) (`facade-example`): Provides a simple façade over a complex subsystem.
- [**Proxy**](./proxy-example) (`proxy-example`): Controls access or delays initialization.
- [**Flyweight**](./flyweight-example) (`flyweight-example`): Shares state to handle many objects efficiently.

The following is ignored:

- **Bridge** (`bridge-example`): Separates abstraction and implementation.
  - **Note**: Go's `interface` makes this pattern largely unnecessary.

### 3. Behavioral Patterns

Patterns concerning communication and responsibility between objects.

- [**Strategy**](./strategy-example) (`strategy-example`): Swaps algorithms at runtime.
- [**Observer**](./observer-example) (`observer-example`): Notifies subscribers on state changes.
- [**Command**](./command-example) (`command-example`): Encapsulates operations for history, undo, etc.
- [**Chain of Responsibility**](./chain-of-responsibility-example) (`chain-of-responsibility-example`): Passes requests along a chain.
- [**State**](./state-example) (`state-example`): Changes behavior by switching internal state.
- [**Memento**](./memento-example) (`memento-example`): Saves and restores state snapshots.

The following are ignored:

- **Iterator** (`iterator-example`): Traverses collections.
  - **Note**: Go has `for range`, and custom iterators usually just need a `Next() (T, bool)` method.
- **Mediator** (`mediator-example`): Centralizes interactions.
  - **Note**: Often leads to "God Objects".
- **Template Method** (`template-method-example`): Shares an algorithm skeleton.
  - **Note**: Go lacks inheritance. Use functions, small interfaces, and composition instead.
- **Visitor** (`visitor-example`): Separates operations from data structures.
  - **Note**: Usually replaced by `switch n := n.(type)` (type switches).
- **Interpreter** (`interpreter-example`): Models grammar rules.
  - **Note**: Skipped in this collection.

### Other Patterns

Commonly used in Go.

- [**Functional Options**](./functional-options-example): A Go‑idiomatic pattern for flexible initialization.

## Notes

- These examples are minimal for learning purposes. Adjust the designs for real‑world requirements.
- See [book](./book/00_introduction.md) for background and diagrams.
