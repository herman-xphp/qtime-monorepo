# Frontend Engineering Tickets 🖥️

## Phase 1: Admin Dashboard (SvelteKit)

### [QTIME-005] Scaffold Admin Web

**Type**: Chore | **Story Points**: 3
**Description**: Initialize SvelteKit project in `apps/admin-web`.
**Tech Notes**:

- Use `shadcn-svelte` for UI components.
  **Acceptance Criteria**:
- [ ] `npm run dev` starts app on port 5173.
- [ ] TailwindCSS configured.

### [QTIME-007] Admin Authentication UI

**Type**: Feature | **Story Points**: 5
**Description**: specific Login Page for Clinic Staff.
**Acceptance Criteria**:

- [ ] Login Form (Email/Password).
- [ ] Integration with Spring Boot Login API.
- [ ] Handle HttpOnly Cookie.

### [QTIME-008] WebSocket Client Integration

**Type**: Feature | **Story Points**: 5
**Description**: Connect to Go Queue Engine WebSocket.
**Acceptance Criteria**:

- [ ] Auto-reconnect on disconnect.
- [ ] Toast notification on new ticket.

### [QTIME-014] Queue Management Dashboard

**Type**: Feature | **Story Points**: 8
**Description**: The main grid view of active queues.
**Acceptance Criteria**:

- [ ] "Call Next" button sends API request.
- [ ] Real-time updates via WebSocket.

## Phase 2: Public Mobile (Next.js)

### [QTIME-006] Scaffold Booking PWA

**Type**: Chore | **Story Points**: 3
**Description**: Initialize Next.js 14 project in `apps/public-mobile`.
**Acceptance Criteria**:

- [ ] `npm run dev` starts app on port 3005.
- [ ] Manifest.json configured for PWA.

### [QTIME-015] Booking Flow

**Type**: Feature | **Story Points**: 8
**Description**: User flow to find clinic and take a number.
**Acceptance Criteria**:

- [ ] Location Search Page.
- [ ] Ticket Success Page with QR Code.

## Phase 3: Advanced

### [QTIME-016] Reporting UI

**Type**: Feature | **Story Points**: 5
**Description**: Visualize daily status using Chart.js.

### [QTIME-017] Offline Support

**Type**: Feature | **Story Points**: 5
**Description**: Service Worker implementation for PWA.
