VS Code Remote SSH on Raspberry Pi 3B+
Environment
Raspberry Pi 3B+
Raspbian Buster
armhf (32-bit)
glibc 2.28
Problem
VS Code Remote SSH failed with:

"The remote host does not meet the prerequisites for running VS Code Server"

Troubleshooting
Verified:

SSH connectivity
User authentication
glibc version
architecture
removal of ~/.vscode-server
Findings
The Raspberry Pi was running:

Raspbian GNU/Linux 10 (Buster)
glibc 2.28
armhf
These versions are likely too old for modern VS Code Server requirements.

Current Decision
Continue with:

VS Code on Windows
Git push/pull
SSH access
Future option:

Migrate to Raspberry Pi OS Bookworm
