# ns/op

 Benchmark (ns/op) | depth=1 | 2 | 3 | 5 | 10 | 20
--- | --- | --- | --- | --- | --- | ---
BenchmarkErrors_New | 123.00 | 212.00 | 309.00 | 511.00 | 1049.00 | 2437.00
BenchmarkFmt_Errorf | 307.00 | 583.00 | 882.00 | 1455.00 | 2984.00 | 6354.00
BenchmarkPkgErrors_WithMessage | 143.00 | 247.00 | 354.00 | 607.00 | 1309.00 | 2853.00
BenchmarkXerrors_Errorf | 2213.00 | 3324.00 | 4363.00 | 6455.00 | 12016.00 | 22459.00
BenchmarkZerrors_SWrap | 201.00 | 290.00 | 361.00 | 511.00 | 926.00 | 1785.00
BenchmarkZerrors_Wrap | 253.00 | 360.00 | 451.00 | 663.00 | 1182.00 | 2282.00

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
