# GoEnv

![GitHub go.mod Go version](https://img.shields.io/github/go-mod/go-version/gookit/gcli?style=flat-square)
[![Unit-Tests](https://github.com/gookit/gcli/actions/workflows/go.yml/badge.svg)](https://github.com/gookit/gcli/actions/workflows/go.yml)
[![GitHub tag (latest SemVer)](https://img.shields.io/github/tag/gookit/gcli)](https://github.com/gookit/gcli)
[![GoDoc](https://godoc.org/github.com/gookit/gcli?status.svg)](https://pkg.go.dev/github.com/gookit/gcli/v3)
[![Go Report Card](https://goreportcard.com/badge/github.com/gookit/gcli)](https://goreportcard.com/report/github.com/gookit/gcli)

Go multi version env manager

- features TODO

> **[中文说明](README.zh-CN.md)**

## Install

**Use go install**

```shell
go install github.com/inherelab/goenv/cmd/goenv
```

## Usage

```shell
goenv
```

switch version:

```shell
goenv switch 1.16
```
Or:

```shell
goenv use 1.16
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

- https://github/gookit/color
- https://github/gookit/config
- https://github/gookit/gcli
- https://github/gookit/goutil

## LICENSE

[MIT](LICENSE)
