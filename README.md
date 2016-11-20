HTML StripTags for Go
=====================

[![Build Status][build-status-svg]][build-status-link]
[![Docs][docs-godoc-svg]][docs-godoc-link]
[![License][license-svg]][license-link]

This is a Go package containing the unexported `stripTags` function in `html/template/html.go`.

## Background

* The `stripTags` function in `html/template/html.go` is very useful, however, it is not exported.
* Requests were made [on GitHub](https://github.com/golang/go/issues/5884) without success.
* This package is a repo for work done by [Christopher Hesse](https://github.com/christopherhesse) to extract this in this [Gist](https://gist.github.com/christopherhesse/d422447a086d373a967f).

## Installation

```bash
$ go get github.com/grokify/html-strip-tags-go
```

## Usage

```go
import(
    "github.com/grokify/html-strip-tags-go" // => strip
)

func main() {
    original := "<h1>Hello World</h1>"
    stripped := strip.StripTags(original) // => "Hello World"
}
```

 [build-status-svg]: https://api.travis-ci.org/grokify/html-strip-tags-go.svg?branch=master
 [build-status-link]: https://travis-ci.org/grokify/html-strip-tags-go
 [docs-godoc-svg]: https://img.shields.io/badge/docs-godoc-blue.svg
 [docs-godoc-link]: https://godoc.org/github.com/grokify/html-strip-tags-go
 [license-svg]: https://img.shields.io/github/license/grokify/html-strip-tags-go.svg
 [license-link]: https://github.com/grokify/html-strip-tags-go/blob/master/LICENSE
