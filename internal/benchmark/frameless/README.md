# ns/op

 Benchmark (ns/op) | depth=1 | 2 | 3 | 5 | 10 | 20
--- | --- | --- | --- | --- | --- | ---
BenchmarkErrors_New | 125.00 | 237.00 | 311.00 | 513.00 | 1193.00 | 2790.00
BenchmarkFmt_Errorf | 313.00 | 614.00 | 890.00 | 1568.00 | 3023.00 | 6241.00
BenchmarkPkgErrors_WithMessage | 179.00 | 289.00 | 434.00 | 681.00 | 1357.00 | 3766.00
BenchmarkXerrors_Errorf | 2247.00 | 3647.00 | 4870.00 | 7047.00 | 12066.00 | 22683.00
BenchmarkZerrors_SWrap | 212.00 | 331.00 | 380.00 | 544.00 | 977.00 | 1887.00
BenchmarkZerrors_Wrap | 254.00 | 357.00 | 462.00 | 677.00 | 1216.00 | 2319.00

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
