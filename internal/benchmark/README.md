# Benchmark results

## frameless

**go test ./frameless/... -test.bench=Bench -test.benchmem | go run ./cmptool > frameless/README.md**

No frame info in the error message. The error message is identical in all errors. See [results](frameless/README.md).

## withframe

**go test ./withframe/... -test.bench=Bench -test.benchmem | go run ./cmptool > withframe/README.md**

With frame info in the error message. The error message is different in the various errors. See [results](withframe/README.md).