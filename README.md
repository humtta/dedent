# Dedent

A Go package to remove common leading whitespace from every line of a multiline string.

## Installation

This package requires Go 1.24 or later. To add it to your module, run the following command from the
module root:

```sh
go get github.com/humtta/dedent@latest
```

## Documentation

The full API reference is available on [Go Packages].

## Benchmark

The benchmark results below were obtained on an AMD Ryzen 7 5700U (Linux, x64) in a [Devbox]
environment, using the `devbox run bench` command.

```txt
BenchmarkD-16    5828864    208.9 ns/op    256 B/op    1 allocs/op
```

## License

This project is licensed under the [MIT License].

[go packages]: https://pkg.go.dev/github.com/humtta/dedent
[devbox]: https://www.jetify.com/docs/devbox
[mit license]: LICENSE.md
