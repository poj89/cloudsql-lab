# cloudsql-lab

CloudSQL Lab is my hands-on platform engineering lab.

I started this project to learn Go by building real tools instead of solving isolated coding exercises. Over time, it has grown into a playground for exploring PostgreSQL, Podman, Linux, APIs, automation, and platform engineering concepts through a single evolving codebase.

The goal is not to build a production system. The goal is to understand how infrastructure, databases, containers, APIs, and automation work together by building them step by step.

## What I'm Learning

- Go fundamentals through practical tools
- PostgreSQL administration and automation
- Podman container workflows
- Linux and Raspberry Pi operations
- API development
- Testing and CI/CD
- Platform engineering concepts
- Monitoring and observability
- Infrastructure automation

## Learning Philosophy

Rather than starting multiple unrelated projects, I am building one project that grows in complexity as my skills improve.

Each milestone introduces a new concept while reinforcing previous lessons.

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

## Current Development Workflow

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

## Project Areas

### PostgreSQL Automation

Build tooling that simulates common cloud database operations and lifecycle management tasks.

Topics:

- Instance management
- Database provisioning
- Configuration management
- Operational workflows

### Raspberry Pi Labs

Use a Raspberry Pi 3B+ as a constrained Linux environment for experimentation.

Topics:

- Linux administration
- SSH workflows
- System monitoring
- Go applications
- Real-world troubleshooting

### Platform Engineering Labs

Explore how applications, databases, containers, and automation fit together as a platform.

Topics:

- PostgreSQL
- Podman
- APIs
- Monitoring
- Documentation
- Operations

## Phase 1: JSON-Backed Instance Registry

The first phase focuses on foundational Go concepts:

- Structs
- Slices
- JSON
- File I/O
- Error handling

The commands below manage instance metadata only. No containers or databases are started at this stage.

Run from the project root:

```powershell
go run ./cmd/cloudsql create dev-db
go run ./cmd/cloudsql list
go run ./cmd/cloudsql delete dev-db
```

The registry is stored in:

```text
data/instances.json
```

Instance names must start with a letter or digit and may contain:

- Letters
- Digits
- Periods
- Underscores
- Hyphens

Creating an existing instance or deleting a non-existent instance returns an error.

## Current Design

### internal/instance

Responsible for:

- Instance representation
- Validation rules
- Business logic related to instances

### internal/storage

Responsible for:

- Reading instance data
- Writing instance data
- JSON serialization

Writes are performed using a temporary file followed by a rename operation to reduce the risk of leaving behind a partially written registry.

### cmd/cloudsql

Responsible for:

- Command-line parsing
- User interaction
- Error reporting

## Raspberry Pi Roadmap

Current project:

```text
Health Monitor
```

Current capabilities:

- Hostname
- Uptime
- Load Average
- CPU Temperature
- Memory Usage

Future roadmap:

### Milestone 2

Store collected metrics in PostgreSQL.

```text
Health Monitor
        ↓
    PostgreSQL
```

### Milestone 3

Query historical metrics and generate reports.

```text
Metrics
    ↓
Reports
```

### Milestone 4

Introduce configuration management.

```text
config.yaml
```

### Milestone 5

Expose metrics through a REST API.

```text
GET /health
GET /metrics
```

### Milestone 6

Containerize the application using Podman.

```text
Podman
├── health-monitor
└── postgres
```

### Milestone 7

Operate a complete monitoring platform.

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

## Build

```powershell
go build -o cloudsql.exe ./cmd/cloudsql
```