# Sprint 4: Production Hardening 🛡️

**Sprint Goal**: System is stable under high load (1000 RPS) and securely deployed to Kubernetes.
**Duration**: 2 Weeks (Mar 12 - Mar 25)
**Status**: Draft

---

## 📅 Timeline

- **Week 1**: Container Optimization & CI/CD Pipeline Setup.
- **Week 2**: Load Testing (k6) & Security Audit (Pentest).

## 👥 Team Capacity

- **DevOps**: 100%
- **All Devs**: 30% (Bug Fixing & Optimization based on Load Test)

## ✅ Definition of Done (DoD)

1.  Docker Image Size < 50MB (Go), < 200MB (Node/Python).
2.  k6 Test Result: p95 Latency < 200ms at 1000 VU.
3.  Security Scan (Trivy) returns 0 Critical/High Vulnerabilities.
4.  Deployed to Staging Environment.

## 🎫 Key Tickets

- **QTIME-030**: Docker Multistage Build Optimization.
- **QTIME-031**: Load Test Scenario (k6).
- **QTIME-032**: GitHub Actions CI/CD.
- **QTIME-033**: Kubernetes Helm Charts.
