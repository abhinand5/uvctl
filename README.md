# uvctl

Manage Python virtual environments created with [uv](https://github.com/astral-sh/uv) from a central location.

uvctl keeps all your environments in one place (`~/.local/uvctl/envs` by default), provides seamless activation that works like conda, and stays out of your way. No magic, no project-local configs, just simple environment management.

## Install

**Quick install (Linux/macOS):**

```bash
curl -fsSL https://github.com/abhinand5/uvctl/releases/latest/download/uvctl-$(uname -s | tr '[:upper:]' '[:lower:]')-$(uname -m | sed 's/x86_64/amd64/;s/aarch64/arm64/') -o uvctl \
  && chmod +x uvctl && sudo mv uvctl /usr/local/bin/
```

**With Go:**

```bash
go install github.com/abhinand5/uvctl@latest
```

**From source:**

```bash
git clone https://github.com/abhinand5/uvctl.git
cd uvctl && make install
```

## Setup

Add the shell hook to your `~/.bashrc` or `~/.zshrc`:

```bash
eval "$(uvctl hook bash)"  # or zsh
```

This enables seamless `uvctl activate` and `uvctl deactivate` commands. Restart your shell or run `source ~/.bashrc` to apply.

## Quick Start

```bash
# Check everything is set up correctly
uvctl doctor

# Create an environment with Python 3.12
uvctl create myproject 3.12

# Activate it
uvctl activate myproject

# Do your work...
python --version  # Python 3.12.x from your env
uv pip install numpy # This command installs in the uvctl managed env

# Deactivate when done
uvctl deactivate
```

## Commands

| Command | Description |
|---------|-------------|
| `uvctl create <name> <python>` | Create a new environment with specified Python version |
| `uvctl activate <name>` | Activate an environment |
| `uvctl deactivate` | Deactivate the current environment |
| `uvctl ls` | List all environments |
| `uvctl which` | Print path to the active environment |
| `uvctl delete <name>` | Delete an environment |
| `uvctl doctor` | Diagnose setup issues |
| `uvctl hook <shell>` | Print shell integration code |
| `uvctl version` | Print version info |

## Configuration

uvctl uses environment variables for configuration:

| Variable | Description | Default |
|----------|-------------|---------|
| `UVCTL_ROOT` | Directory where environments are stored | `~/.local/uvctl/envs` |
| `UVCTL_ACTIVE` | Currently active environment (set automatically) | - |

## Environment Layout

All environments live under `$UVCTL_ROOT`:

```
~/.local/uvctl/envs/
├── myproject/
│   └── .venv/
│       ├── bin/
│       │   ├── activate
│       │   └── python
│       └── ...
└── another-env/
    └── .venv/
        └── ...
```

## Requirements

- [uv](https://docs.astral.sh/uv/getting-started/installation/) must be installed and in your PATH

## License

MIT
