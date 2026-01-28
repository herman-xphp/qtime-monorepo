# Technology Stack & Architecture Decisions (ADR) 🧠

This document records the **strategic rationale** behind our technology choices.
In an enterprise environment, we don't pick tools because they are "hype", but because they solve specific constraints.

---

## 1. Core Language Strategy: "Polyglot"

**Decision**: Use distinct languages for Queue (Go), Business (Java), and Workers (Node/Python).
**Status**: Accepted.

### Why not Monolith (e.g., All Java)?

- **Scale**: The "Queue Engine" needs to handle 10k+ concurrent WebSocket connections during flash sales. Java threads are heavy (1MB stack). Go goroutines are light (2KB).
- **Maintenance**: Splitting domains prevents "Spaghetti Code".

### Why not All Go?

- **Complexity**: Complex business logic (Billing, Tax, Audit, Role Management) is verbose in Go. Java Spring Boot provides mature standards (JPA, Security) that speed up enterprise development.

---

## 2. Queue Engine: Golang + Fiber

**Context**: High Write Throughput, Low Latency, Websocket Hub.

- **Choice**: **Go 1.25 (Golang)**
- **Framework**: **Fiber v3** (over Gin/Echo).
- **Reasoning**:
  - Fiber is built on `fasthttp`, offering zero memory allocation path for hot routers.
  - Simple concurrency model (Channels) for broadcasting WebSocket events.

---

### 3. Core Backend: Java Spring Boot 4

**Context**: Complex Domain Model, Transaction Management, Reporting.

- **Choice**: **Java 21 (LTS)**
- **Framework**: **Spring Boot 4.0.2** (Latest Stable).
- **Reasoning**:
  - **Performance**: Spring Boot 4 brings optimizations for Virtual Threads out of the box.
  - **Future Proof**: Aligned with Jakarta EE 11 standard.
  - **Robustness**: Compile-time safety for complex billing logic.

---

## 4. Message Broker: Redpanda

**Context**: Event Sourcing between Queue and Notification services.

- **Choice**: **Redpanda**
- **Alternatives**: Apache Kafka, RabbitMQ.
- **Reasoning**:
  - **Performance**: C++ based, 10x lower tail latency than JVM-based Kafka.
  - **Simplicity**: Single binary (No Zookeeper, No JVM tuning required).
  - **Compatibility**: 100% Kafka API compatible (we can use standard Kafka Clients).

---

## 5. Database Strategy

### Primary: PostgreSQL 18

- **Usage**: Tenant Data, Users, Transactions (ACID).
- **Reasoning**: Reliability. JSONB support allows hybrid schema for flexible merchant config.

### Caching / Lock: Redis 8

- **Usage**: Distributed Locking (`SETNX`), Session Store, Hot Queue Counter.
- **Critical**: We use **Lua Scripts** in Redis to ensure atomic ticket generation. This prevents "Double Booking" at the database level.

---

## 6. Frontend Strategy: Monorepo "Apps"

**Question**: Where does the Frontend live? Separate Repo or Here?
**Decision**: Inside this Monorepo, under `apps/` directory.

- **Structure**:
  - `apps/admin-dashboard`: SvelteKit (Web for Clinic Staff).
  - `apps/booking-pwa`: React/Next.js (Mobile Web for Patients).
  - `services/notification-worker`: Node.js (Async Notifications).
  - `services/intelligence-worker`: Python (AI/ETA Prediction).
  - `services/queue-engine`: Go (Core Logic).
  - `services/core-backend`: Java (Admin/API).
- **Reasoning**:
  - **Atomic Commits**: If we change the API in `queue-engine`, we can fix the Frontend code in the _same_ PR. No "Version Mismatch" nightmare.
  - **Shared Contracts**: Frontend can directly import `openapi.yaml` or TypeScript types generated from the backend.

---

_Verified by Technical Architect_
