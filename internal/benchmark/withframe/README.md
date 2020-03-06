# ns/op

 Benchmark (ns/op) | depth=1 | 2 | 3 | 5 | 10 | 20
--- | --- | --- | --- | --- | --- | ---
BenchmarkPkgErrors_WithStackAndMessage | 7949.00 | 10647.00 | 11122.00 | 14278.00 | 22310.00 | 43539.00
BenchmarkXerrors_Errorf | 6051.00 | 8772.00 | 11191.00 | 17101.00 | 31451.00 | 60593.00
BenchmarkZerrors_SWrapLikeXerrors | 2028.00 | 3282.00 | 4085.00 | 6148.00 | 11742.00 | 21835.00
BenchmarkZerrors_SWrapWithFrame | 1505.00 | 2317.00 | 3088.00 | 5454.00 | 11270.00 | 15477.00
BenchmarkZerrors_WrapLikeXerrors | 2454.00 | 3945.00 | 5059.00 | 7497.00 | 14233.00 | 26351.00
BenchmarkZerrors_WrapWithFrame | 1857.00 | 2852.00 | 3902.00 | 5815.00 | 10562.00 | 20905.00

 # B/op

 Benchmark (B/op) | depth=1 | 2 | 3 | 5 | 10 | 20
--- | --- | --- | --- | --- | --- | ---
BenchmarkPkgErrors_WithStackAndMessage | 1352 | 1528 | 1832 | 2440 | 3962 | 6753
BenchmarkXerrors_Errorf | 2104 | 3705 | 4345 | 7546 | 15167 | 30028
BenchmarkZerrors_SWrapLikeXerrors | 1696 | 3040 | 3552 | 6240 | 12384 | 25312
BenchmarkZerrors_SWrapWithFrame | 576 | 864 | 1152 | 1728 | 3168 | 6081
BenchmarkZerrors_WrapLikeXerrors | 1648 | 3024 | 3536 | 6224 | 12368 | 24016
BenchmarkZerrors_WrapWithFrame | 576 | 864 | 1152 | 1728 | 3168 | 6081

 # allocs/op

 Benchmark (allocs/op) | depth=1 | 2 | 3 | 5 | 10 | 20
--- | --- | --- | --- | --- | --- | ---
BenchmarkPkgErrors_WithStackAndMessage |  17 |  20 |  23 |  29 |  44 |  74
BenchmarkXerrors_Errorf |  27 |  40 |  52 |  77 |  138 |  259
BenchmarkZerrors_SWrapLikeXerrors |  8 |  11 |  13 |  18 |  29 |  50
BenchmarkZerrors_SWrapWithFrame |  5 |  7 |  9 |  13 |  23 |  43
BenchmarkZerrors_WrapLikeXerrors |  10 |  14 |  17 |  24 |  40 |  71
BenchmarkZerrors_WrapWithFrame |  7 |  10 |  13 |  19 |  34 |  64
