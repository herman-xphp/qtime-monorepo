# Frontend Architecture 🖥️

Q-Time utilizes a **Dual-Frontend Strategy** to catering to two very different user personas.

## 1. Application Breakdown

| App                 | Persona               | Tech Stack          | Type                 | Responsibility                                  |
| :------------------ | :-------------------- | :------------------ | :------------------- | :---------------------------------------------- |
| **Admin Dashboard** | Clinic Staff, Tellers | **SvelteKit**       | SPA (CSR/SSR hybrid) | Queue Control, Reporting, Configuration.        |
| **Booking PWA**     | End Users (Patients)  | **Next.js** (React) | PWA (Mobile First)   | Booking numbers, Checking ETA, Realtime Status. |

---

## 2. Admin Dashboard (`apps/admin-web`)

- **Why SvelteKit?**
  - **Reactivity**: Handling real-time queue updates (WebSocket) in Svelte is much simpler and more performant than React (`store` vs `useEffect`).
  - **Bundle Size**: Compiler-based approach results in tiny bundles, ideal for slow clinic computers.
- **Key Libraries**:
  - UI: `shadcn-svelte` + TailwindCSS.
  - State: Svelte Stores (Native).
  - Chart: `Chart.js` for Analytics.

## 3. Booking PWA (`apps/public-mobile`)

- **Why Next.js?**
  - **Ecosystem**: React has the best ecosystem for mapping (Google Maps components), QR Scanning, etc.
  - **SEO**: Critical for "Klinik Gigi Terdekat" discovery via Google Search.
- **Key Features**:
  - **Offline Mode**: Service Workers to view ticket even when offline.
  - **Push Notifications**: FCM Integration.

---

## 4. Shared Strategy (Monorepo)

### Design System

We share a design token set (Tailwind Config) to ensure consistent branding:

- Primary Color: `Teal-600`
- Font: `Inter`
- Spacing: Standard `4px` grid.

### API Consumption

Both apps use the **BFF (Backend for Frontend)** pattern implicitly via their server-side load functions, or consume the Core API directly.

- **Client**: Generated using `openapi-typescript` from our `openapi.yaml`.
- **Auth**: HttpOnly Cookies (JWT) set by the API Gateway or Proxy.

---

## 5. Directory Structure

```text
apps/
├── admin-web/          # SvelteKit
│   ├── src/routes/     # File-based Routing
│   ├── src/lib/        # Components
│   └── static/
├── public-mobile/      # Next.js
│   ├── app/            # App Router
│   ├── components/
│   └── public/
```
