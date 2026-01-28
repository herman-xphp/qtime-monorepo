# API Specification (Complete) 📡

This document complements `openapi.yaml`. While the YAML file is for tools, this document is for **Humans**.

## 🔄 Lifecycle of a Queue Ticket

1.  **Created**: User calls `POST /queue/take`. Status: `WAITING`.
2.  **Called**: Staff calls `POST /queue/call`. Status: `CALLED`.
3.  **Finished**: Staff marks as `POST /queue/complete` or `POST /queue/skip`. Status: `COMPLETED` / `SKIPPED`.

## 1. Queue Engine (Go) - Port 3000

Start command: `cd services/queue-engine && go run main.go`

| Mtd    | Endpoint                | Auth   | Desc                                    |
| :----- | :---------------------- | :----- | :-------------------------------------- |
| `POST` | `/queue/take`           | User   | **Atomic**. User takes a number.        |
| `POST` | `/queue/call`           | Staff  | **Atomic**. Staff calls next number.    |
| `PUT`  | `/queue/{id}/status`    | Staff  | Mark as Completed, Skipped, or No-Show. |
| `GET`  | `/queue/{mid}/realtime` | Public | Get dashboard state for TV Display.     |
| `GET`  | `/queue/my-ticket`      | User   | Get current user's active ticket.       |

### Business Rules (Go)

- **Duplicate Check**: A user cannot take a second ticket if they have a `WAITING` or `CALLED` ticket in the same Merchant efficiently check via Redis Set.
- **Geo-fencing**: (Optional) User can only take ticket if within 5km radius (Validated by coordinates in body).

---

## 2. Core Backend (Java) - Port 8080

Start command: `cd services/core-backend && ./mvnw spring-boot:run`

| Mtd    | Endpoint             | Auth   | Desc                               |
| :----- | :------------------- | :----- | :--------------------------------- |
| `POST` | `/api/auth/register` | Public | Register new Tenant (Clinic/Bank). |
| `POST` | `/api/auth/login`    | Public | Login for Admin/Staff/User.        |
| `GET`  | `/api/merchants/me`  | Admin  | Get own merchant profile.          |
| `PUT`  | `/api/merchants/me`  | Admin  | Update config (Hours, Logo).       |
| `POST` | `/api/counters`      | Admin  | Add physical counter (Loket).      |
| `POST` | `/api/staff`         | Admin  | Invite staff member (Role based).  |
| `GET`  | `/api/reports/daily` | Admin  | PDF Download of daily stats.       |

### Business Rules (Java)

- **Tenant Isolation**: All queries must enforce `WHERE merchant_id = ?` based on JWT claims.
- **Audit Log**: All WRITE operations must be logged to `audit_logs` table (MongoDB).

---

## 3. Intelligence (Python) - Port 3002

Start command: `cd services/eta-service && uvicorn main:app --reload`

| Mtd   | Endpoint     | Desc                                   |
| :---- | :----------- | :------------------------------------- |
| `GET` | `/eta/{mid}` | Returns estimated wait time (in mins). |

---

## 4. Notifications (Node.js) - Worker

- No Public HTTP Endpoints.
- Listens to Redpanda Topic: `queue-events`.
- Triggers:
  - `TICKET_CREATED` -> WA: "Tiket Anda A-001."
  - `TICKET_CALLED` -> WA: "Silakan meuju Loket 1."
  - `TICKET_REMINDER` -> WA: "Giliran Anda 5 menit lagi."

---

## 5. Deployment / Gateway Configuration

In production, all these services are hidden behind **Nginx/Traefik**.

- `api.qtime.id/queue/*` -> `queue-engine:3000`
- `api.qtime.id/api/*` -> `core-backend:8080`
- `api.qtime.id/eta/*` -> `eta-service:3002`
