# `talosctx`: Get More Out of Talos

Based on the work done by kubectx, talosctx aims to assist you in your daily actions on [Talos](https://www.talos.dev/) clusters.

## Features

- **Built with Golang**
- **TUI**: [pterm](https://github.com/pterm/pterm)

## Installation

```bash
go install github.com/lederniermetre/talosctx/cmd/talosctx@latest
```

or check [binaries release](https://github.com/lederniermetre/talosctx/releases)

## Usage

## Configuration

`talosctx` use TALOSCONFIG environment variable or fallback to `$HOME/.talos/config` to retrieve talosconfig file and contextes.

### Examples

```bash

# Display interactive select between contexts
talosctx

# Get current context (PS1 usage)
talosctx --get-context
```

## Development

### Requirements

- [goreleaser](https://goreleaser.com/) for build snapshot version
- [Task](https://taskfile.dev/) for devEx

### Get started

```bash
task init
task tests
```

### Help

In order to retrieve all commands

```bash
task -l
```
