# Sprint 6: API Gateway & Gateway Security 🛡️

**Goal**: Centralize traffic and secure the ecosystem.

### 🎫 Tickets
- [ ] **QTIME-034**: API Gateway Implementation (Traefik/Kong).
- [ ] **QTIME-038**: Secrets Management (SOPS/Vault).
- [ ] **QTIME-009**: Security Hardening (IP Whitelisting/Rate Limiting at Gateway).

### ✅ Definition of Done (DoD)
- [ ] All external traffic goes through port 80/443.
- [ ] Rate limiting active for `/queue/take` endpoint.
- [ ] No plain-text secrets in repository.
