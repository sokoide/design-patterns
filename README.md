# Go Design Patterns (Clean Architecture)

[日本語版はこちら](./README_ja.md)

This repository is a collection of Gang of Four design pattern samples built with **Go** using Clean Architecture principles. It helps you learn how to separate business logic (Usecase) from implementation details (Adapter) and keep designs loosely coupled.

See [the book](./book/00_introduction.md) for a deeper explanation.

## 📂 Project Structure

Each pattern lives in its own directory (a Go module).

### 1. Creational Patterns

These patterns make object creation more flexible.

- [**Abstract Factory**](./abstract-factory-example) (`abstract-factory-example`): Produces families of related objects without specifying their concrete classes.
- [**Builder**](./builder-example) (`builder-example`): Encapsulates complex construction so the same process can yield different representations.
- [**Factory Method**](./factory-example) (`factory-example`): Delegates object creation to subclasses (or implementations).
- [**Prototype**](./prototype-example) (`prototype-example`): Creates new objects by cloning a prototypical instance.
- [**Singleton**](./singleton-example) (`singleton-example`): Ensures a class has only one instance.

### 2. Structural Patterns

These patterns compose objects and classes into larger structures.

- [**Adapter**](./adapter-example) (`adapter-example`): Makes incompatible interfaces work together.
- [**Bridge**](./bridge-example) (`bridge-example`): Separates abstractions from their implementations so each can evolve on its own.
- [**Composite**](./composite-example) (`composite-example`): Treats part-whole hierarchies uniformly, enabling recursive tree structures.
- [**Decorator**](./decorator-example) (`decorator-example`): Dynamically adds responsibilities to objects.
- [**Facade**](./facade-example) (`facade-example`): Provides a simple interface to a complex subsystem.
- [**Flyweight**](./flyweight-example) (`flyweight-example`): Shares instances efficiently when many similar objects are needed.
- [**Proxy**](./proxy-example) (`proxy-example`): Controls access to another object, such as for access control or lazy initialization.

### 3. Behavioral Patterns

These patterns concern communication and responsibility between objects.

- [**Chain of Responsibility**](./chain-of-responsibility-example) (`chain-of-responsibility-example`): Passes a request along a chain until an object handles it.
- [**Command**](./command-example) (`command-example`): Encapsulates requests as objects for history tracking, undo, etc.
- [**Interpreter**](./interpreter-example) (`interpreter-example`): Defines grammar rules as classes to interpret a language.
- [**Iterator**](./iterator-example) (`iterator-example`): Traverses a collection without exposing its internal structure.
- [**Mediator**](./mediator-example) (`mediator-example`): Centralizes interactions between multiple objects to reduce coupling.
- [**Memento**](./memento-example) (`memento-example`): Captures and restores an object’s state.
- [**Observer**](./observer-example) (`observer-example`): Notifies dependent objects about state changes.
- [**State**](./state-example) (`state-example`): Changes behavior when an object’s internal state changes.
- [**Strategy**](./strategy-example) (`strategy-example`): Makes algorithms interchangeable.
- [**Template Method**](./template-method-example) (`template-method-example`): Defines the skeleton of an algorithm and delegates specific steps to subclasses.
- [**Visitor**](./visitor-example) (`visitor-example`): Separates algorithms from the data structures they operate on.

### Other Patterns

- [**Functional Options**](./functional-options-example): A Go-specific pattern for flexible struct initialization.
- [**Entitlement (Gateway)**](./entitlement-example): An example of authority management using Clean Architecture. It uses a Gateway (= Adapter) pattern to abstract access to external resources (AD/Cache).

## 🏗 Common Architecture

Each directory follows a Clean Architecture structure:

- **`domain/`**: Defines core business rules, models, and interfaces. It has no dependencies on external libraries.
- **`usecase/`**: Implements application-specific business logic using the `domain` interfaces.
- **`adapter/`**: Contains concrete implementations (e.g., repositories, API clients, algorithms) of `domain` interfaces.
- **`main.go`**: The entry point that wires dependencies via DI and runs the components.

## 🚀 Running Examples

Move into a pattern directory and run `go run main.go`.

```bash
# Example: Run the Factory Method pattern
cd factory-example
go run main.go
```
