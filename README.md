HTML StripTags for Go
=====================

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
    strip "github.com/grokify/html-strip-tags-go"
)

func main() {
	original := "<h1>Hello World</h1>"
	stripped := strip.StripTags(original) // => "Hello World"
}
```
