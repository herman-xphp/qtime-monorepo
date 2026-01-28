# Git Workflow & Standards 🌲

## 1. Branching Strategy (Git Flow)

We use a standard **Git Flow** centered around `develop`.

| Branch    | Source    | Description                                       |
| :-------- | :-------- | :------------------------------------------------ |
| `main`    | -         | **Production Ready**. Only via PR from `develop`. |
| `develop` | `main`    | **Integration Branch**. The ongoing work stream.  |
| `feat/*`  | `develop` | New features (e.g., `feat/queue-logic`).          |
| `fix/*`   | `develop` | Bug fixes (e.g., `fix/memory-leak`).              |
| `chore/*` | `develop` | Maintenance (e.g., `chore/setup-infra`).          |

---

## 2. Commit Messages (Strict Format) 📜

We enforce a **3-part structure** for all commits to ensure clarity and traceability.

**Structure**:

```text
<type>(<scope>): <subject>

- <detailed bullet point 1>
- <detailed bullet point 2>

Ticket: <TICKET-ID>
```

**Types**: `feat`, `fix`, `docs`, `chore`, `test`, `refactor`.

**Example**:

```text
feat(service): implement CustomerService with soft-delete

- Implement create() with duplicate validation (KTP, Email)
- Implement update() for partial update
- Implement soft-delete in delete() using 'CLOSED' status

Ticket: SYR-005
```

---

## 3. Pull Request (PR) Lifecycle

Since we are simulating a team environment, every feature must go through a **PR Process**:

1.  **Create Branch**:
    ```bash
    git checkout -b feat/my-feature develop
    ```
2.  **Work & Commit**:
    ```bash
    git commit -m "feat(scope): my work"
    ```
3.  **Open PR (Simulation)**:
    - Review files.
    - Run tests.
4.  **Merge (Simulation)**:
    Perusahaan mewajibkan **Merge Commit** (`--no-ff`) agar history fitur terlihat jelas.
    ```bash
    git checkout develop
    git merge --no-ff feat/my-feature -m "merge: PR #<ID> <Title>"
    git branch -d feat/my-feature
    ```
