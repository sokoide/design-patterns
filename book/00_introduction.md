# Introduction: Go Design Patterns - A Journey to Clean Architecture

Welcome to the world of design patterns!
This book is not just about memorizing "Design Patterns."
It is an adventure book for writing "change-resistant and testable code" by leveraging the characteristics of the **Go language** and incorporating the philosophy of **Clean Architecture**.

## 🎯 Goals of This Book

1.  **Understand the "Why"**: Learn the background of why patterns were born and the problems they solve.
2.  **Write Idiomatic Go**: Master implementations that follow Go idioms (Interfaces, Structs, Embedding) rather than literal translations from Java or C++.
3.  **Experience Clean Architecture**: Practice designs that decouple business logic (Domain) from technical details (Adapter).

## 👥 Target Audience

- Engineers who understand the **basics of Go** (syntax, structs, interfaces).
- Intermediate developers worried that their "code is becoming spaghetti..."
- Those interested in architectural design but unsure where to start.

## 🗺️ How to Walk Through This Book

This book is organized to efficiently learn the design patterns that are truly important in Go.

| Day       | Theme                                                                                                     | Target Patterns                                                  | Goal                                                                           |
| :-------- | :-------------------------------------------------------------------------------------------------------- | :--------------------------------------------------------------- | :----------------------------------------------------------------------------- |
| **Day 1** | [**Idiomatic Initialization** (Initialization)](./01_day1_creational.md)                                   | **Functional Options**                                           | Learn the most frequently used flexible initialization pattern in Go.          |
| **Day 2** | [**Organizing & Extending Structure** (Structural 1)](./02_day2_structural_1.md)                          | Adapter, Decorator, Composite                                    | Learn how to integrate incompatible interfaces and add features dynamically.   |
| **Day 3** | [**Hiding Complexity & Control** (Structural 2)](./03_day3_structural_2.md)                                  | Facade, Proxy, Flyweight                                         | Learn to organize complex subsystems, control access, and optimize resources.  |
| **Day 4** | [**Algorithms & Notification** (Behavioral 1)](./04_day4_behavioral_1.md)                                 | Strategy, Observer, Command                                      | Learn to swap algorithms, notify state changes, and encapsulate operations.    |
| **Day 5** | [**Managing State & Responsibility** (Behavioral 2)](./05_day5_behavioral_2.md)                           | Chain of Responsibility, State, Memento                          | Learn to process via chains, manage state transitions, and save/restore state. |

Each chapter (Day) consists of the following sections:

1.  **Story**: Intuitive understanding of difficult concepts through familiar analogies.
2.  **Concept**: Visualization of the pattern's definition and structure using Mermaid diagrams.
3.  **The Essence of Go Implementation**: Explanation of implementation points using code from this repository.
4.  **Hands-on**: Small tasks to deepen understanding by actually moving your hands.
5.  **Quiz**: A 3-choice quiz to check your understanding.

## 🛠️ Prerequisites

- Go 1.18 or higher
- Your favorite editor (VS Code, GoLand, Vim, etc.)
- Curiosity and a sense of playfulness

Are you ready?
Open the door to Day 1 and go learn the magic of idiomatic initialization!