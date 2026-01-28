# Security Guidelines & Standards 🛡️

## 1. Authentication & Authorization

- **Standard**: OAuth2 / OpenID Connect (OIDC).
- **Token**: JWT (JSON Web Tokens).
  - **Access Token**: Short-lived (15 min).
  - **Refresh Token**: HTTP-Only Cookie (Secure, SameSite).
- **RBAC (Role Based Access Control)**:
  - `ROLE_SUPER_ADMIN` (Platform Owner)
  - `ROLE_MERCHANT_ADMIN` (Clinic Owner)
  - `ROLE_COUNTER_STAFF` (Operator)
  - `ROLE_USER` (End Customer)

## 2. Network Security

- **TLS/SSL**: Wajib untuk semua komunikasi (In-transit encryption).
- **Internal Network**: Database port (5432, 6379) TIDAK BOLEH terekspos ke public internet. Hanya bisa diakses via private network docker/k8s.
- **API Gateway**: Implement Rate Limiting (Token Bucket Algorithm) di layer depan (Traefik/Nginx) untuk mencegah DDoS.

## 3. Data Protection

- **PII (Personally Identifiable Information)**: Data NIK, No HP, Email harus dienkripsi di database (AES-256) atau minimal hashing jika tidak perlu decrypt.
- **Sanitization**: Semua input user di-sanitize untuk mencegah SQL Injection (walaupun pake ORM) dan XSS.

## 4. Dependencies

- Use `npm audit` and `govulncheck` in CI/CD pipeline to detect vulnerable libraries.
- Docker images should use slim/alpine versions and scanned for vulnerabilities (Trivy).
