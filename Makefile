.PHONY: help up down ps build logs logs-q logs-n logs-c logs-i restart-q restart-n restart-c restart-i test-load test-api db-shell redis-shell topic-list lint test clean

# Default target
help:
	@echo "🛠️  Q-Time Monorepo Helper"
	@echo "========================="
	@echo "Infrastructure & Lifecycle:"
	@echo "  make up          : Start all services (detached)"
	@echo "  make up-q        : Start Queue Engine only"
	@echo "  make up-n        : Start Notification Worker only"
	@echo "  make up-c        : Start Core Backend only"
	@echo "  make up-i        : Start Intelligence Worker only"
	@echo "  make down        : Stop and remove all containers"
	@echo "  make build       : Rebuild all Docker images"
	@echo "  make ps          : Show running containers"
	@echo "  make clean       : Prune docker system (use with caution)"
	@echo ""
	@echo "Logs & Monitoring:"
	@echo "  make logs        : Follow logs from all services"
	@echo "  make logs-q      : Logs - Queue Engine (Go)"
	@echo "  make logs-n      : Logs - Notification Worker (Node)"
	@echo "  make logs-c      : Logs - Core Backend (Java)"
	@echo "  make logs-i      : Logs - Intelligence Worker (Python)"
	@echo ""
	@echo "Service Management:"
	@echo "  make restart-q   : Restart Queue Engine"
	@echo "  make restart-n   : Restart Notification Worker"
	@echo "  make restart-c   : Restart Core Backend"
	@echo "  make restart-i   : Restart Intelligence Worker"
	@echo ""
	@echo "Database & Tools:"
	@echo "  make db-shell    : Access Postgres shell (psql)"
	@echo "  make redis-shell : Access Redis shell (redis-cli)"
	@echo "  make topic-list  : List Redpanda topics"
	@echo ""
	@echo "Testing & QA:"
	@echo "  make test-load   : Run K6 Load Test (1000 VUs)"
	@echo "  make test-api    : Run single cURL request (Take Ticket)"
	@echo "  make test        : Run Unit Tests for all services (Local)"
	@echo "  make lint        : Run Linters for all services (Local)"

# --- Infrastructure ---
up:
	docker compose up -d

up-q:
	docker compose up -d queue-engine

up-n:
	docker compose up -d notification-worker

up-c:
	docker compose up -d core-backend

up-i:
	docker compose up -d intelligence-worker

down:
	docker compose down

ps:
	docker compose ps

build:
	docker compose build

clean:
	docker system prune -f

# --- Logs ---
logs:
	docker compose logs -f

logs-q:
	docker compose logs -f queue-engine

logs-n:
	docker compose logs -f notification-worker

logs-c:
	docker compose logs -f core-backend

logs-i:
	docker compose logs -f intelligence-worker

# --- Restart ---
restart-q:
	docker compose restart queue-engine && docker compose logs -f queue-engine

restart-n:
	docker compose restart notification-worker && docker compose logs -f notification-worker

restart-c:
	docker compose restart core-backend && docker compose logs -f core-backend

restart-i:
	docker compose restart intelligence-worker && docker compose logs -f intelligence-worker

# --- Tools ---
db-shell:
	docker compose exec postgres psql -U user -d qtime_db

redis-shell:
	docker compose exec redis redis-cli

topic-list:
	docker compose exec redpanda rpk topic list

# --- Testing ---
test-load:
	k6 run tests/k6/load-test.js

test-api:
	@echo "🎟️ Taking a ticket for merchant-123..."
	curl -X POST http://localhost:3000/queue/take \
		-H "Content-Type: application/json" \
		-d '{"merchant_id": "merchant-123"}'
	@echo "\n"

test:
	@echo "🧪 Testing Queue Engine (Go)..."
	cd services/queue-engine && go test ./...
	@echo "🧪 Testing Notification Worker (Node)..."
	cd services/notification-worker && npm test
	@echo "🧪 Testing Core Backend (Java)..."
	cd services/core-backend && ./mvnw test
	@echo "🧪 Testing Intelligence Worker (Python)..."
	# cd services/intelligence-worker && pytest

lint:
	@echo "🔍 Linting Queue Engine (Go)..."
	# cd services/queue-engine && golangci-lint run
	@echo "🔍 Linting Notification Worker (Node)..."
	cd services/notification-worker && npm run lint
