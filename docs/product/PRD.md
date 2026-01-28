# Product Requirements Document (PRD) - Q-Time ⏳

## 1. Executive Summary

Q-Time is a **Universal Queue Management System (SaaS)** aimed at eliminating the "uncertainty of waiting" for service-based businesses (clinics, salons, banks) and their customers.

**Value Proposition**:

- **For Users**: "Don't wait in line. Wait at home/coffee shop. Arrive just in time."
- **For Business**: Reduce crowded waiting rooms, reduce "no-shows" via deposits/penalties, and gain operational insights.

---

## 2. User Interfaces

1.  **Mobile Web (PWA) - For Patients**:
    - **Search**: Find nearest clinic.
    - **Take Number**: Select service -> Get Ticket.
    - **Live Status**: "5 people ahead of you".
    - **My Tickets**: History and active bookings.

2.  **Admin Dashboard (Web) - For Staff**:
    - **Counter Control**: "Call Next", "Recall", "Skip".
    - **Overview**: Total waiting, average wait time.
    - **Settings**: Configure counters and breaks.

## 3. User Personas

1.  **Business Owner (Admin)**: Needs to manage queues, call numbers, and see daily stats.
2.  **End Customer (User)**: Wants to take a number remotely, track status in real-time, and get notified.
3.  **Front Desk Officer**: Needs a simple dashboard to print numbers (for walk-in guests) and manage exceptions.

---

## 3. MVP Features (Phase 1)

### A. Queue Management (Core)

- **Take Queue (Remote)**: User selects service → Gets Number & ETA.
- **Take Queue (On-site)**: Front desk prints QR Code/Paper.
- **Real-time Dashboard**: WebSocket connection updating current number serving.

### B. Notification Engine

- **Channels**: WhatsApp (Priority), Push Notification (App).
- **Triggers**:
  - "Booking Confirmed"
  - "5 People ahead of you" (Warning)
  - "Now Serving You" (Call)

### C. Intelligence (AI Lite)

- **Dynamic ETA**: Calculate Waiting Time based on moving average of last 10 served customers.

---

## 4. Success Metrics (KPIs)

- **Technical**: < 200ms API Latency for Queue Status.
- **Business**: Completion Rate (Booked vs Served). Target > 90% (Low No-Show).

## 5. Non-Functional Requirements

- **Reliability**: Offline-first capability for Business Dashboard (if internet blips).
- **Scalability**: Handle Monday Morning spike (8:00 - 9:00 AM).
