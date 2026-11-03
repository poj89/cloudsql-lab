# cloudsql-lab
A learning project focused on Go, PostgreSQL, Podman, and cloud database automation.

# Goals
- Learn Go through real infrastructure tooling
- Provision PostgreSQL containers using Podman
- Simulate Cloud SQL operations
- Practice testing and CI/CD
- Explore concurrency patterns
- Build a portfolio-quality project

# Phase 1: JSON-backed instance registry

Phase 1 teaches structs, slices, JSON, file I/O, and explicit error handling.
The commands below manage instance metadata only; they do not start containers.

Run them from the project directory:

```powershell
go run ./cmd/cloudsql create dev-db
go run ./cmd/cloudsql list
go run ./cmd/cloudsql delete dev-db
```

The registry is stored in `data/instances.json`. Instance names must start with
a letter or digit and may contain letters, digits, periods, underscores, and
hyphens. Creating an existing name and deleting an unknown name return errors.

## Go design

- `internal/instance` defines and validates the `Instance` struct.
- `internal/storage` reads and writes a slice of instances as a JSON array.
  Writes use a temporary file followed by a rename to reduce the chance of
  leaving a partially written registry.
- `cmd/cloudsql` parses the command-line arguments and reports errors.

## Build

```powershell
go build -o cloudsql.exe ./cmd/cloudsql
```
