# ns/op

 Benchmark (ns/op) | depth=1 | 2 | 3 | 5 | 10 | 20
--- | --- | --- | --- | --- | --- | ---
BenchmarkErrors_New | 123.00 | 212.00 | 306.00 | 509.00 | 1073.00 | 2503.00
BenchmarkFmt_Errorf | 312.00 | 587.00 | 881.00 | 1462.00 | 2997.00 | 6291.00
BenchmarkPkgErrors_WithMessage | 144.00 | 258.00 | 362.00 | 610.00 | 1291.00 | 2852.00
BenchmarkXerrors_Errorf | 2201.00 | 3307.00 | 4419.00 | 6452.00 | 11917.00 | 22809.00
BenchmarkZerrors_SWrap | 207.00 | 290.00 | 363.00 | 532.00 | 942.00 | 1746.00
BenchmarkZerrors_Wrap | 251.00 | 355.00 | 448.00 | 652.00 | 1168.00 | 2270.00

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
