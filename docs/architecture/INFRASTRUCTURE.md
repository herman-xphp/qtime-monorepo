# Infrastructure & DevOps ☁️

## 1. Local Development (Docker Compose)

See root `docker-compose.yml`.

- **Postgres 15**: Persistence.
- **Redis 7**: Caching & Locking.
- **Redpanda**: Event Streaming.
- **Mailhog** (Optional): SMTP Testing tool.

## 2. CI/CD Pipeline Strategy

We use GitHub Actions / GitLab CI.

### Stages:

1.  **Commit**: Developer pushes code.
2.  **Lint & Test**:
    - Go: `golangci-lint`, `go test ./...`
    - Java: `mvn verify`, `checkstyle`
    - Node: `eslint`, `npm run test`
    - Python: `black`, `pytest`
3.  **Build**: Create Docker Images (`qtime/queue-engine:sha-xyz`). Scanner security (Trivy).
4.  **Deploy (Staging)**: Update Kubernetes Manifest (ArgoCD sets auto-sync).

## 3. Production Deployment (Kubernetes)

- **Ingress Controller**: Nginx / Traefik.
- **Cert Manager**: Auto SSL (Let's Encrypt).
- **Horizontal Pod Autoscaler (HPA)**: Auto-scale replicas based on CPU/Memory usage (Crucial for Monday morning spikes).

## 4. Observability

- **Logs**: Stdout json format -> Fluentd -> Elastic/Loki.
- **Metrics**: Prometheus scrapes `/metrics` endpoint.
- **Tracing**: OpenTelemetry agent injected in sidecar.
