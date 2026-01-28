# Backend Engineering Tickets 🏗️

## Phase 1: Foundation (Go & Java)

### [QTIME-001] Setup Monorepo & Infrastructure

**Type**: Chore | **Story Points**: 3
**Description**: Initialize project structure and docker-compose for local development.
**Acceptance Criteria**:

- [ ] `docker-compose up` starts Postgres (5432), Redis (6379), Redpanda (9092).
- [ ] `services/` folder contains subfolders for `queue-engine`, `notif-service`, `eta-service`.
- [ ] `.gitignore` is properly configured.

### [QTIME-002] Queue Engine Skeleton (Golang)

**Type**: Feature | **Story Points**: 5
**Description**: Create the main Golang service using Fiber framework.
**Tech Notes**:

- Use `github.com/gofiber/fiber/v2`
- Use Clean Architecture: `internal/handler`, `internal/service`, `internal/repository`.
  **Acceptance Criteria**:
- [ ] Application runs on port 3000.
- [ ] `GET /health` returns "OK".
- [ ] Graceful shutdown implemented.

### [QTIME-003] API: Take Ticket Number (Redis Atomic)

**Type**: Feature | **Story Points**: 8
**Description**: Endpoint for user to take a queue number using Redis Lua Script.
**Tech Notes**:

- Use `EVAL` for atomic increment.
  **Acceptance Criteria**:
- [ ] `POST /queue/take` returns unique ticket number.
- [ ] No duplicate numbers under load.

### [QTIME-004] Spring Boot Scaffold & Flyway

**Type**: Chore | **Story Points**: 5
**Description**: Initialize Java Core Backend with Flyway for DB Migrations.
**Tech Notes**:

- **Spring Boot 4.0.2**, Java 21 (LTS).
- Enable Virtual Threads (`spring.threads.virtual.enabled=true`).
- Disable `spring.jpa.hibernate.ddl-auto`.
  **Acceptance Criteria**:
- [ ] `mvn clean install` success.
- [ ] `V1__init_schema.sql` runs on startup to create `tenants` table.

## Phase 1.5: Production Hardening

### [QTIME-009] Spring Security & JWT

**Type**: Feature | **Story Points**: 8
**Description**: Implement stateless JWT Authentication.
**Acceptance Criteria**:

- [ ] `JwtAuthenticationFilter` validates Bearer token.
- [ ] `@PreAuthorize` works on controllers.

### [QTIME-013] Observability Setup

**Type**: Chore | **Story Points**: 5
**Description**: Add Actuator and Prometheus metrics.
**Acceptance Criteria**:

- [ ] `GET /actuator/prometheus` returns metrics.

## Phase 2: Connectivity

### [QTIME-010] Go: Redpanda Producer

**Type**: Feature | **Story Points**: 5
**Description**: Publish `TICKET_CREATED` events.

### [QTIME-011] Node: Redpanda Consumer

**Type**: Feature | **Story Points**: 5
**Description**: Consume events and log them.

## Phase 3: Intelligence

### [QTIME-020] Python Service Scaffold

**Type**: Chore | **Story Points**: 3
**Description**: Setup FastAPI project.

### [QTIME-030] Docker Optimization

**Type**: DevOps | **Story Points**: 5
**Description**: Multistage builds for all services.
