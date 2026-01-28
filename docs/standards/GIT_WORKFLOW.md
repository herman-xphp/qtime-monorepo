# Git Workflow Standard & Conventions 🛡️

Dokumen ini adalah **Single Source of Truth** untuk workflow Git di project Q-Time. Semua kontributor (Manusia & AI) wajib mengikuti standar ini untuk menjaga kualitas history code.

---

## 1. Branching Strategy

Kita menggunakan **Feature Branch Workflow** (adaptasi dari GitHub Flow).

- **`main`**: Protected branch. Kode di sini harus _deployable_ ke Production. Dilarang commit langsung ke main (kecuali setup awal).
- **`feat/name`**: Untuk fitur baru.
- **`fix/name`**: Untuk perbaikan bug.
- **`chore/name`**: Untuk perubahan non-code (config, docker, docs).
- **`refactor/name`**: Untuk merapikan code tanpa mengubah fitur.

**Naming Convention:**
Format: `type/context-description`

- ✅ `feat/queue-api-endpoint`
- ✅ `fix/redis-connection-timeout`
- ✅ `chore/update-docker-compose`
- ❌ `feat/login` (Terlalu umum)
- ❌ `budi/fix-bug` (Jangan pakai nama orang)

---

## 2. Commit Message Convention

Wajib menggunakan **Conventional Commits** agar changelog bisa digenerate otomatis.

**Format**: `type(scope): description`

- **Types**:
  - `feat`: Fitur baru untuk user.
  - `fix`: Perbaikan bug user.
  - `chore`: Update build task, package manager, config (no production code change).
  - `docs`: Perubahan dokumentasi.
  - `style`: Formatting, missing semi colons, etc.
  - `refactor`: Code change that neither fixes a bug nor adds a feature.
  - `test`: Adding missing tests.

- **Examples**:
  - ✅ `feat(queue): add POST /take endpoint`
  - ✅ `fix(auth): handle jwt token expiration`
  - ✅ `chore(infra): add redpanda to docker-compose`
  - ❌ `update code` (Tidak jelas)
  - ❌ `fix bug` (Bug yang mana?)

---

## 3. The Workflow (Step-by-Step)

### A. Mulai Pekerjaan

1.  Selalu mulai dari main yang terbaru:
    ```bash
    git checkout main
    git pull origin main
    ```
2.  Buat branch baru:
    ```bash
    git checkout -b feat/setup-queue-engine
    ```

### B. Selama Coding

3.  Lakukan perubahan atomic. Jangan gabung fitur A dan B dalam satu commit.
4.  Commit dengan pesan yang deskriptif:
    ```bash
    git commit -m "feat(queue): initialize fiber app and redis connection"
    ```

### C. Selesai Coding

5.  Push ke remote:
    ```bash
    git push origin feat/setup-queue-engine
    ```
6.  Buat **Pull Request (PR)** di GitHub/GitLab.
7.  **Self-Review**: Pastikan CI pass dan tidak ada conflict.
8.  Minta review dari rekan tim (atau AI Architect).
9.  **Squash & Merge** ke `main`.

---

## 4. Emergency (Hotfix)

Jika ada bug kritis di Production:

1.  Branch dari `main`: `git checkout -b hotfix/payment-gateway-down`
2.  Fix bug & Commit: `fix(payment): resolve timeout issue`
3.  Merge segera ke `main`.

---

## 5. AI Agent Protocol (Manual Mode) 🤖

**IMPORTANT**: Project ini berjalan dalam **Manual Mode**.

- **No Auto-Execution**: AI Agent **DILARANG** membuat file, folder, atau menjalankan perintah Git/Terminal secara otomatis.
- **Proposal First**: AI harus mengajukan perubahan (isi file atau command) terlebih dahulu.
- **Explicit Consent**: Eksekusi hanya boleh dilakukan setelah User memberikan instruksi jelas (misal: "buatkan filenya", "jalankan commandnya").
