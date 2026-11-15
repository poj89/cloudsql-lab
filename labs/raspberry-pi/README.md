# Raspberry Pi Platform Journey

The Raspberry Pi labs follow a deliberate progression from a simple Go command-line utility to a small, containerized platform.

```text
┌───────────────────────────┐
│ Phase 1                   │
│ Foundations               │
├───────────────────────────┤
│ ✅ Go Fundamentals        │
│ ✅ Linux via SSH          │
│ ✅ Health Monitor CLI     │
└─────────────┬─────────────┘
              │
              ▼
┌───────────────────────────┐
│ Phase 2                   │
│ Data Persistence          │
├───────────────────────────┤
│ 🔜 PostgreSQL Container   │
│ 🔜 health_metrics Table   │
│ 🔜 Metric Storage         │
└─────────────┬─────────────┘
              │
              ▼
┌───────────────────────────┐
│ Phase 3                   │
│ Historical Insights       │
├───────────────────────────┤
│ 🔜 SQL Queries            │
│ 🔜 Trend Analysis         │
│ 🔜 Reporting Commands     │
└─────────────┬─────────────┘
              │
              ▼
┌───────────────────────────┐
│ Phase 4                   │
│ API Layer                 │
├───────────────────────────┤
│ 🔜 GET /health            │
│ 🔜 GET /metrics           │
│ 🔜 JSON Responses         │
└─────────────┬─────────────┘
              │
              ▼
┌───────────────────────────┐
│ Phase 5                   │
│ Containerization          │
├───────────────────────────┤
│ 🔜 Health Monitor Image   │
│ 🔜 PostgreSQL Container   │
│ 🔜 Podman Networking      │
└─────────────┬─────────────┘
              │
              ▼
┌───────────────────────────┐
│ Phase 6                   │
│ Platform Operations       │
├───────────────────────────┤
│ 🔜 Monitoring             │
│ 🔜 Troubleshooting        │
│ 🔜 Operational Runbooks   │
└─────────────┬─────────────┘
              │
              ▼
┌───────────────────────────┐
│ Future Track              │
├───────────────────────────┤
│ 🔜 Go Unit Tests          │
│ 🔜 Integration Tests      │
│ 🔜 GitHub Actions Tests   │
└───────────────────────────┘
```

**Current Status**

```text
✅ PostgreSQL Container Running
✅ Podman Machine Running
✅ Health Monitor Working

Current Focus
    ↓
🔜 Create health_metrics table
🔜 Store metrics in PostgreSQL
🔜 Query historical metrics
```

The objective is not simply to learn Go, PostgreSQL, or Podman individually. The goal is to understand how applications, databases, containers, APIs, and operational workflows fit together to form a complete platform.