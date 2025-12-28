# Product Guidelines

## Code Style
- **Idiomatic Go:** Adhere strictly to `gofmt` standards and "Effective Go" guidelines.
- **Naming Conventions:** Use clear, descriptive names. Exported identifiers (Capitalized) should have comments.
- **Error Handling:** Handle errors explicitly. Avoid panics in library code.

## Documentation
- **Bilingual Support:** maintain `README.md` (English) and `README_ja.md` (Japanese) for each pattern to ensure accessibility for both audiences.
- **Clarity:** Explanations should be concise and focused on the "why" and "how" of the pattern.

## Visuals
- **Mermaid Diagrams:** Use Mermaid syntax within Markdown for Class Diagrams and Sequence Diagrams to illustrate pattern structure and flow.

## Testing
- **Standard Library:** Use Go's built-in `testing` package.
- **Table-Driven Tests:** Prefer table-driven tests for covering multiple scenarios efficiently.
- **Coverage:** Aim for high test coverage for the core pattern logic.

## Architecture & Structure
- **Clean Architecture:** Each example directory MUST follow the layered structure:
    - `domain/`: Business entities and interfaces.
    - `usecase/`: Application logic.
    - `adapter/`: Interface adapters (implementations).
- **Independence:** Each pattern directory must be a self-contained Go module.
