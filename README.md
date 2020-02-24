# zerrors 
[![GoDoc](https://godoc.org/github.com/JavierZunzunegui/zerrors?status.svg)](http://godoc.org/github.com/JavierZunzunegui/zerrors) 
[![Report card](https://goreportcard.com/badge/github.com/JavierZunzunegui/zerrors)](https://goreportcard.com/report/github.com/JavierZunzunegui/zerrors)

Package zerrors provides additional functionality on top of the go 1.13 error wrapping features,
particularly frame information and flexibility over the wrapped error formatting, without sacrificing on performance.

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

Global errors should also be replaced to use `zerrors.New` and `zerrors.SNew`.

The resulting error's `Error() string` method is of format `last message: ...: first wrap message: base message`,
which is produced in a efficient manner.
Alternative serialisation of the error is also possible, such that one can have a custom `MyFormat(error) string`
to produce `last message - ... - first wrap message - base message`, or any other custom desirable format.
Additionally, if `import _ "github.com/JavierZunzunegui/zerrors/zmain"` is included the errors come with frame info
such that `Error() string` becomes `last message (file.go:10): ...: base message (base_file.go:11)`,
where the files and lines indicate where the error was created.
Custom formatting is also supported for frames.

The `errors.Is`, `errors.As` and `errors.Unwrap` methods from go1.13 are supported as expected,
and are intended to remain the primary means to examine the contents of errors.
Two new additional methods, `Value(error) error` and `Frame(error) (runtime.Frame, bool)`
are introduced and can be used in conjunction with `errors.Unwrap` to implement more sophisticated functionality.
Error inspection via type assertion is further discouraged.

More details and examples can be found in the various examples and tests in the source code.

## General guidance

Package zerrors works best if no other errors implement `interface{ Unwrap() error }`,
as these will not support frames, error formatting or the performance optimisations inside zerrors.
Note libraries may use zerrors without those importing it knowing about it or changing their handling of errors,
the callers will simply not make use of any of the additional functionality this package offers.
Library callers - the `package main` - are best suited to make use of these,
and through `zerrors/zmain` and custom formatters decide how all errors are serialised,
including errors coming from external libraries.

## Support and Future work

This is not an officially supported Google product.

As of Feb 2020, zerrors is new and remains untested in production.
Its implementation is minimal and consists of only a few exported methods.
Please contribute to improve this library by sharing your experience via an issue or pull request.
