# `talosctx`: Get More Out of Talos contexts

![Latest GitHub release](https://img.shields.io/github/release/bdronneau/talosctx.svg) [![CI](https://github.com/bdronneau/talosctx/actions/workflows/ci.yaml/badge.svg)](https://github.com/bdronneau/talosctx/actions/workflows/ci.yaml) [![Go Report Card](https://goreportcard.com/badge/github.com/bdronneau/talosctx)](https://goreportcard.com/report/github.com/bdronneau/talosctx)

**talosctx** is a tool for make easy switch between [Talos](https://www.talos.dev/) contexts. This tool is based on [kubectx/kubens](https://github.com/ahmetb/kubectx) work.

## Features

- **TUI**: [pterm](https://github.com/pterm/pterm)
- **Built with Golang**

![](./img/demo.gif)

## Installation

```bash
go install github.com/bdronneau/talosctx/cmd/talosctx@latest
```

or check [binaries releases](https://github.com/bdronneau/talosctx/releases)

## Usage

## Configuration

**talosctx** use `TALOSCONFIG` environment variable or fallback to `$HOME/.talos/config` to retrieve talosconfig file and contextes.

### Examples

```bash

# Display interactive select between contexts
talosctx

# Get current context (PS1 usage, no return line)
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
