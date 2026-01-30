# Frontend Engineering Tickets 🖥️

## Phase 1: Admin Dashboard (SvelteKit)

### [QTIME-005] Scaffold Admin Web [DONE]
**Type**: Chore | **Points**: 3
**Description**: Initialize SvelteKit project in `apps/admin-web` with TypeScript.
**Success Criteria**:
- [ ] TailwindCSS & Shadcn-Svelte configured.
- [ ] Base layout with sidebar and breadcrumbs.

### [QTIME-007] Admin Authentication UI
**Type**: Feature | **Points**: 5
**Description**: Secure login page for Staff/Admins.
**Success Criteria**:
- [ ] Form validation (Zod + Superforms).
- [ ] JWT handling via HttpOnly cookies (Server-side load).

### [QTIME-008] Real-time Dashboard (Socket.io)
**Type**: Feature | **Points**: 8
**Description**: Live queue display for staff counters.
**Success Criteria**:
- [ ] Svelte Store integration for live state management.
- [ ] WebSocket connection to Go Queue Engine.

---

## Phase 2: Public Booking PWA (Next.js)

### [QTIME-006] Scaffold Booking PWA [DONE]
**Type**: Chore | **Points**: 3
**Description**: Initialize Next.js 14+ (App Router) in `apps/public-mobile`.
**Success Criteria**:
- [ ] PWA Manifest and Service Worker (next-pwa).
- [ ] Mobile-first UI components.

### [QTIME-015] Ticket Booking Flow
**Type**: Feature | **Points**: 8
**Description**: End-to-end user flow for taking a queue number.
**Success Criteria**:
- [ ] Merchant discovery (Search/List).
- [ ] Ticket generation API call & Success state with QR Code.

### [QTIME-017] Offline-First Support
**Type**: Feature | **Points**: 5
**Description**: Allow users to see their ticket even without internet.
**Success Criteria**:
- [ ] Cache active ticket in IndexedDB/LocalStorage.
- [ ] Background sync for status updates.

---

## Phase 3: Analytics & Final Polish

### [QTIME-016] Business Intelligence Dashboard
**Type**: Feature | **Points**: 5
**Description**: Visual data reports for Merchant Admins.
**Success Criteria**:
- [ ] Integration with Java Core Backend Analytics API.
- [ ] Responsive charts using Chart.js or Recharts.

### [QTIME-018] Performance & SEO Audit
**Type**: QA | **Points**: 3
**Description**: Ensure PWA score is 100 on Lighthouse.
