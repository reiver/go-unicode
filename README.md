# go-unicode

Package unicode provides tools for working with Unicode characters, for the Go programming language.

## Documention

Online documentation, which includes examples, can be found at: http://godoc.org/github.com/reiver/go-unicode

[![GoDoc](https://godoc.org/github.com/reiver/go-unicode?status.svg)](https://godoc.org/github.com/reiver/go-unicode)


## Example
```
import (
	"github.com/reiver/go-unicode"
)

// ...

var r rune

// ...

switch r {
case unicode.BackSpace:
	// ...
case unicode.HorizontalTab:
	// ...
case unicode.LineFeed:
	// ...
case unicode.Delete:
	// ...
default:
	// ...
}
```

## Import

To import package **unicode** use `import` code like the follownig:
```
import "github.com/reiver/go-unicode"
```

## Installation

To install package **unicode** do the following:
```
GOPROXY=direct go get https://github.com/reiver/go-unicode
```

## Author

Package **unicode** was written by [Charles Iliya Krempeaux](http://reiver.link)
