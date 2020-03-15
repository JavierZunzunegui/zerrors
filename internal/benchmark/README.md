# Benchmarks

Compares the performance of the two error wrapping methods in zerrors (`Wrap` and `SWrap`) to other common error
wrapping methods:
- the standard library's `errors.New`.
- the standard library's `fmt.Errorf` with the `%w` pattern.
- `github.com/pkg/errors`.
- `golang.org/x/xerrors`.

## frameless

**go test ./frameless/... -test.bench=Bench -test.benchmem | go run ./cmptool > frameless/README.md**

No frame info in the error message. The error message is identical in all errors. See [results](frameless/README.md).

## withframe

**go test ./withframe/... -test.bench=Bench -test.benchmem | go run ./cmptool > withframe/README.md**

With frame info in the error message. The error message is different in the various errors. See 
[results](withframe/README.md). Note however `golang.org/x/xerrors` is identically replicated in the `...LikeXerrors` 
benchmarks and `github.com/pkg/errors` is partially replicated in the `...LikePkgErrors` benchmarks.