# Roadmap to Production: Q-Time 🚀

Ini adalah High-Level Plan untuk mencapai "Production Ready" dalam 4 Sprint (Estimasi 8 Minggu).

## 🏃‍♂️ Sprint 1: Foundation & Core Engine

- **Theme**: "The Heartbeat"
- **Goal**: Sistem bisa menerima antrian (Take) dan memanggil antrian (Call) dengan data yang konsisten.
- **Key Deliverables**:
  - Monorepo Skeleton & Infra (Redis/PG/Redpanda).
  - `queue-engine` (Go) running with Lua Script locking.
  - `core-backend` (Java Spring Boot) for Admin/Tenant Management.
  - `apps/admin-web` (SvelteKit) Basic Layout.

## 🔔 Sprint 2: Connectivity & Notifications

- **Theme**: "Connecting People"
- **Goal**: User mendapatkan update status tanpa harus refresh halaman (Real-time & Async).
- **Key Deliverables**:
  - Inter-service communication (Go -> Redpanda -> Node).
  * `apps/public-mobile` (Next.js) Booking Flow.
  * `notif-service` (Node.js) consuming events.
  * Simulasi kirim WA/Email.
  * Frontend Dashboard awal (Svelte) untuk Admin.

- **Key Deliverables**:
  - `eta-service` (Python) calculating moving averages.
  - API Endpoint `GET /queue/{id}/eta`.
  - Admin Analytics (Grafik kepadatan harian).

## 🛡️ Sprint 4: Production Hardening

- **Theme**: "Bulletproof"
- **Goal**: Sistem siap dipukuli traffic tinggi dan aman dari serangan.
- **Key Deliverables**:
  - **Load Testing**: k6 script menembak 1000 RPS.
  - **Security**: Rate Limiting (Traefik), JWT Rotation.
  - **CI/CD**: GitHub Actions pipeline (Build, Test, Push to Registry).
  - **Deployment**: Kubernetes Manifests (Deployment, Service, Ingress).
