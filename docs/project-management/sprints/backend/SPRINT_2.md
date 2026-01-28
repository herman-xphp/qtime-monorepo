# Backend Sprint 2: Connectivity & Events 🔗

**Goal**: Enable Asynchronous Communication.
**Duration**: 2 Weeks

## 🎯 Objectives

1.  **Queue Engine (Go)**:
    - Publish `TICKET_CREATED` events to Redpanda.
2.  **Notification (Node)**:
    - Setup NestJS Service.
    - Consume events from Redpanda.
    - Mock WhatsApp Sender.

## 🎫 Key Tickets

- **[QTIME-010]** Go: Redpanda Producer.
- **[QTIME-011]** Node: Redpanda Consumer.
