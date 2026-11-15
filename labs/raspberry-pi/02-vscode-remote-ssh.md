# VS Code Remote SSH on Raspberry Pi 3B+

## Objective

Connect VS Code Remote SSH from a Windows workstation to a Raspberry Pi 3B+ to enable remote Go development directly from VS Code.

---

## Environment

| Component | Value |
|------------|--------|
| Device | Raspberry Pi 3 Model B+ Rev 1.3 |
| Operating System | Raspbian GNU/Linux 10 (Buster) |
| Architecture | armhf (32-bit) |
| glibc Version | 2.28 |
| VS Code Version | 1.144.1 |
| Remote SSH Extension | 0.128.0 |
| Connection Method | SSH |

---

## Problem

VS Code Remote SSH failed to connect with the following error:

```text
The remote host may not meet VS Code Server's prerequisites for glibc and libstdc++
The remote host does not meet the prerequisites for running VS Code Server
```

Standard SSH access worked successfully:

```bash
ssh pi@10.0.0.40
```

---

## Initial Hypotheses

Possible causes included:

- SSH configuration issue
- Authentication problem
- Damaged VS Code Server installation
- Unsupported architecture
- Outdated Linux runtime libraries
- Outdated Raspberry Pi OS release

---

## Troubleshooting Performed

### Verify SSH Connectivity

Confirmed successful SSH access:

```bash
ssh pi@10.0.0.40
```

Result:

```text
SUCCESS
```

### Verify Operating System

```bash
cat /etc/os-release
```

Result:

```text
PRETTY_NAME="Raspbian GNU/Linux 10 (buster)"
```

### Verify Architecture

```bash
dpkg --print-architecture
```

Result:

```text
armhf
```

### Verify glibc Version

```bash
ldd --version
```

Result:

```text
glibc 2.28
```

### Remove Existing VS Code Server Files

```bash
rm -rf ~/.vscode-server
rm -rf ~/.vscode-remote
```

Result:

```text
No change.
Same VS Code Remote SSH error remained.
```

### Update Installed Packages

```bash
sudo apt update
sudo apt upgrade
```

Result:

```text
584 packages upgraded successfully.
```

### Reboot System

```bash
sudo reboot
```

Result:

```text
System rebooted successfully.
```

---

## Findings

The following factors were identified:

### SSH Was Not the Problem

SSH connectivity functioned normally throughout the investigation.

Evidence:

```bash
ssh pi@10.0.0.40
```

worked successfully.

### Operating System Is Very Old

The Raspberry Pi is running:

```text
Raspbian GNU/Linux 10 (Buster)
```

This release is several generations behind current Raspberry Pi OS releases.

### Runtime Libraries Are Old

The system uses:

```text
glibc 2.28
```

Modern VS Code Server versions require more recent Linux runtimes than those typically found on Buster systems.

### 32-bit Platform

The system architecture is:

```text
armhf
```

Current developer tooling increasingly targets newer and 64-bit Linux environments.

---

## Root Cause Assessment

Most likely cause:

```text
VS Code Server prerequisites are not satisfied by the
current Raspberry Pi OS Buster environment.
```

Supporting factors:

- Buster release
- glibc 2.28
- armhf architecture
- Successful SSH connectivity
- Persistent failure after VS Code Server cleanup
- Persistent failure after package upgrades

---

## Decision

For now, continue using:

```text
VS Code on Windows
↓
Git Push
↓
GitHub
↓
Git Pull on Raspberry Pi
↓
Execute via SSH
```

This workflow provides:

- Source control
- Version history
- Remote deployment practice
- No dependency on VS Code Remote SSH

---

## Lessons Learned

### Patch vs Rebuild

Updating packages:

```bash
apt upgrade
```

is not the same as upgrading the operating system itself.

Current state:

```text
Buster + latest packages
```

Desired future state:

```text
Bookworm + latest packages
```

### Real-World Troubleshooting Workflow

The investigation followed a structured approach:

1. Verify connectivity
2. Verify authentication
3. Check operating system
4. Check architecture
5. Check runtime libraries
6. Remove existing installation artifacts
7. Apply package updates
8. Evaluate results
9. Choose a practical workaround

---

## Current Workflow

Development cycle:

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

---

## Future Improvements

Potential future migration:

```text
Raspberry Pi OS Bookworm (32-