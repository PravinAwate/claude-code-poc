# CI/CD Quality & Security Standards

## Code Quality Rules
- All HTTP endpoints must explicitly log incoming requests.
- Avoid using empty or unhandled writer fragments.

## Security Audit Rules
- 🛑 CRITICAL: No strings matching API tokens or keys (`sk_...`) are allowed in source code.
