# zerrors 
[![](https://github.com/JavierZunzunegui/zerrors/workflows/Go/badge.svg)](https://github.com/JavierZunzunegui/zerrors/actions?query=workflow%3AGo)
[![GoDoc](https://godoc.org/github.com/JavierZunzunegui/zerrors?status.svg)](http://godoc.org/github.com/JavierZunzunegui/zerrors) 
[![Report card](https://goreportcard.com/badge/github.com/JavierZunzunegui/zerrors)](https://goreportcard.com/report/github.com/JavierZunzunegui/zerrors)

Package zerrors provides additional functionality on top of the go 1.13 error wrapping features,
particularly frame information and flexibility over the wrapped error formatting, without sacrificing on performance.

## Packages

-   [github.com/JavierZunzunegui/zerrors](https://godoc.org/github.com/JavierZunzunegui/zerrors): the primary 
    package, it is through this than one can wrap, inspect and serialise errors. Most users should only require this.
-   [github.com/JavierZunzunegui/zerrors/zmain](https://godoc.org/github.com/JavierZunzunegui/zerrors/zmain): 
    an auxiliary package to configure default behaviour, namely the format of `Error()`, format's `%+v`, as well as
    disabling frame capturing for improved performance. Most users will never require this, and for those who do use it
    it should be only used within the package initialisation phase (`init()` or in a global `var`).

## How to use it

Replace any existing
```go
if err != nil {
    return err
}
```
Or
```go
if err != nil {
    return fmt.Errorf("additional info: %w", err)
}
```

Or any similar pattern, with:

```go
if err != nil {
    return zerrors.SWrap(err, "additional info")
}
```
Or
```go
if err != nil {
    return zerrors.Wrap(err, &CustomErrorType{})
}
```

Global errors should also be replaced to use `zerrors.New` or `zerrors.SNew`.

The resulting error's `Error() string` method is of format `last message: ...: first wrap message: base message`.
A more detailed message is produced by `zerrors.Detail(error) string` or via the `%+v` pattern of the `fmt` methods, of 
the form `last message (file.go:LINE): ... : first wrap message (file.go:LINE): base message (file.go:LINE)`.

Alternative serialisation of the error is also possible, either via a custom method (say, `MyFormat(error) string`) or 
by changing the default encodings via [zmain](https://godoc.org/github.com/JavierZunzunegui/zerrors/zmain). 
This way one can produce any alternative message (such as `last message - ... - first wrap message - base message`, or 
any variant with frame information), using the same inputs as the default encoding uses.

The `errors.Is`, `errors.As` and `errors.Unwrap` methods from go1.13 are supported as expected,
and are intended to remain the primary means to examine the contents of errors.
Two new additional methods, `zerrors.Value(error) error` and `zerros.Frame(error) (runtime.Frame, bool)` are introduced
and can be used in conjunction with `errors.Unwrap` to implement more sophisticated functionality, such as the 
alternative serialisation formats.
Error inspection via type assertion (discouraged since go1.13) is further discouraged,
as it will be of no use for errors returned by zerrors.

More details and examples can be found in the various examples, tests and benchmarks in the source code.

### Migrating From...

- [`github.com/pkg/errors`](https://github.com/pkg/errors): 
change `errors.Wrap(err, "...")`
to `zerrors.SWrap(err, "...")`.
- [`golang.org/x/xerrors`](https://github.com/golang/xerrors) or 
[`fmt.Errorf`](https://golang.org/pkg/fmt/#Errorf) with `%w`:
change `[xerrors/fmt].Errorf("...: %w", ..., err)` to
`zerrors.SWrap(err, "...")` or `zerrors.SWrap(err, fmt.Sprintf("...", ...))`.
- Custom `MyError{..., err: err}` for `type MyError` with `Unwrap() error` and
`err error` parameter:
remove `err` and `Unwrap`, change to `zerrors.Wrap(err, MyError{...})`

## General guidance

Package zerrors works best if no other errors implement `interface{ Unwrap() error }`,
as these will not support frames, error formatting or the performance optimisations inside zerrors.
Note libraries may use zerrors without those importing it knowing about it or changing their handling of errors,
the callers will simply not make use of any of the additional functionality this package offers.
Library callers - the `package main` - are best suited to make use of these,
and through `zerrors/zmain` and custom formatters decide how all errors are serialised,
including errors coming from external libraries.

## Benchmarks

See [benchmark/README.md](internal/benchmark/README.md).
There are performance comparisons to all current mayor strategies for error wrapping: 
- the standard library's [`errors.New`](https://golang.org/pkg/errors/#New) and
[`fmt.Errorf`](https://golang.org/pkg/fmt/#Errorf) with the `%w` pattern
- [`github.com/pkg/errors`](https://github.com/pkg/errors)
- [`golang.org/x/xerrors`](https://github.com/golang/xerrors)

Benchmarks show zerrors to have a better generally performance than all the above, while being more flexible.

## Support and Future work

This is not an officially supported Google product.

As of Feb 2020, zerrors is new and remains untested in production.
Its implementation is minimal and consists of only a few exported methods.
Please contribute to improve this library by sharing your experience via an issue or pull request.
