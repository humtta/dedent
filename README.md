# Dedent

A Go package to remove shared indentation from a multiline string.

In Go, raw string literals preserve every character from the source, including leading whitespace.
When a multiline string is indented to align with surrounding code, that indentation becomes part of
the value. This package identifies the shared indentation across lines and removes it.

## Requirements

- Go 1.24 or later.

## Installation

Install the latest version of the package by running the following command:

```sh
go get github.com/humtta/dedent@latest
```

## Usage

This package provides only two functions:

- [`D`] accepts a string and returns a new one with the shared indentation removed from each line.
- [`Df`] formats the given string with [`fmt.Sprintf`] and passes the result to [`D`].

Here's an example:

```go
package main

import (
  "fmt"

  "github.com/humtta/dedent"
)

func main() {
  html := `
    <div>
      <h1>Hello, %s!</h1>
    </div>
  `

  fmt.Print(dedent.D(html))
  // Output:
  // <div>
  //   <h1>Hello, %s!</h1>
  // </div>
  //

  fmt.Print(dedent.Df(html, "World"))
  // Output:
  // <div>
  //   <h1>Hello, World!</h1>
  // </div>
  //
}
```

## Documentation

The full API reference is available on [Go Packages].

## License

This project is licensed under the [MIT License].

[`d`]: https://pkg.go.dev/github.com/humtta/dedent#D
[`df`]: https://pkg.go.dev/github.com/humtta/dedent#Df
[`fmt.sprintf`]: https://pkg.go.dev/fmt#Sprintf
[go packages]: https://pkg.go.dev/github.com/humtta/dedent
[mit license]: LICENSE.md
