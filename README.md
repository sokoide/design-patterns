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
# Example: Run the Factory Method pattern
cd factory-example
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

Each pattern lives in its own directory (a Go module).

### 1. Creational Patterns

These patterns make object creation more flexible.

- [**Abstract Factory**](./abstract-factory-example) (`abstract-factory-example`): Switches between families of related products without depending on concrete types.
- [**Builder**](./builder-example) (`builder-example`): Separates complex construction steps so one process can build different representations.
- [**Factory Method**](./factory-example) (`factory-example`): Delegates creation to implementations, reducing the caller’s dependency on concretes.
- [**Prototype**](./prototype-example) (`prototype-example`): Creates new objects by cloning existing instances.
- [**Singleton**](./singleton-example) (`singleton-example`): Manages a single shared instance.

### 2. Structural Patterns

These patterns compose objects and classes into larger structures.

- [**Adapter**](./adapter-example) (`adapter-example`): Connects incompatible interfaces.
- [**Bridge**](./bridge-example) (`bridge-example`): Separates abstraction and implementation so both can evolve independently.
- [**Composite**](./composite-example) (`composite-example`): Treats parts and wholes uniformly to model tree structures.
- [**Decorator**](./decorator-example) (`decorator-example`): Layers additional behavior dynamically.
- [**Facade**](./facade-example) (`facade-example`): Provides a simple façade over a complex subsystem.
- [**Flyweight**](./flyweight-example) (`flyweight-example`): Shares state to handle many similar objects efficiently.
- [**Proxy**](./proxy-example) (`proxy-example`): Uses a stand‑in to control access or delay initialization.

### 3. Behavioral Patterns

These patterns concern communication and responsibility between objects.

- [**Chain of Responsibility**](./chain-of-responsibility-example) (`chain-of-responsibility-example`): Passes requests along a chain until one handler accepts them.
- [**Command**](./command-example) (`command-example`): Encapsulates operations for history, undo, etc.
- [**Interpreter**](./interpreter-example) (`interpreter-example`): Models grammar rules and evaluates expressions.
- [**Iterator**](./iterator-example) (`iterator-example`): Traverses collections without exposing internals.
- [**Mediator**](./mediator-example) (`mediator-example`): Centralizes interactions to simplify dependencies.
- [**Memento**](./memento-example) (`memento-example`): Saves and restores object state snapshots.
- [**Observer**](./observer-example) (`observer-example`): Notifies subscribers on state changes.
- [**State**](./state-example) (`state-example`): Changes behavior by switching internal state.
- [**Strategy**](./strategy-example) (`strategy-example`): Swaps algorithms at runtime.
- [**Template Method**](./template-method-example) (`template-method-example`): Shares an algorithm skeleton and lets implementations fill in steps.
- [**Visitor**](./visitor-example) (`visitor-example`): Separates operations from data structures to add new behavior without changing them.

### Other Patterns

- [**Functional Options**](./functional-options-example): A Go‑idiomatic pattern for flexible initialization.
- [**Entitlement (Gateway)**](./entitlement-example): Authority management example. Uses a Gateway (= Adapter) to abstract external resources (AD/Cache).

## Notes

- These examples are minimal for learning purposes. Adjust the designs for real‑world requirements.
- See `book/` for background and diagrams.
