# Polyglot Coding Standards 💻

## 1. General Principles

- **English Only**: Code comments, variable names, database columns must be in English.
- **Clean Code**: Functions should do one thing. DRY (Don't Repeat Yourself).
- **API Response**: Standard JSON Envelope.
  ```json
  {
    "success": true,
    "data": { ... },
    "error": null,
    "meta": { "page": 1, "limit": 10 }
  }
  ```

## 2. Golang Standard

- **Style**: Standard `gofmt`.
- **Error Handling**: Never underscore errors (`_`). Always handle or wrap errors.
  - _Bad_: `val, _ := strconv.Atoi(...)`
  - _Good_: `if err != nil { return fmt.Errorf("context: %w", err) }`
- **Project Layout**: Follow standard-go-project-layout (`cmd/`, `internal/`, `pkg/`).

## 3. Node.js / TypeScript Standard

- **Style**: ESLint + Prettier (Standard Config).
- **Architecture**: NestJS Modular Architecture.
- **Async**: Always use `async/await` instead of raw Promises/Callbacks.
- **Types**: No `any`. Define Interface/DTO for every payload.

## 4. Python Standard

- **Style**: PEP8 (enforced by Black/Flake8).
- **Type Hinting**: Mandatory for function arguments and return types.
  ```python
  def calculate_eta(queue_size: int) -> float:
  ```
- **Dependencies**: Manage via `poetry` or `pipenv` (deterministic lock file).

## 5. Git Convention

See `git-workflow.md` in root directory.
