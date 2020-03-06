# ns/op

 Benchmark (ns/op) | depth=1 | 2 | 3 | 5 | 10 | 20
--- | --- | --- | --- | --- | --- | ---
BenchmarkPkgErrors_WithStackAndMessage | 7922.00 | 9419.00 | 11090.00 | 14366.00 | 22361.00 | 39658.00
BenchmarkXerrors_Errorf | 5466.00 | 8499.00 | 11082.00 | 16924.00 | 31111.00 | 59512.00
BenchmarkZerrors_SWrapWithFrame | 1500.00 | 2302.00 | 2993.00 | 4469.00 | 8135.00 | 15747.00
BenchmarkZerrors_WrapWithFrame | 1857.00 | 2825.00 | 3814.00 | 5738.00 | 10723.00 | 20241.00

 # B/op

 Benchmark (B/op) | depth=1 | 2 | 3 | 5 | 10 | 20
--- | --- | --- | --- | --- | --- | ---
BenchmarkPkgErrors_WithStackAndMessage | 1352 | 1528 | 1832 | 2441 | 3962 | 6753
BenchmarkXerrors_Errorf | 2104 | 3705 | 4345 | 7546 | 15168 | 30034
BenchmarkZerrors_SWrapWithFrame | 576 | 864 | 1152 | 1728 | 3168 | 6081
BenchmarkZerrors_WrapWithFrame | 576 | 864 | 1152 | 1728 | 3168 | 6081

 # allocs/op

 Benchmark (allocs/op) | depth=1 | 2 | 3 | 5 | 10 | 20
--- | --- | --- | --- | --- | --- | ---
BenchmarkPkgErrors_WithStackAndMessage |  17 |  20 |  23 |  29 |  44 |  74
BenchmarkXerrors_Errorf |  27 |  40 |  52 |  77 |  138 |  259
BenchmarkZerrors_SWrapWithFrame |  5 |  7 |  9 |  13 |  23 |  43
BenchmarkZerrors_WrapWithFrame |  7 |  10 |  13 |  19 |  34 |  64
