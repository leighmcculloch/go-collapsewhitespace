# collapsewhitespace

[![Build Status](https://img.shields.io/travis/leighmcculloch/go-collapsewhitespace.svg?label=linux%20%26%20osx)](https://travis-ci.org/leighmcculloch/go-collapsewhitespace)
[![Go Report Card](https://goreportcard.com/badge/github.com/leighmcculloch/go-collapsewhitespace)](https://goreportcard.com/report/github.com/leighmcculloch/go-collapsewhitespace)
[![Go docs](https://img.shields.io/badge/godoc-reference-blue.svg)](https://godoc.org/4d63.com/collapsewhitespace)

Efficiently collapse multiple whitespace characters in a string into a single space.

## Usage

```go
import "4d63.com/collapsewhitespace"
```

```go
collapsewhitespace.String("a  b	\nc") // "a b c"
```
