# Industrial Logging Standards - Q-Time 🪵

## 1. Core Principles

To ensure observability in a polyglot microservices environment, all services must adhere to **Structured Logging** using **JSON format**. This allows for easy parsing by Log Aggregators (ELK, Grafana Loki, Datadog).

### A. Format Requirements

Every log entry MUST be a single-line JSON object containing these mandatory fields:

| Field       | Type    | Description                                      | Example                      |
| :---------- | :------ | :----------------------------------------------- | :--------------------------- |
| `timestamp` | ISO8601 | UTC timestamp of the event.                      | `"2026-01-30T20:05:00.000Z"` |
| `level`     | String  | Log level (lowercase or uppercase).              | `"info"`, `"error"`          |
| `service`   | String  | Name of the microservice.                        | `"queue-engine"`             |
| `trace_id`  | String  | Unique ID to correlate requests across services. | `"v1-a8b2c5..."`             |
| `message`   | String  | Human readable description of the event.         | `"Ticket created"`           |
| `context`   | Object  | Arbitrary key-value pairs relevant to the event. | `{"merchant_id": "M-1"}`     |

---

## 2. Log Levels

| Level     | Usage                                                                 | Production? |
| :-------- | :-------------------------------------------------------------------- | :---------- |
| **DEBUG** | Fine-grained info for local development/debugging.                    | No          |
| **INFO**  | Normal operational events (e.g. Service started, Ticket created).     | Yes         |
| **WARN**  | Unusual events that are not errors (e.g. Slow DB query, Retrying).    | Yes         |
| **ERROR** | Failed operations that require attention (e.g. DB connection failed). | Yes         |
| **FATAL** | Critical failure that causes service shutdown.                        | Yes         |

---

## 3. Polyglot implementation recommendations

### 🚀 Golang (Queue Engine)

Use **`uber-go/zap`** or the native **`log/slog`** (Go 1.21+).

```go
// slog example
logger.Info("ticket_created",
    slog.String("merchant_id", "M-123"),
    slog.String("trace_id", traceID),
)
```

### ☕ Java (Core Backend)

Use **Logback** with **SLF4J**. Configure `logback.xml` to use `LogstashEncoder`.

```java
log.info("Process payment for merchant: {}", merchantId, StructuredArguments.kv("trace_id", tid));
```

### 🟢 Node.js (Notification Worker)

Use **`pino`** or **`winston`**. Pino is preferred for performance.

```typescript
logger.info({ merchantId, traceId }, "Dispatching WhatsApp message");
```

### 🐍 Python (Intelligence Worker)

Use **`structlog`** or **`loguru`** with JSON formatting.

```python
logger.info("predicting_eta", merchant_id="M-1", trace_id=tid)
```

---

## 4. Security & Privacy (PII) 🛡️

**NEVER** log sensitive information:

- Passwords, Tokens, Secret Keys.
- Full Credit Card Numbers (Mask to `****-****-****-1234`).
- National ID (NIK) or Private Phone Numbers (Redact or Hash).

---

## 5. Traceability (Correlation ID)

1.  **Incoming Request**: API Gateway generates a `X-Trace-ID`.
2.  **Propagation**: Every internal gRPC/HTTP call or Messaging (Redpanda) MUST carry this `trace_id`.
3.  **Logs**: Every log entry related to that request MUST include the `trace_id`.

---

## 6. Anti-Patterns (What NOT to do) ❌

- **Avoid raw strings**: `log.info("User " + user + " logged in")` -> Hard to parse.
- **No Multi-line**: Stacks should be formatted into a single JSON field `stack_trace`.
- **Level Abuse**: Don't use ERROR for business validation (e.g., "Invalid OTP" is INFO/WARN).
