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

## 2. Commit Messages (Conventional Commits)

We enforce **Conventional Commits** to automate changelogs and versioning.

**Format**:
`type(scope): subject`

**Types**:

- `feat`: A new feature (Correlates with MINOR version).
- `fix`: A bug fix (Correlates with PATCH version).
- `docs`: Documentation only changes.
- `style`: Formatting, missing semi colons, etc (no code change).
- `refactor`: Code change that neither fixes a bug nor adds a feature.
- `test`: Adding missing tests or correcting existing tests.
- `chore`: Changes to build process or auxiliary tools.

**Examples**:

- ✅ `feat(queue): add redis atomic lua script`
- ✅ `fix(auth): handle expired jwt token gracefully`
- ✅ `docs(readme): update architecture diagram`
- ❌ `update code` (Too vague)
- ❌ `fixing bug` (No type/scope)

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
