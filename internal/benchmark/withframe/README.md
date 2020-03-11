# ns/op

 Benchmark (ns/op) | depth=1 | 2 | 3 | 5 | 10 | 20
--- | --- | --- | --- | --- | --- | ---
BenchmarkPkgErrors_WithStackAndMessage | 8077.00 | 9561.00 | 11350.00 | 14211.00 | 22386.00 | 40420.00
BenchmarkXerrors_Errorf | 5546.00 | 8623.00 | 10993.00 | 16817.00 | 30895.00 | 60334.00
BenchmarkZerrors_SWrapLikePkgErrors | 2637.00 | 4160.00 | 5197.00 | 7893.00 | 12982.00 | 24145.00
BenchmarkZerrors_SWrapLikeXerrors | 2109.00 | 3378.00 | 4178.00 | 6365.00 | 11580.00 | 22182.00
BenchmarkZerrors_SWrapWithFrame | 1522.00 | 2369.00 | 3073.00 | 4586.00 | 8337.00 | 15848.00
BenchmarkZerrors_WrapLikePkgErrors | 2944.00 | 4750.00 | 6039.00 | 9208.00 | 15778.00 | 29527.00
BenchmarkZerrors_WrapLikeXerrors | 2587.00 | 3982.00 | 5004.00 | 7887.00 | 14755.00 | 27191.00
BenchmarkZerrors_WrapWithFrame | 1912.00 | 3030.00 | 3932.00 | 5863.00 | 10891.00 | 20804.00

 # B/op

 Benchmark (B/op) | depth=1 | 2 | 3 | 5 | 10 | 20
--- | --- | --- | --- | --- | --- | ---
BenchmarkPkgErrors_WithStackAndMessage | 1352 | 1528 | 1832 | 2440 | 3961 | 6751
BenchmarkXerrors_Errorf | 2104 | 3705 | 4345 | 7546 | 15167 | 30034
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
