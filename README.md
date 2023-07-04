# Veribi CLI

Command line interface for Veribi platform

[![License MIT](https://img.shields.io/badge/license-MIT-blue.svg?style=flat-square)](/LICENSE)
[![Go Version](https://img.shields.io/github/go-mod/go-version/vaclav-dvorak/veribi-cli.svg?style=flat-square)](go.mod)

## 🚀 Quick start

Get the latest Veribi CLI release:

Automatic install for Linux/macOS with curl. This script downloads the CLI and puts it in your `/usr/local/bin`

```sh
curl -fsSL https://raw.githubusercontent.com/vaclav-dvorak/veribi-cli/main/scripts/install.sh | sh
```

### 🍺 Homebrew

You can use our private tap to install cli via homebrew.

```sh
brew tap vaclav-dvorak/tap
brew install veribi-cli
```

## 📖 Command overview

| Command | Alias | Description |
| --- | --- | --- |
| `init` | `i` | Saves initial config values. |
| `login` | `l` | Logs into Veribi platform and save token for further usage. |
| `offers` | `o` | Lists all currently available offers. |
| `version` | `v` | Outputs version of CLI and notify about possible update. |

## `offers` command

Lists all currently available offers.

```bash
veribi offers [flags]
```

### Options

- `-t`, `--ths <bool>`

  sort output table by THS/$ value

- `-a`, `--add-auctions <bool>`

  add auctions to the list of offers

## `version` command

Outputs version of CLI and notify about possible update.

```bash
veribi version [flags]
```

### Options

- `-j`, `--json <bool>`

  get output as JSON object
