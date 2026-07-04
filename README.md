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

This package provides only two functions: [`D`] and [`Df`].

`D` accepts a single string and returns a new one with the common indentation removed from each
line. Blank lines are normalized to empty lines and ignored when calculating the common prefix. If
the first line is blank, it's removed. Here's an example:

```go
package main

import (
  "fmt"

  "github.com/humtta/dedent"
)

func main() {
  html := `
    <div>
      <h1>Hello, World!</h1>
    </div>
  `
  fmt.Print(dedent.D(html))
  // Output:
  // <div>
  //   <h1>Hello, World!</h1>
  // </div>
  //
}
```

`Df` is just a convenience wrapper that formats the input using `fmt.Sprintf` and passes the result
to `D`. Here's another example:

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
  fmt.Print(dedent.Df(html, "World"))
  // Output:
  // <div>
  //   <h1>Hello, World!</h1>
  // </div>
  //
}
```

The indentation is removed based on a common prefix, byte by byte. Mixing tabs and spaces can
prevent a match, leaving the indentation unchanged. For predictable behavior, use a single type of
whitespace character consistently across all lines (like any normal person).

## Documentation

The full API reference is available on [Go Packages].

## Benchmark

The benchmark results below were obtained on an AMD Ryzen 7 5700U (Linux, x64) in a [Devbox]
environment, using the `devbox run bench` command.

| Benchmark       | Operations | Nanoseconds/op | Bytes/op | Allocations/op |
|:---------------:|:----------:|:--------------:|:--------:|:--------------:|
| `BenchmarkD-16` | 5,828,864  | 208.9          | 256      | 1              |

## License

This project is licensed under the [MIT License].

[`d`]: https://pkg.go.dev/github.com/humtta/dedent#D
[`df`]: https://pkg.go.dev/github.com/humtta/dedent#Df
[go packages]: https://pkg.go.dev/github.com/humtta/dedent
[devbox]: https://www.jetify.com/docs/devbox
[mit license]: LICENSE.md
