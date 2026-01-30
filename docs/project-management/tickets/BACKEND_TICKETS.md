# Backend Engineering Tickets 🏗️

## Phase 1: Foundation & Core Services

### [QTIME-001] Setup Monorepo & Infrastructure [DONE]
**Type**: Chore | **Points**: 3
**Description**: Initialize project structure and docker-compose.

### [QTIME-002] Queue Engine Skeleton (Golang) [DONE]
**Type**: Feature | **Points**: 5
**Description**: Basic Go Fiber service.

### [QTIME-003] API: Take Ticket Number (Redis Atomic) [DONE]
**Type**: Feature | **Points**: 8
**Description**: Atomic ticket generation using Lua script.

### [QTIME-004] Spring Boot Scaffold & Flyway [DONE]
**Type**: Chore | **Points**: 5
**Description**: Base Java service with DB migrations.

---

## Phase 2: Security & Business Logic (Java)

### [QTIME-009] Spring Security & JWT Implementation
**Type**: Feature | **Points**: 8
**Description**: Implement stateless authentication with JWT for multiple roles.

### [QTIME-010] Tenant (Merchant) Management API
**Type**: Feature | **Points**: 8
**Description**: CRUD for Merchants, Counters, and Staff. Enforce Tenant Isolation.

### [QTIME-011] Staff Interaction: Call, Skip, Complete
**Type**: Feature | **Points**: 8
**Description**: API for staff to manage queue lifecycle.

### [QTIME-012] Reporting Engine: Daily Stats
**Type**: Feature | **Points**: 5
**Description**: Aggregate queue data for daily business reports.

---

## Phase 3: Event-Driven & Intelligence

### [QTIME-013] Redpanda Event Integration
**Type**: Feature | **Points**: 5
**Description**: Publish queue events (Created, Called) to Redpanda.

### [QTIME-014] Notification Worker (Node.js)
**Type**: Feature | **Points**: 5
**Description**: Consume events and send WhatsApp/Email stubs.

### [QTIME-015] ETA Service (Python)
**Type**: Feature | **Points**: 8
**Description**: Moving average calculation for dynamic ETAs.

---

## Phase 4: Production Hardening

### [QTIME-030] Industrial Logging & Observability
**Type**: Chore | **Points**: 5
**Description**: Implement JSON logging, Actuator, and Prometheus metrics.

### [QTIME-031] Container Optimization & K8s Manifests
**Type**: DevOps | **Points**: 8
**Description**: Multistage Dockerfiles and Helm charts.

### [QTIME-032] Load Testing & Performance Tuning
**Type**: QA | **Points**: 13
**Description**: k6 load testing and JVM/Go tuning.
