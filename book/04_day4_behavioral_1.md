# Day 4: Algorithms & Notification (Behavioral Patterns Part 1)

Welcome to Day 4!
Yesterday, we learned how to make complex structures look simple.
From today, we move to "Behavior," meaning how objects cooperate with each other and change dynamically.
By mastering these, your code will be able to flexibly swap algorithms and automatically communicate state changes.

Today, we will learn the following three patterns:

1.  **Strategy**: Swapping algorithms like cassettes
2.  **Observer**: "Let me know when it's updated!"
3.  **Command**: Boxing up a command

---

## 7. Strategy

### 📖 Story: Weapons in an RPG

A hero fights monsters.
If they equip a "sword," it's a "slashing" attack; if they equip a "bow," it's a "shooting" attack.
Without changing the hero (Context) themselves, you can switch the attack method (algorithm) just by changing the equipment (Strategy).

### 💡 Concept

Encapsulates algorithms and makes them interchangeable at runtime.

```mermaid
classDiagram
    class Context {
        -strategy Strategy
        +Execute()
    }
    class Strategy {
        <<interface>>
        +Algorithm()
    }
    class ConcreteStrategyA {
        +Algorithm()
    }
    Context o-- Strategy
    Strategy <|.. ConcreteStrategyA
```

### 🐹 The Essence of Go Implementation

This is one of the most basic and powerful ways to use interfaces in Go.
Implementing it as a function type (`type StrategyFunc func()`) and passing the function itself is also simple and very "Go-like."
Passing a comparison function to `sort.Slice` is also a type of Strategy pattern.

```go
type PaymentMethod interface {
    Pay(amount float64) error
}

type CreditCard struct {}
func (c *CreditCard) Pay(amount float64) error { ... }

type PayPal struct {}
func (p *PayPal) Pay(amount float64) error { ... }
```

### 🧪 Hands-on

In `strategy-example` (payment and shipping example), try adding a new payment method (e.g., `BankTransfer`) and confirm that switching it at runtime changes the behavior.

### ❓ Quiz

**Q1. What can be avoided by using the Strategy pattern?**
A. Giant `if-else` or `switch` statements.
B. Defining interfaces.
C. Using structs.

<details>
<summary>Answer</summary>
**A**. Since algorithm branching can be expressed by switching classes (or functions), you can avoid a storm of conditional branching.
</details>

---

## 8. Observer

### 📖 Story: Stock Price Alerts

You are monitoring specific stock prices.
When a stock price (Subject) fluctuates, a notification is sent to all investors (Observers) watching that stock.
The stock price system doesn't need to know who is watching in detail. They just send a notification to the "watch list."
If you stop watching, you won't receive notifications anymore.

### 💡 Concept

Automatically notifies other dependent objects when an object's state changes.

```mermaid
classDiagram
    class Subject {
        <<interface>>
        +Register(Observer)
        +Unregister(Observer)
        +NotifyAll()
    }
    class Observer {
        <<interface>>
        +OnUpdate(string)
    }
    Subject --> Observer
```

### 🐹 The Essence of Go Implementation

While implementing it using interfaces is standard in Go, using **Go Channels** allows you to create a more Go-like asynchronous event notification system.
However, it's important to have a mechanism to properly unregister Observers when they are no longer needed to prevent memory leaks.

### 🧪 Hands-on

Let's look at `observer-example`.
Try creating a new type of Observer (e.g., `MobileAppListener`) and registering it with the Subject to receive notifications.

### ❓ Quiz

**Q2. The Observer pattern is the foundation for which architecture?**
A. MVC (Model-View-Controller).
B. REST API.
C. Batch processing.

<details>
<summary>Answer</summary>
**A**. It is used in the heart of MVC as a mechanism to notify the View (display) of changes in the Model (data).
</details>

---

## 9. Command

### 📖 Story: Text Editor Operations

Imagine you are "typing" or "deleting" text in a text editor.
Each operation is objectified as a "command (Command)."
Because it is an "object," you can stack them up to easily implement "Undo" and "Redo."

### 💡 Concept

Encapsulates a request as an object, thereby letting you parameterize clients with different requests, queue or log requests, and support undoable operations.

```mermaid
classDiagram
    class Invoker {
        -command Command
        +SetCommand(Command)
        +ExecuteCommand()
    }
    class Command {
        <<interface>>
        +Do(buffer)
        +Undo(buffer)
    }
    Invoker o-- Command
```

### 🐹 The Essence of Go Implementation

The basic idea is to have a struct with a `Do()` method.
When building CLI tools, implementing each sub-command (`git commit`, `git push`, etc.) using the Command pattern keeps things neatly organized.

### 🧪 Hands-on

`command-example` is a text editor example.
Try adding a new command (e.g., `InsertTextCommand`) and implement `Do` and `Undo` to see the buffer state change.

### ❓ Quiz

**Q3. What is the advantage of the Command pattern?**
A. You can delay the timing of processing or keep a history.
B. You can make the class inheritance hierarchy deep.
C. It increases database speed.

<details>
<summary>Answer</summary>
**A**. Because requests can be treated as "objects," they can be saved or passed around freely.
</details>

---

Great job on Day 4!
You've learned patterns to control dynamic behavior, such as "swapping algorithms" and "notification."
Tomorrow is the final day. We will learn more advanced patterns like managing state and chains of responsibility.
Keep up the good work for the final sprint!