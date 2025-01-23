---
title: Server Installation
---

Here's all the different ways you can download and install the server binary

## Go Install

The simplest way to install Pockets if you have Go installed is to use the `go install` command:

```bash
go install github.com/matfire/pockets/server@latest
```

This will download and compile the latest version of the server binary and install it in your `$GOPATH/bin` directory.

## Build from Source

If you want to build from source, you can clone the repository and use the provided Makefile:

```bash
# Clone the repository
git clone https://github.com/matfire/pockets.git
cd pockets

# Build the server binary
make build/server
```

The binary will be created in the root directory of the project.

## Github Releases

You can download pre-built binaries from the [GitHub releases page](https://github.com/matfire/pockets/releases). Choose the appropriate binary for your operating system and architecture.

After downloading:

1. Extract the archive if necessary
2. Move the binary to a location in your system PATH (optional)
3. Make the binary executable (Unix-based systems only):
   ```bash
   chmod +x ./pockets
   ```

## Homebrew

If you're on macOS or Linux, you can install Pockets using Homebrew:

```bash
brew install matfire/matfire/pockets
```

The server binary will be installed and available in your system PATH.

## Configuration

For detailed configuration options and examples, please refer to the [Server Configuration](/references/server) documentation.

Basic configuration can be done using either:
- A configuration file (`pockets.yaml` or `pockets.toml`) in the current directory
- Environment variables prefixed with `PKTS_`

Default configuration values:
- Port: 3000
- Admin User: admin@example.com
- Admin Password: test1234
