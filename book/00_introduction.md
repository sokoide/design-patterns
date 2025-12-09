# Introduction: Go Design Patterns ~A Journey to Clean Architecture~

Welcome to the world of design patterns!
This book is not just for memorizing "design patterns."
It is an adventure guide for writing "code that is robust to change and easy to test,"
leveraging the characteristics of the **Go language** and incorporating the philosophy of **Clean Architecture**.

## 🎯 Purpose of This Book

1. **Understand the "Why"**: Learn the background of why patterns were created and the problems they solve.
2. **Write Idiomatic Go**: Master implementations that follow Go's idioms (Interfaces, Structs, Embedding), rather than direct translations from Java or C++.
3. **Experience Clean Architecture**: Practice designing systems that separate business logic (Domain) from technical details (Adapter).

## 👥 Target Audience

- Engineers who understand the **basics of the Go language** (syntax, structs, interfaces).
- Intermediate developers who are struggling with "spaghetti code."
- Those who are interested in architectural design but don't know where to start.

## 🗺️ How to Navigate This Book

This book is structured to conquer the 23 GoF design patterns in **5 days**.

| Day | Theme | Patterns Covered | Goal |
| :--- | :--- | :--- | :--- |
| **Day 1** | [**Fundamentals of Object Creation** (Creational)](./01_day1_creational.md) | Singleton, Factory Method, Abstract Factory, Builder, Prototype | Learn how to separate object creation from its use. |
| **Day 2** | [**Organizing and Extending Structure** (Structural 1)](./02_day2_structural_1.md) | Adapter, Bridge, Composite, Decorator | Learn how to integrate different interfaces and dynamically add functionality. |
| **Day 3** | [**Efficiency and Interfaces** (Structural 2 & Behavioral Intro)](./03_day3_structural_2_behavioral_1.md) | Facade, Flyweight, Proxy, Strategy | Learn to hide complexity, improve resource efficiency, and switch algorithms (Strategy). |
| **Day 4** | [**Separating Behavior and Notification** (Behavioral 1)](./04_day4_behavioral_2.md) | Observer, Command, State, Template Method, Iterator | Learn about state change notifications, encapsulating processes, and state management. |
| **Day 5** | [**Complex Flows and Responsibilities** (Behavioral 2)](./05_day5_behavioral_3.md) | Chain of Responsibility, Mediator, Memento, Visitor, Interpreter | Learn about complex interactions between objects, chaining responsibilities, and separating data structures from processing. |

Each chapter (Day) consists of the following sections:

1. **Story**: Intuitively understand difficult concepts through relatable analogies.
2. **Concept**: Visualize the pattern's definition and structure with Mermaid diagrams.
3. **Go Implementation Tips**: Explain key implementation points using code from this repository as examples.
4. **Hands-on**: Small exercises to deepen understanding by getting your hands dirty.
5. **Quiz**: A multiple-choice quiz to check your understanding.

## 🛠️ What You'll Need

- Go 1.18 or higher
- Your favorite editor (VS Code, GoLand, Vim, etc.)
- Curiosity and a playful spirit

So, are you ready?
Let's open the door to Day 1 and learn the magic of object creation!
