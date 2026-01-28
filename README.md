# Q-Time: Universal Queue Management System ⏳

![Build Status](https://img.shields.io/badge/build-passing-brightgreen)
![Quality Gate](https://img.shields.io/badge/quality-A-brightgreen)
![Coverage](https://img.shields.io/badge/coverage-80%25-green)
![License](https://img.shields.io/badge/license-MIT-blue)

**Q-Time** is an enterprise-grade SaaS solution designed to eliminate the uncertainty of physical waiting lines. It orchestrates real-time queuing for clinics, banks, and service centers using a robust **Polyglot Microservices Architecture**.

> **Simulation Context**: This project simulates a Series-A startup environment ("Industrial Standard"), enforcing strict separation of concerns, CI/CD gates, and Production-Ready code standards from Day 1.

---

## 📑 Table of Contents

- [Architecture & Tech Stack](#-architecture--tech-stack)
- [Monorepo Structure](#-monorepo-structure)
- [Getting Started](#-getting-started)
- [Port Mapping](#-port-mapping-local-dev)
- [Development Workflow](#-development-workflow)
- [Documentation Library](#-documentation-library)

---

## 🏗️ Architecture & Tech Stack

We utilize a **Best-for-Job Strategy**, leveraging distinct languages for specific domains.

```mermaid
graph LR
    Client -->|HTTPS| PublicApp["PWA (Next.js)"]
    Staff -->|HTTPS| AdminApp["Admin (SvelteKit)"]

    PublicApp --> Gateway
    AdminApp --> Gateway

    Gateway["API Gateway"] --> QueueSvc("Go: Queue Engine")
    Gateway --> CoreSvc("Java: Core Backend")

    QueueSvc -->|Events| Redpanda
    Redpanda --> NotifSvc("Node: Notifications")
    Redpanda --> ML("Python: Intelligence")
```

| Domain               | Service         | Stack                           | Key Responsibility                                     |
| :------------------- | :-------------- | :------------------------------ | :----------------------------------------------------- |
| **Edge**             | **Apps**        | **Next.js 16** / **SvelteKit**  | Public Booking PWA & Admin Dashboard.                  |
| **High Performance** | `queue-engine`  | **Go 1.25** + **Fiber**         | Atomic Ticket Generation (Redis Lua), WebSocket Hub.   |
| **Business Core**    | `core-backend`  | **Java 25** + **Spring Boot 4** | Billing, Tenants, Reporting. Uses **Virtual Threads**. |
| **IO Bound**         | `notif-service` | **Node.js 24** + **NestJS**     | WhatsApp/Email Dispatcher.                             |
| **Data/AI**          | `eta-service`   | **Python 3.14** + **FastAPI**   | ETA Prediction (Moving Average).                       |
| **Infrastructure**   | -               | Docker, Redpanda, Redis, PG     | Event Streaming & Persistence.                         |

👉 _Read the full Rationale in [Tech Stack Decisions (ADR)](docs/architecture/TECH_STACK_DECISIONS.md)._

---

## 📦 Monorepo Structure

```text
qtime-monorepo/
├── apps/                       # 🖥️ Frontend Applications
│   ├── admin-web/              # SvelteKit (Admin Dashboard)
│   └── public-mobile/          # Next.js (Patient PWA)
├── services/                   # 🚀 Backend Microservices
│   ├── queue-engine/           # Go (High Perf)
│   ├── core-backend/           # Java (Business Logic)
│   ├── notif-service/          # Node.js (Async)
│   └── eta-service/            # Python (Data)
├── docs/                       # 📚 Documentation
│   ├── architecture/           # Arch, Infra, ADRs
│   ├── product/                # PRD & User Flows
│   └── project-management/     # Sprints & Tickets
├── docker-compose.yml          # 🐳 Local Infrastructure
└── Makefile                    # 🛠️ Shortcut Commands
```

---

## 🚀 Getting Started

### 1. Prerequisites

- **Docker & Docker Compose** (Required)
- **Go 1.25+**, **Java 25 LTS**, **Node 24 LTS** (Recommended for local dev)

### 2. Quick Start

Clone and start the infrastructure:

```bash
git clone https://github.com/your-org/qtime-monorepo.git
cd qtime-monorepo

# Start DB, Cache, and Message Broker
docker-compose up -d
```

### 3. Verify Health

Access the services (once running):

- **Queue API**: `http://localhost:3000/health`
- **Core API**: `http://localhost:8080/actuator/health`
- **Redpanda Console**: `http://localhost:8081`

---

## 🔌 Port Mapping (Local Dev)

| Service           | Port   | Debug  | Database      | Validated URL           |
| :---------------- | :----- | :----- | :------------ | :---------------------- |
| **Queue Engine**  | `3000` | `2345` | Redis `8`     | `POST /queue/take`      |
| **Core Backend**  | `8080` | `5005` | Postgres `18` | `POST /auth/login`      |
| **Notifications** | `3001` | `9229` | -             | -                       |
| **Intelligence**  | `3002` | -      | -             | `GET /eta/{id}`         |
| **Admin Web**     | `5173` | -      | -             | `http://localhost:5173` |
| **Public PWA**    | `3005` | -      | -             | `http://localhost:3005` |

---

## 📚 Documentation Library

Please read these before writing code:

### 🧠 Planning & Architecture

- **[Vision & PRD](docs/product/PRD.md)**: Product Requirements.
- **[Architecture](docs/architecture/ARCHITECTURE.md)**: System Design.
- **[Frontend Architecture](docs/architecture/FRONTEND_ARCHITECTURE.md)**: UI Strategy.
- **[API Specification](docs/api/API_SPEC.md)**: Clean URL Contracts (No `/v1`).

### ⚙️ Standards

- **[Coding Standards](docs/standards/CODING_STANDARDS.md)**: Go, Java, TS, Python rules.
- **[Git Workflow & AI Protocol](docs/standards/GIT_WORKFLOW.md)**: How to commit & collaborate.

### 📅 Project Management

- **[Roadmap](docs/project-management/ROADMAP_TO_PRODUCTION.md)**: Timeline.
- **Active Sprints**:
  - [Backend Plan](docs/project-management/sprints/backend/SPRINT_1.md)
  - [Frontend Plan](docs/project-management/sprints/frontend/SPRINT_1.md)

---

## 🤝 Contribution Workflow

1.  Pick a ticket from:
    - **[Backend Tickets](docs/project-management/tickets/BACKEND_TICKETS.md)**
    - **[Frontend Tickets](docs/project-management/tickets/FRONTEND_TICKETS.md)**
2.  Create branch: `git checkout -b feat/queue-logic`
3.  Commit (Conventional): `git commit -m "feat(queue): implement atomic redis lock"`
4.  Push & PR.

---

_Built with ❤️ by the Q-Time Engineering Team_
