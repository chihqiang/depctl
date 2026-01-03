# depctl

`depctl` — Push it, roll it, own it.

A powerful, fast, and reliable remote deployment tool for modern applications. Deploy to multiple hosts simultaneously, manage versions with ease, and rollback instantly when needed.

## Features

- 🚀 **Multi-host deployment** - Deploy to multiple servers in one command
- 🔄 **Version management** - Track and manage deployment versions
- ⚡ **Quick rollback** - Instantly rollback to any previous version
- 🔧 **Deployment hooks** - Execute pre/post deployment commands
-  **Smart packaging** - Include/exclude files intelligently
- 🔐 **Flexible authentication** - Support SSH key and password authentication
- 📊 **Deployment history** - View deployment history across all hosts
- 🌐 **Environment variables** - Configure via environment variables

## Installation

### Build from source

```bash
git clone https://github.com/chihqiang/depctl.git
cd depctl&& go build -o depctl
```

## Quick Start

```bash
# Deploy to multiple hosts
depctl publish --hosts "root:123456@127.0.0.1:8022" --dir . --version "v1.0.0"

# View deployment history
depctl --hosts "root:123456@127.0.0.1:8022" history

# Rollback to previous version
depctl rollback --hosts "root:123456@127.0.0.1:8022" --version "v1.0.0"
```

## Commands

### publish

Deploy your application to remote hosts.

```bash
depctl publish [options]
```

### history

View deployment history across all hosts.

```bash
depctl history
```

### rollback

Rollback to a previous deployment version.

```bash
depctl rollback --version VERSION
```

## Options

### Global Options

These options are available for all commands:

- `--hosts string` - List of remote hosts (format: `user[:password]@host[:port]`) [$DEPCTL_HOSTS]
- `--key string` - Path to SSH private key [$DEPCTL_KEY]
- `--passphrase string` - Passphrase for SSH private key [$DEPCTL_PASSPHRASE]
- `--timeout duration` - SSH connection timeout (default: 30s) [$DEPCTL_TIMEOUT]
- `--hook-pre-host string` - Remote command to run before deployment [$DEPCTL_HOOK_PRE]
- `--hook-post-host string` - Remote command to run after deployment [$DEPCTL_HOOK_POST]
- `--remote-repo string` - Remote deployment repository path (default: "/data/wwwroot/{basename}/releases")
- `--current-link string` - Symbolic link path to current version (default: "/data/wwwroot/{basename}/current")

### Publish Command Options

- `--dir string` - Local directory to deploy (default: current directory)
- `--include string` - Files/directories to include when packaging
- `--exclude string` - Files/directories to exclude when packaging
- `--version string` - Version tag (default: timestamp format)

### Rollback Command Options

- `--version string` - Version to rollback to

## Deployment Structure

depctl uses a standard deployment structure on remote hosts:

```
/data/wwwroot/{project-name}/
├── releases/
│   ├── 20241201123456/
│   ├── 20241201123500/
│   └── 20241201130000/
└── current -> /data/wwwroot/{project-name}/releases/20241201130000
```

## Environment Variables

All options can be configured via environment variables:

- `DEPCTL_HOSTS` - Remote hosts list
- `DEPCTL_KEY` - SSH private key path
- `DEPCTL_PASSPHRASE` - SSH key passphrase
- `DEPCTL_TIMEOUT` - SSH connection timeout
- `DEPCTL_HOOK_PRE` - Pre-deployment hook command
- `DEPCTL_HOOK_POST` - Post-deployment hook command

### Permission Issues

Ensure the deployment user has:
- SSH access to remote hosts
- Write permissions to the deployment directory
- Execute permissions for hook commands

## Contributing

Feel free to submit Issues and Pull Requests!

## License

Apache License 2.0

---

Made with ❤️ for fast, reliable deployments.