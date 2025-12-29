# Go State Pattern Example (Clean Architecture)

This project is an educational sample code that implements the **State Pattern** using the **Go** language. Following the structure of Clean Architecture, it separates the logic of state transitions.

## What This Example Shows

- Encapsulating state‑specific behavior in separate state objects
- Delegating transitions from the context to the current state
- Using `usecase` to coordinate the flow, utilizing a `Logger` for output (avoiding direct `fmt` calls in business logic)

## Quick Start

In the `state-example` directory:

```bash
go run main.go
```

## 🚪 Scenario: Door Opening and Closing System

A door has three states, and its state transitions through two actions (A and B).

### States

1. **Locked**: The door is locked.
2. **ClosedUnlocked**: The door is closed but unlocked.
3. **Open**: The door is open.

### Actions

* **Action A**:
  * When Locked → **Unlock**
  * When ClosedUnlocked → **Open**
  * When Open → Do nothing
* **Action B**:
  * When Open → **Close**
  * When ClosedUnlocked → **Lock**
  * When Locked → Do nothing

## 🏗 Architecture

```mermaid
classDiagram
    namespace domain {
        class Action {
            <<enumeration>>
            A
            B
        }
        class Logger {
            <<interface>>
            +Log(message: string)
        }
        class DoorState {
            <<interface>>
            +Name() string
            +Handle(action Action) (DoorState, string, error)
        }
    }

    namespace usecase {
        class DoorContext {
            -currentState: DoorState
            -logger: Logger
            +ExecuteAction(action Action)
            +GetStateName() string
        }
    }

    namespace adapter {
        class LockedState {
            +Handle(action Action)
        }
        class ClosedUnlockedState {
            +Handle(action Action)
        }
        class OpenState {
            +Handle(action Action)
        }
        class ConsoleLogger {
            +Log(message: string)
        }
    }

    %% Relationships
    DoorContext o-- DoorState : Holds Current
    DoorContext o-- Logger : Uses
    LockedState ..|> DoorState : Implements
    ClosedUnlockedState ..|> DoorState : Implements
    OpenState ..|> DoorState : Implements
    ConsoleLogger ..|> Logger : Implements

    %% Transitions (Conceptual)
    LockedState ..> ClosedUnlockedState : Action A
    ClosedUnlockedState ..> OpenState : Action A
    OpenState ..> ClosedUnlockedState : Action B
    ClosedUnlockedState ..> LockedState : Action B
```

### Role of Each Layer

1. **Domain (`/domain`)**:
    * `DoorState` interface: Defines the behavior that all states must have (`Handle`).
    * `Action` constant: The common language used within the system.
    * `Logger` interface: Abstract logging capability.
2. **Usecase (`/usecase`)**:
    * **Context (`DoorContext`)**: A container that holds the current state (`currentState`).
    * When it receives user input, it doesn't make decisions itself but delegates the task to the current state (`currentState.Handle`).
    * It uses `domain.Logger` to output results, ensuring no direct dependency on `fmt` or external systems.
3. **Adapter (`/adapter`)**:
    * **Concrete States**: A place for the specific logic of each state, such as `LockedState` and `OpenState`.
    * The **transition rules**, like "when in the Locked state and button A is pressed, the next state is ClosedUnlocked," are described here.
    * **ConsoleLogger**: Concrete implementation of the logger.

## 💡 Architectural Design Notes (Q&A)

For engineers learning Clean Architecture, this section explains the important intentions behind this design.

### Q1. Where should the state transition rules (if/switch statements) be written?

**A. In the State Pattern, they are written inside each "State class (Adapter)."**

If you write a giant `switch` statement (`if state == Locked then ...`) in the `Usecase`, you would need to modify that huge function every time a new state is added, making it a breeding ground for bugs.
In the State Pattern, the "behavior in the Locked state" is encapsulated within the `LockedState` class. This makes the code for each state independent and easier to understand.

### Q2. How well does it fit with Clean Architecture?

**A. It fits very well.**

* **Domain**: Defines the state interface.
* **Adapter**: Implements the concrete state transition logic.
* **Usecase**: Simply holds and uses the state.

This clean separation of responsibilities makes it a very effective design for creating applications with complex state machines (like games, workflow engines, payment flows, etc.).

## 🚀 How to Run

```bash
go run main.go
```