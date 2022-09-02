# GoEnv

![GitHub go.mod Go version](https://img.shields.io/github/go-mod/go-version/inherelab/goenv?style=flat-square)
[![Unit-Tests](https://github.com/inherelab/goenv/actions/workflows/go.yml/badge.svg)](https://github.com/inherelab/goenv/actions/workflows/go.yml)
[![GitHub tag (latest SemVer)](https://img.shields.io/github/tag/inherelab/goenv)](https://github.com/inherelab/goenv)
[![GoDoc](https://godoc.org/github.com/inherelab/goenv?status.svg)](https://pkg.go.dev/github.com/inherelab/goenv/v3)
[![Go Report Card](https://goreportcard.com/badge/github.com/inherelab/goenv)](https://goreportcard.com/report/github.com/inherelab/goenv)

Go multi version env manager

- features TODO

![goenv](_example/help.png)

> **[EN README](README.md)**

## Install

**Use go install**

```shell
go install github.com/inherelab/goenv/cmd/goenv
```

## Usage

```shell
goenv
```

### Switch

Switch to a exists version

```shell
goenv switch 1.16
```
Or:

```shell
goenv use 1.16
```

### Install

Install new version

```shell
goenv install 1.18
```

### Update

```shell
goenv update 1.18
```

### Uninstall

```shell
goenv uninstall 1.18
```

## Development

### Clone

```shell
go clone https://github.com/inherelab/goenv
cd goenv
```

### Run

```bash
go run ./cmd/goenv
```

### Install

```bash
go install ./cmd/goenv
```

## Base on

- https://github.com/gookit/color
- https://github.com/gookit/config
- https://github.com/gookit/gcli
- https://github.com/gookit/goutil

## LICENSE

[MIT](LICENSE)
