# Technical Architecture - Q-Time 🏗️

## 1. High Level Design

Q-Time utilizes a **Polyglot Configuration** to leverage the best tool for each specific problem domain.

```mermaid
graph TD
    Client[User Device] -->|HTTPS| PublicApp[Booking PWA (Next.js)]
    Staff[Clinic Staff] -->|HTTPS| AdminApp[Admin Dash (Svelte)]

    PublicApp -->|HTTPS/WSS| Gateway[API Gateway]
    AdminApp -->|HTTPS/WSS| Gateway

    subgraph "Core Services"
        Gateway -->|HTTP/REST| QueueSvc[Queue Engine (Go)]
        Gateway -->|HTTP/REST| CoreSvc[Core Backend (Java)]
    end

    subgraph "Async Workers"
        QueueSvc -->|Pub/Sub| Redpanda{Redpanda/Kafka}
        Redpanda -->|Consume| NotifSvc[Notification (Node.js)]
        Redpanda -->|Consume| EtaSvc[ETA Predictor (Python)]
    end

    subgraph "Data Layer"
        QueueSvc -.-> Redis[(Redis - Hot Data)]
        AdminSvc -.-> Postgres[(Postgres - Cold Data)]
    end
```

## 2. Service Responsibilities

| Service           | Language             | Port | Database      | Responsibility                                                                          |
| :---------------- | :------------------- | :--- | :------------ | :-------------------------------------------------------------------------------------- |
| **queue-engine**  | **Go (Fiber)**       | 3000 | Redis + PG    | Handle High-Concurrrency Queue operations (Take, Call, Skip). Atomic Locking via Redis. |
| **core-backend**  | **Java (Spring)**    | 8080 | Postgres      | **Core Business Logic**: Tenant Mgmt, Billing, User Auth, Reporting.                    |
| **notif-service** | **Node.js (Nest)**   | 3001 | Redis (Queue) | Integration with WhatsApp/Email Providers. managing async jobs.                         |
| **eta-service**   | **Python (FastAPI)** | 3002 | PG (Read)     | Read historical data, calculate Moving Average for accurate ETA.                        |
| **admin-web**     | **SvelteKit**        | 5173 | -             | Admin Dashboard Frontend.                                                               |
| **public-mobile** | **Next.js**          | 3005 | -             | Public Booking PWA.                                                                     |

## 3. Communication Protocols

- **External (Client -> Server)**: REST API (JSON) + WebSocket (Socket.io) for realtime updates.
- **Internal (Service -> Service)**:
  - **Synchronous**: gRPC (if strict consistency needed).
  - **Asynchronous**: Redpanda (Kafka Protocol) for decoupling booking flow from notification.

## 4. Scalability Strategy

- **Horizontal Scaling**: All services are stateless. Can run multiple replicas behind Load Balancer.
- **Database**:
  - Postgres: Read Replicas for Analytics.
  - Redis: Cluster Mode for High Availability.
