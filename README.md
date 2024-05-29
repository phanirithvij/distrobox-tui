# distrobox-tui

[![Go Report Card](https://goreportcard.com/badge/github.com/phanirithvij/distrobox-tui)](https://goreportcard.com/report/github.com/phanirithvij/distrobox-tui)

![screenshot.png](/screenshot.png)

A minimal TUI for [Distrobox](https://github.com/89luca89/distrobox) using [Bubbletea](https://github.com/charmbracelet/bubbletea).

Features [Catppuccin](https://github.com/catppuccin/catppuccin) color palette. Support for theme selection coming in future release.

My intention is to learn the Bubbletea framework by creating something (sort of?) useful.
## Install

### Requirements
* Go 1.18+
* Distrobox

```bash
go install github.com/phanirithvij/distrobox-tui@main
```

## Usage

* Must be run from the host OS
* Ensure `$GOPATH/bin` is in your shell's $PATH

```bash
distrobox-tui
```

Currently it is not possible to *create* Distroboxes in the TUI, but this might be added in the future.

For other planned things see [docs/ROADMAP.md](./docs/ROADMAP.md)

## Project history

This project is a continuation of the [original](https://github.com/hyperreal64/distrobox-tui) project, created by [@hyperreal64](https://github.com/hyperreal64).
