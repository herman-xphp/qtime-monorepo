# DevOps & Infrastructure Tickets 🚀

## Phase 4: Infrastructure & Hardening

### [QTIME-034] API Gateway Implementation
**Type**: Chore | **Points**: 5
**Description**: Setup Traefik or Kong as a single entry point for all services.
**Acceptance Criteria**:
- [ ] Single port exposure (80/443).
- [ ] Path-based routing: `/queue/` -> Go, `/api/` -> Java, `/eta/` -> Python.

### [QTIME-035] Distributed Tracing (OpenTelemetry)
**Type**: Chore | **Points**: 8
**Description**: Integrate OTel into all polyglot services and setup Jaeger.
**Acceptance Criteria**:
- [ ] Correlation ID propagated across HTTP and Redpanda headers.

### [QTIME-036] Polyglot DB Migration Standard
**Type**: Chore | **Points**: 3
**Description**: Setup `golang-migrate` for Go and `alembic` for Python.
**Acceptance Criteria**:
- [ ] Migrations run automatically during container startup.

---

## Phase 5: Automation & Reliability

### [QTIME-037] CI/CD Pipeline (GitHub Actions)
**Type**: DevOps | **Points**: 5
**Description**: Automate Lint, Test, and Build for all services.
**Acceptance Criteria**:
- [ ] Workflows for each service folder.
- [ ] Docker Hub auto-push on green build.

### [QTIME-038] Secrets Management Implementation
**Type**: Security | **Points**: 5
**Description**: Move from plain Env Vars to SOPS or HashiCorp Vault.
**Acceptance Criteria**:
- [ ] Encrypted secrets at rest in the repository.

### [QTIME-039] E2E Integration Testing
**Type**: QA | **Points**: 8
**Description**: Setup cross-service tests using Testcontainers.
**Acceptance Criteria**:
- [ ] Complete flow (Booking -> Redpanda -> Notification) verified in isolated Docker env.
