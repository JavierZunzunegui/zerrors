# ns/op

 Benchmark (ns/op) | depth=1 | 2 | 3 | 5 | 10 | 20
--- | --- | --- | --- | --- | --- | ---
BenchmarkErrors_New | 125.00 | 217.00 | 304.00 | 533.00 | 1116.00 | 2430.00
BenchmarkFmt_Errorf | 309.00 | 577.00 | 873.00 | 1443.00 | 2979.00 | 6171.00
BenchmarkPkgErrors_WithMessage | 143.00 | 252.00 | 357.00 | 600.00 | 1284.00 | 2958.00
BenchmarkXerrors_Errorf | 2184.00 | 3291.00 | 4363.00 | 6428.00 | 11775.00 | 22475.00
BenchmarkZerrors_SWrap | 206.00 | 292.00 | 355.00 | 516.00 | 930.00 | 1788.00
BenchmarkZerrors_Wrap | 246.00 | 358.00 | 456.00 | 655.00 | 1189.00 | 2305.00

 # B/op

 Benchmark (B/op) | depth=1 | 2 | 3 | 5 | 10 | 20
--- | --- | --- | --- | --- | --- | ---
BenchmarkErrors_New | 48 | 96 | 144 | 288 | 784 | 2448
BenchmarkFmt_Errorf | 64 | 128 | 192 | 368 | 944 | 2768
BenchmarkPkgErrors_WithMessage | 64 | 128 | 192 | 368 | 944 | 2768
BenchmarkXerrors_Errorf | 344 | 480 | 600 | 872 | 1648 | 3232
BenchmarkZerrors_SWrap | 112 | 176 | 224 | 352 | 624 | 1200
BenchmarkZerrors_Wrap | 112 | 176 | 224 | 352 | 624 | 1200

 # allocs/op

 Benchmark (allocs/op) | depth=1 | 2 | 3 | 5 | 10 | 20
--- | --- | --- | --- | --- | --- | ---
BenchmarkErrors_New |  3 |  5 |  7 |  11 |  21 |  41
BenchmarkFmt_Errorf |  3 |  5 |  7 |  11 |  21 |  41
BenchmarkPkgErrors_WithMessage |  3 |  5 |  7 |  11 |  21 |  41
BenchmarkXerrors_Errorf |  11 |  16 |  21 |  31 |  57 |  108
BenchmarkZerrors_SWrap |  3 |  4 |  5 |  7 |  12 |  22
BenchmarkZerrors_Wrap |  5 |  7 |  9 |  13 |  23 |  43
