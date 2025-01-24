---
title: CLI Installation
---

Here's all the different ways you can download and install the CLI binary

## Go Install

The simplest way to install the Pockets CLI if you have Go installed is to use the `go install` command:

```bash
go install github.com/matfire/pockets/cli@latest
```

This will download and compile the latest version of the CLI binary and install it in your `$GOPATH/bin` directory.

## Build from Source

If you want to build from source, you can clone the repository and use the provided Makefile:

```bash
# Clone the repository
git clone https://github.com/matfire/pockets.git
cd pockets

# Build the CLI binary
make build/cli
```

The binary will be created in the root directory of the project.

## Github Releases

You can download pre-built binaries from the [GitHub releases page](https://github.com/matfire/pockets/releases). Choose the appropriate binary for your operating system and architecture.

After downloading:

1. Extract the archive if necessary
2. Move the binary to a location in your system PATH (optional)
3. Make the binary executable (Unix-based systems only):
   ```bash
   chmod +x ./pockets-cli
   ```

## Homebrew

If you're on macOS or Linux, you can install the Pockets CLI using Homebrew:

```bash
brew install matfire/matfire/pocketsctl
```
