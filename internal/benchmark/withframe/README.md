# ns/op

 Benchmark (ns/op) | depth=1 | 2 | 3 | 5 | 10 | 20
--- | --- | --- | --- | --- | --- | ---
BenchmarkPkgErrors_WithStackAndMessage | 7949.00 | 10647.00 | 11122.00 | 14278.00 | 22310.00 | 43539.00
BenchmarkXerrors_Errorf | 6051.00 | 8772.00 | 11191.00 | 17101.00 | 31451.00 | 60593.00
BenchmarkZerrors_SWrapWithFrame | 1505.00 | 2317.00 | 3088.00 | 5454.00 | 11270.00 | 15477.00
BenchmarkZerrors_WrapWithFrame | 1857.00 | 2852.00 | 3902.00 | 5815.00 | 10562.00 | 20905.00

 # B/op

 Benchmark (B/op) | depth=1 | 2 | 3 | 5 | 10 | 20
--- | --- | --- | --- | --- | --- | ---
BenchmarkPkgErrors_WithStackAndMessage | 1352 | 1528 | 1832 | 2440 | 3962 | 6753
BenchmarkXerrors_Errorf | 2104 | 3705 | 4345 | 7546 | 15167 | 30028
BenchmarkZerrors_SWrapWithFrame | 576 | 864 | 1152 | 1728 | 3168 | 6081
BenchmarkZerrors_WrapWithFrame | 576 | 864 | 1152 | 1728 | 3168 | 6081

 # allocs/op

 Benchmark (allocs/op) | depth=1 | 2 | 3 | 5 | 10 | 20
--- | --- | --- | --- | --- | --- | ---
BenchmarkPkgErrors_WithStackAndMessage |  17 |  20 |  23 |  29 |  44 |  74
BenchmarkXerrors_Errorf |  27 |  40 |  52 |  77 |  138 |  259
BenchmarkZerrors_SWrapWithFrame |  5 |  7 |  9 |  13 |  23 |  43
BenchmarkZerrors_WrapWithFrame |  7 |  10 |  13 |  19 |  34 |  64
