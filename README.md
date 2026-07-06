# Dedent

A Go package to remove common leading whitespace from every line of a multiline string.

In Go, string literals preserve every character in the source, including spaces and tabs. When a
multiline literal is indented to align with surrounding code, that whitespace becomes part of the
string value. This package computes the longest common indentation prefix across non-blank lines and
removes it from each line.

## Requirements

- Go 1.24 or later.

## Installation

To install the latest version of this package, run the following command:

```sh
go get github.com/humtta/dedent@latest
```

Then, import it into your code:

```go
import "github.com/humtta/dedent"
```

## Usage

This package provides only two functions:

- [`D`] accepts a string and returns a new one with the common indentation removed from each line.
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
