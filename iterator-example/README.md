# Go Iterator Pattern Example (Clean Architecture)

This project is an educational sample code that implements the **Iterator Pattern** using the **Go** language. You will learn how to access the elements of a collection sequentially without exposing its underlying representation.

## What This Example Shows

- Traversing a collection without exposing its internal slice
- Implementing a simple iterator with `HasNext` / `GetNext`

## Quick Start

In the `iterator-example` directory:

```bash
go run main.go
```

## 🔄 Scenario: Traversing a User List

We have a `UserCollection` that holds a list of `User`s.
The client (`main.go`) does not want to know whether this collection uses an array, a linked list, or a map internally.
It just wants to say, "Give me the next element (`Next`)." We will achieve this using an Iterator.

### Characters

1. **Iterator (`domain.Iterator`)**: The interface for traversal, with methods like "Is there a next element? (`HasNext`)" and "What is the next element? (`GetNext`)."
2. **Aggregate (`domain.Collection`)**: An interface for creating an Iterator (`CreateIterator`).
3. **Concrete Iterator (`adapter.UserIterator`)**: The class that actually manages the array index and performs the traversal.
4. **Concrete Aggregate (`adapter.UserCollection`)**: The actual data holder.

## 🏗 Architecture

```mermaid
classDiagram
    direction TB

    %% Domain Layer
    class Iterator {
        <<interface>>
        +HasNext() bool
        +GetNext() interface{}
    }

    class Collection {
        <<interface>>
        +CreateIterator() Iterator
    }

    %% Adapter Layer
    class UserCollection {
        +Users: User[]
        +CreateIterator() Iterator
    }

    class UserIterator {
        -users: User[]
        -index: int
        +HasNext() bool
        +GetNext() interface{}
    }

    %% Relationships
    UserCollection ..|> Collection : Implements
    UserIterator ..|> Iterator : Implements
    UserCollection ..> UserIterator : Creates
```

### Role of Each Layer

1. **Domain (`/domain`)**:
    * `Iterator`: Defines the standard API for traversal.
    * `Collection`: Defines something that provides an iterator.
2. **Adapter (`/adapter`)**:
    * `UserCollection`: A concrete list of users.
    * `UserIterator`: An iterator specific to `UserCollection`. It holds the current progress (`index`) as its state.

## 💡 Architectural Design Notes (Q&A)

### Q1. Go has `for range` loops. What's the point of using this?

**A. It's effective when you want to hide (encapsulate) the internal structure of a collection.**

The standard `range` can only be used for Go's built-in types like slices and maps.
For complex data structures like "traversing a tree structure with depth-first search" or "traversing a graph structure," implementing the Iterator pattern allows the user to handle them uniformly with just `HasNext() / GetNext()`, without worrying about the internal algorithm.

### Q2. `GetNext()` returns `interface{}`. What about type safety?

**A. Type safety is lost in this implementation example (pre-generics Go style).**

Using **Generics** from Go 1.18 onwards, it's possible to create a type-safe iterator like `Iterator[T]`.
This example uses `interface{}` to mimic the classic structure of the GoF pattern, but in practice, using generics is recommended.

## 🚀 How to Run

```bash
go run main.go
```
