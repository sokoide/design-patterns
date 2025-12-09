# Design Patterns in Go (Clean Architecture)

This repository contains a collection of educational examples demonstrating standard software design patterns implemented in **Go**. Each example is structured according to **Clean Architecture** principles to illustrate how to decouple business logic from implementation details.

## 📂 Project Structure

The repository is organized by design pattern. Each directory is a self-contained Go module.

### Available Patterns

*   `command-example/`: Command Pattern
*   `decorator-example/`: Decorator Pattern
*   `entitlement-example/`: (Likely Proxy or Strategy related to permissions)
*   `facade-example/`: Facade Pattern
*   `factory-example/`: Factory Method Pattern
*   `functional-options-example/`: Functional Options Pattern
*   `observer-example/`: Observer Pattern
*   `state-example/`: State Pattern
*   `strategy-example/`: Strategy Pattern

### Common Internal Structure

Each pattern directory typically follows this Clean Architecture layout:

*   **`domain/`**: Contains core business entities and interfaces (e.g., `Repository` interfaces, `Model` structs). This layer has no external dependencies.
*   **`usecase/`**: Contains application business rules. It orchestrates data flow to and from the entities, using interfaces defined in the domain.
*   **`adapter/`**: Contains interface adapters (implementations). This is where concrete implementations of domain interfaces reside (e.g., specific notification services, database drivers).
*   **`infrastructure/`** (optional): Frameworks and drivers (e.g., database connections, external API clients).
*   **`main.go`**: The entry point. It is responsible for **dependency injection**, wiring up the adapters and use cases, and starting the application.

## 🚀 usage

Each example is a standalone Go module. To run a specific pattern:

1.  Navigate to the pattern's directory:
    ```bash
    cd factory-example
    ```

2.  Run the `main.go` file:
    ```bash
    go run main.go
    ```

## 🛠 Development Conventions

*   **Dependency Injection**: Dependencies are injected via constructor functions (e.g., `NewService(repo Repository)`), typically wired up in `main.go`.
*   **Interfaces**: Key behaviors are defined as interfaces in the `domain` layer to invert dependencies.
*   **Language**: Go (Golang).
*   **Documentation**: Each directory contains a `README.md` (in Japanese) explaining the specific scenario and architecture, often including Mermaid diagrams.
