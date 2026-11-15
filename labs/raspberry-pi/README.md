# Raspberry Pi Labs

This section of CloudSQL Lab extends platform engineering concepts beyond local development environments and onto real hardware.

The Raspberry Pi serves as a lightweight Linux platform for learning and practicing:

- Go development
- Linux administration
- SSH-based deployments
- System monitoring
- PostgreSQL
- Podman
- APIs
- Platform engineering workflows

## Current Environment

| Component | Value |
|------------|--------|
| Device | Raspberry Pi 3B+ |
| Access Method | SSH |
| Development Workflow | VS Code → GitHub → Raspberry Pi |
| Operating System | Raspbian Buster |
| Language | Go |

## Development Approach

VS Code Remote SSH was evaluated but not adopted for the current Raspberry Pi environment.

The project currently uses a Git and SSH based workflow:

```text
VS Code (Windows)
        ↓
     Git Push
        ↓
      GitHub
        ↓
 Raspberry Pi
     Git Pull
        ↓
   Build / Run
```

This approach keeps the workflow simple, works reliably with the current Raspberry Pi environment, and reinforces source control and deployment fundamentals.

---

# Raspberry Pi Platform Journey

The Raspberry Pi labs follow a deliberate progression from a simple Go command-line utility to a small platform composed of applications, databases, APIs, containers, and operational workflows.

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

## Current Status

```text
✅ Health Monitor Working
✅ Podman Machine Running
✅ PostgreSQL Container Running

Current Focus
    ↓
🔜 Create health_metrics table
🔜 Store metrics in PostgreSQL
🔜 Query historical metrics
```

The goal is not simply to learn Go, PostgreSQL, or Podman individually.

The objective is to understand how applications, databases, containers, APIs, and operational workflows fit together to form a complete platform.

---

# Current Labs

## Raspberry Pi Health Monitor

A lightweight Go application that collects:

- Hostname
- Uptime
- Load Average
- CPU Temperature
- Memory Usage

Example output:

```text
=== Raspberry Pi Health Monitor ===

Hostname: raspberrypi
Uptime: 1h 25m
Load Average: 0.11 0.04 0.01
CPU Temp: 52.6°C
Memory: 17.2% (159 MB / 923 MB)
```

Current execution model:

```text
VS Code
    ↓
Git Push
    ↓
GitHub
    ↓
Git Pull
    ↓
Build and Run on Raspberry Pi
```

---

# Learning Progression

Each milestone introduces a single new capability while reinforcing concepts learned in previous phases.

```text
Collect Metrics
       ↓
Store Metrics
       ↓
Query Metrics
       ↓
Expose Metrics
       ↓
Containerize Solution
       ↓
Operate Platform
```

---

# Platform Vision

The long-term goal is to evolve a simple Go CLI into a small monitoring platform.

Capabilities:

1. Collect health metrics from a Raspberry Pi
2. Store metrics in PostgreSQL
3. Query historical data
4. Expose metrics through a REST API
5. Run components in containers
6. Operate and maintain the platform

Target architecture:

```text
+--------------------+
| Health Monitor API |
+--------------------+
          |
          v
+--------------------+
| PostgreSQL         |
+--------------------+
```

Containerized deployment:

```text
Windows
    ↓
Podman Machine (cloudsql-lab)
    ↓
Podman
├── health-monitor
└── postgres
```

---

# Roadmap

## Milestone 1 ✅ Foundations

Completed:

- Configure SSH access to Raspberry Pi
- Establish Git-based deployment workflow
- Build a Go-based health monitor
- Run PostgreSQL 16 in Podman
- Publish CloudSQL Lab on LinkedIn

Skills learned:

- Go fundamentals
- Linux administration
- SSH
- Git
- GitHub workflows
- Podman fundamentals
- PostgreSQL basics
- Troubleshooting

---

## Milestone 2 🚧 PostgreSQL Persistence

### Goal

Move beyond terminal output and persist collected metrics in PostgreSQL.

Current:

```text
Read Metrics
      ↓
Print To Terminal
```

Target:

```text
Read Metrics
      ↓
Store In PostgreSQL
```

Architecture:

```text
Health Monitor
       ↓
   PostgreSQL
```

Schema:

```sql
CREATE TABLE health_metrics (
    id SERIAL PRIMARY KEY,
    recorded_at TIMESTAMP NOT NULL,
    hostname TEXT NOT NULL,
    cpu_temp NUMERIC,
    memory_percent NUMERIC,
    load_avg NUMERIC
);
```

Deliverables:

- PostgreSQL database
- health_metrics table
- Metric storage
- SQL verification queries

---

## Milestone 3: Historical Reporting

Focus areas:

- SQL SELECT statements
- Aggregate functions
- Trend analysis
- Reporting workflows

---

## Milestone 4: Configuration Management

Focus areas:

- Configuration files
- Environment separation
- Operational best practices

---

## Milestone 5: REST API

Focus areas:

- HTTP servers
- REST APIs
- JSON serialization
- Service design

---

## Milestone 6: Containerization with Podman

Focus areas:

- Container images
- Podman networking
- Environment variables
- Deployment workflows

---

## Milestone 7: Platform Operations

Focus areas:

- Monitoring
- Troubleshooting
- Observability
- Documentation
- Operational runbooks

---

# Future Milestone: Automated Testing

Current state:

```text
GitHub Actions
    ↓
go test ./...
    ↓
[no test files]
```

Future state:

```text
GitHub Actions
    ↓
Automated Tests
    ↓
Validated Application Behavior
```

Topics:

- `_test.go` files
- Unit tests
- Table-driven tests
- Integration tests
- Testing PostgreSQL interactions
- GitHub Actions integration

Goal:

Replace:

```text
[no test files]
```

with meaningful automated tests that verify application behavior and prevent regressions.

---

# Success Criteria

By completing this roadmap, the Raspberry Pi labs will demonstrate hands-on experience with:

- Go application development
- Linux administration
- PostgreSQL
- SQL and data modeling
- REST APIs
- Podman
- Monitoring systems
- Troubleshooting
- Automated testing
- Documentation
- Platform engineering principles

The final outcome is a small but complete platform engineering lab that evolves from a simple command-line utility into a containerized monitoring platform.