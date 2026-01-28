# Master Ticket Backlog 🎫

Ini adalah daftar tiket detail untuk seluruh fase project.

## 🔁 Sprint 2: Notification & Async

**[QTIME-010] Setup Redpanda Producer in Go**

- **Type**: Feature | **Points**: 3
- **Tech**: Sarama / Redpanda Client
- **Criteria**:
  - Helper function `PublishEvent(topic, payload)` created in Go.
  - Event `TICKET_CREATED` published when user takes a queue.

**[QTIME-011] Notification Service (Node.js)**

- **Type**: Feature | **Points**: 5
- **Tech**: NestJS + Kafka Microservice Strategy
- **Criteria**:
  - Service consume topic `TICKET_CREATED`.
  - Log message "Sending WA to User..." to console.

**[QTIME-012] WebSocket Hub Implementation**

- **Type**: Feature | **Points**: 8
- **Tech**: Go Fiber Websocket / Gorilla
- **Criteria**:
  - Client can connect to `ws://host/ws/queue/{merchant_id}`.
  - Server broadcasts update to room when Admin calls a number.

## 🧠 Sprint 3: Intelligence (Python)

**[QTIME-020] Setup Python ETL Worker**

- **Type**: Support | **Points**: 3
- **Tech**: FastAPI + Pandas
- **Criteria**:
  - Job runs every 10 minutes.
  - Reads `completed_queues` table from Postgres.

**[QTIME-021] Implement Moving Average Algo**

- **Type**: Algorithm | **Points**: 5
- **Logic**: `AVG(end_time - start_time)` for last N customers.
- **Criteria**:
  - Unit test with sample dataset.
  - API returns `estimated_wait_seconds`.

## 🛡️ Sprint 4: Production Ready

**[QTIME-030] Docker Multi-stage Build**

- **Type**: DevOps | **Points**: 3
- **Criteria**:
  - Go binary image size < 20MB (Distroless/Alpine).
  - No source code in final image.

**[QTIME-031] Load Test with k6**

- **Type**: QA | **Points**: 5
- **Scenario**: 500 users taking number simultaneously.
- **Success**: 0 Errors, p95 latency < 500ms.

**[QTIME-032] CI/CD Pipeline Config**

- **Type**: DevOps | **Points**: 5
- **Criteria**:
  - GitHub Action `on: push to main`.
  - Runs `go test`, `npm test`.
  - Builds and pushes Docker image.
