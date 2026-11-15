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

---

# Current Labs

## VS Code Remote SSH Troubleshooting

Documented investigation into VS Code Remote SSH compatibility issues involving:

- ARM architecture
- glibc requirements
- Raspberry Pi OS Buster
- VS Code Server prerequisites

Reference:

```text
02-vscode-remote-ssh.md
```

---

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

The Raspberry Pi labs follow a deliberate platform engineering progression.

Each milestone builds on the previous one while introducing a single new capability.

```text
Collect metrics
    ↓
Store metrics
    ↓
Query metrics
    ↓
Expose metrics
    ↓
Containerize solution
    ↓
Operate platform
```

The goal is not simply to learn Go, PostgreSQL, or Podman individually.

The objective is to understand how these technologies work together as a complete platform.

---

# Platform Vision

The long-term goal is to evolve a simple Go CLI into a small platform that:

1. Collects health metrics from a Raspberry Pi
2. Stores those metrics in PostgreSQL
3. Retrieves and reports historical metrics
4. Exposes data through a REST API
5. Runs inside Podman containers
6. Operates as a complete monitoring platform

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
Podman
├── health-monitor
└── postgres
```

---

# Roadmap

## Milestone 1 ✅

Completed:

- Configure SSH access to Raspberry Pi
- Establish Git-based deployment workflow
- Build first Go-based health monitor
- Document Linux troubleshooting investigations
- Publish CloudSQL Lab on LinkedIn

Skills learned:

- Go fundamentals
- Linux file systems
- SSH
- Git
- GitHub workflows
- Troubleshooting

---

# Milestone 2: Persist Metrics in PostgreSQL

## Goal

Move beyond terminal output by storing collected metrics in PostgreSQL.

Current:

```text
Read metrics
    ↓
Print to terminal
```

Target:

```text
Read metrics
    ↓
Store in PostgreSQL
```

Architecture:

```text
Health Monitor
        ↓
    PostgreSQL
```

Example schema:

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

Learning objectives:

- PostgreSQL fundamentals
- SQL INSERT statements
- Go database/sql package
- Schema design
- Data persistence

Deliverables:

- PostgreSQL database
- health_metrics table
- Insert metric collection into database
- Verify data with SQL queries

---

# Milestone 3: Historical Reporting

## Goal

Generate reports from stored metrics.

Current:

```text
What is the CPU temperature right now?
```

Target:

```text
What happened over the last day, week, or month?
```

Example:

```bash
health-monitor report
```

Example output:

```text
Average CPU Temp: 48.2°C
Maximum CPU Temp: 61.1°C
Average Memory Usage: 21.5%
Samples Recorded: 500
```

Learning objectives:

- SQL SELECT statements
- Aggregate functions
- Data analysis
- Reporting workflows

Deliverables:

- Report command
- Historical statistics
- Summary output

---

# Milestone 4: Configuration Management

## Goal

Remove hard-coded application settings.

Current:

```go
host := "localhost"
```

Target:

```yaml
database:
  host: localhost
  port: 5432
  user: postgres
  dbname: metrics
```

Learning objectives:

- Configuration management
- Environment separation
- Operational best practices

Deliverables:

- YAML configuration file
- Config loader package
- Environment-specific settings

---

# Milestone 5: REST API

## Goal

Expose collected metrics through HTTP endpoints.

Example endpoints:

```text
GET /health
GET /metrics
GET /metrics/latest
```

Example response:

```json
{
  "hostname": "raspberrypi",
  "cpu_temp": 52.6,
  "memory_percent": 17.2,
  "load_avg": 0.11
}
```

Learning objectives:

- HTTP servers
- REST APIs
- JSON serialization
- Service design

Deliverables:

- HTTP server
- JSON responses
- API documentation

---

# Milestone 6: Containerization with Podman

## Goal

Package the application and database as containers.

Container architecture:

```text
Podman
├── health-monitor
└── postgres
```

Example workflow:

```bash
podman build -t health-monitor .
```

```bash
podman run health-monitor
```

Learning objectives:

- Container images
- Podman fundamentals
- Networking
- Environment variables
- Deployment workflows

Deliverables:

- Containerfile
- PostgreSQL container
- Health Monitor container

---

# Milestone 7: Operate the Platform

## Goal

Run a complete monitoring platform.

Architecture:

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

Deployment:

```text
Podman
├── health-monitor
└── postgres
```

Capabilities:

- Metric collection
- Persistent storage
- Historical reporting
- REST API access
- Containerized deployment

Learning objectives:

- Monitoring
- Troubleshooting
- Observability concepts
- Platform operations
- Documentation

Deliverables:

- End-to-end monitoring solution
- Operational runbooks
- Deployment documentation

---

# Development Workflow

Current workflow:

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

Future workflow:

```text
VS Code
    ↓
GitHub
    ↓
Raspberry Pi
    ↓
Podman
    ├── health-monitor
    └── postgres
```

---

# Success Criteria

By completing this roadmap, CloudSQL Lab will demonstrate practical experience with:

- Go application development
- Linux administration
- PostgreSQL
- SQL and database design
- REST APIs
- Podman
- Monitoring systems
- Troubleshooting
- Documentation
- Platform engineering principles

The final outcome is a platform engineering lab that evolves from a simple command-line application into a small, containerized monitoring platform.