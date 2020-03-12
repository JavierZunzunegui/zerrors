# ns/op

 Benchmark (ns/op) | depth=1 | 2 | 3 | 5 | 10 | 20
--- | --- | --- | --- | --- | --- | ---
BenchmarkPkgErrors_WithStackAndMessage | 8004.00 | 9596.00 | 11233.00 | 14328.00 | 22558.00 | 40283.00
BenchmarkXerrors_Errorf | 5540.00 | 8631.00 | 11071.00 | 17012.00 | 31966.00 | 59144.00
BenchmarkZerrors_SWrapLikePkgErrors | 2547.00 | 4122.00 | 5115.00 | 7927.00 | 13056.00 | 24750.00
BenchmarkZerrors_SWrapLikeXerrors | 2127.00 | 3454.00 | 4254.00 | 6366.00 | 11459.00 | 22402.00
BenchmarkZerrors_SWrapWithFrame | 1532.00 | 2376.00 | 3139.00 | 4504.00 | 8312.00 | 15980.00
BenchmarkZerrors_WrapLikePkgErrors | 2918.00 | 4766.00 | 5975.00 | 9297.00 | 15456.00 | 29629.00
BenchmarkZerrors_WrapLikeXerrors | 2565.00 | 4019.00 | 5166.00 | 7838.00 | 14540.00 | 27286.00
BenchmarkZerrors_WrapWithFrame | 1927.00 | 2900.00 | 4044.00 | 5879.00 | 10856.00 | 20732.00

 # B/op

 Benchmark (B/op) | depth=1 | 2 | 3 | 5 | 10 | 20
--- | --- | --- | --- | --- | --- | ---
BenchmarkPkgErrors_WithStackAndMessage | 1352 | 1528 | 1832 | 2441 | 3962 | 6748
BenchmarkXerrors_Errorf | 2104 | 3705 | 4345 | 7546 | 15167 | 30032
BenchmarkZerrors_SWrapLikePkgErrors | 2022 | 4041 | 4556 | 8531 | 12387 | 25603
BenchmarkZerrors_SWrapLikeXerrors | 1696 | 3040 | 3552 | 6240 | 12384 | 25312
BenchmarkZerrors_SWrapWithFrame | 576 | 864 | 1152 | 1744 | 3200 | 6081
BenchmarkZerrors_WrapLikePkgErrors | 2006 | 4025 | 4540 | 8515 | 12371 | 25587
BenchmarkZerrors_WrapLikeXerrors | 1648 | 3024 | 3536 | 6224 | 12368 | 24016
BenchmarkZerrors_WrapWithFrame | 576 | 864 | 1152 | 1744 | 3200 | 6081

 # allocs/op

 Benchmark (allocs/op) | depth=1 | 2 | 3 | 5 | 10 | 20
--- | --- | --- | --- | --- | --- | ---
BenchmarkPkgErrors_WithStackAndMessage |  17 |  20 |  23 |  29 |  44 |  74
BenchmarkXerrors_Errorf |  27 |  40 |  52 |  77 |  138 |  259
BenchmarkZerrors_SWrapLikePkgErrors |  14 |  20 |  23 |  32 |  49 |  82
BenchmarkZerrors_SWrapLikeXerrors |  8 |  11 |  13 |  18 |  29 |  50
BenchmarkZerrors_SWrapWithFrame |  5 |  7 |  9 |  13 |  23 |  43
BenchmarkZerrors_WrapLikePkgErrors |  16 |  23 |  27 |  38 |  60 |  103
BenchmarkZerrors_WrapLikeXerrors |  10 |  14 |  17 |  24 |  40 |  71
BenchmarkZerrors_WrapWithFrame |  7 |  10 |  13 |  19 |  34 |  64
