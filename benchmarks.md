# Benchmark results

**go test ./ -test.bench=Bench --test.benchmem**

```
goos: linux
goarch: amd64
pkg: github.com/JavierZunzunegui/zerrors
```

## Internal benchmarks

Comparisons relevant within zerrors only.

```
# zerrors.Wrap
BenchmarkWrappingError_Error/Wrap/depth-1-4 	        13298210      85.8 ns/op      16 B/op    1 allocs/op
BenchmarkWrappingError_Error/Wrap/depth-2-4 	        10205412       114 ns/op      32 B/op    1 allocs/op
BenchmarkWrappingError_Error/Wrap/depth-3-4 	         8872419       131 ns/op      32 B/op    1 allocs/op
BenchmarkWrappingError_Error/Wrap/depth-5-4 	         6203349       184 ns/op      64 B/op    1 allocs/op
BenchmarkWrappingError_Error/Wrap/depth-10-4         	 3828380       310 ns/op      96 B/op    1 allocs/op
BenchmarkWrappingError_Error/Wrap/depth-20-4         	 2165709       550 ns/op     192 B/op    1 allocs/op

# zerrors.Wrap, with Frame
BenchmarkWrappingError_Error/Wrap_WithFrame/depth-1-4  	 3954945       301 ns/op      64 B/op    1 allocs/op
BenchmarkWrappingError_Error/Wrap_WithFrame/depth-2-4  	 2773854       425 ns/op     112 B/op    1 allocs/op
BenchmarkWrappingError_Error/Wrap_WithFrame/depth-3-4  	 2097567       574 ns/op     144 B/op    1 allocs/op
BenchmarkWrappingError_Error/Wrap_WithFrame/depth-5-4  	 1442390       826 ns/op     208 B/op    1 allocs/op
BenchmarkWrappingError_Error/Wrap_WithFrame/depth-10-4 	  713192      1492 ns/op     384 B/op    1 allocs/op
BenchmarkWrappingError_Error/Wrap_WithFrame/depth-20-4 	  412788      2740 ns/op     768 B/op    1 allocs/op

# zerrors.SWrap
BenchmarkWrappingError_Error/SWrap/depth-1-4           	13501796      84.6 ns/op      16 B/op    1 allocs/op
BenchmarkWrappingError_Error/SWrap/depth-2-4           	10289352       111 ns/op      32 B/op    1 allocs/op
BenchmarkWrappingError_Error/SWrap/depth-3-4           	 8860008       130 ns/op      32 B/op    1 allocs/op
BenchmarkWrappingError_Error/SWrap/depth-5-4           	 6490771       185 ns/op      64 B/op    1 allocs/op
BenchmarkWrappingError_Error/SWrap/depth-10-4          	 3861034       306 ns/op      96 B/op    1 allocs/op
BenchmarkWrappingError_Error/SWrap/depth-20-4          	 2205504       556 ns/op     192 B/op    1 allocs/op

# zerrors.SWrap, with Frame
BenchmarkWrappingError_Error/SWrap_WithFrame/depth-1-4 	 4068465       292 ns/op      64 B/op    1 allocs/op
BenchmarkWrappingError_Error/SWrap_WithFrame/depth-2-4 	 2827851       420 ns/op     112 B/op    1 allocs/op
BenchmarkWrappingError_Error/SWrap_WithFrame/depth-3-4 	 2117974       561 ns/op     144 B/op    1 allocs/op
BenchmarkWrappingError_Error/SWrap_WithFrame/depth-5-4 	 1484380       803 ns/op     208 B/op    1 allocs/op
BenchmarkWrappingError_Error/SWrap_WithFrame/depth-10-4	  814632      1468 ns/op     384 B/op    1 allocs/op
BenchmarkWrappingError_Error/SWrap_WithFrame/depth-20-4	  398394      2753 ns/op     768 B/op    1 allocs/op
```

## External benchmarks

Comparisons to external error frameworks.
All benchmarks produce the same error message at each depth level.
Except for frames, which carry additional information.

```
# zerrors.Wrap
BenchmarkCreateAndError/Wrap/depth-1-4                 	 4844085       246 ns/op     112 B/op    5 allocs/op
BenchmarkCreateAndError/Wrap/depth-2-4                 	 3426805       346 ns/op     176 B/op    7 allocs/op
BenchmarkCreateAndError/Wrap/depth-3-4                 	 2607721       455 ns/op     224 B/op    9 allocs/op
BenchmarkCreateAndError/Wrap/depth-5-4                 	 1874023       643 ns/op     352 B/op   13 allocs/op
BenchmarkCreateAndError/Wrap/depth-10-4                	 1000000      1164 ns/op     624 B/op   23 allocs/op
BenchmarkCreateAndError/Wrap/depth-20-4                	  489519      2165 ns/op    1200 B/op   43 allocs/op

# zerrors.Wrap, with Frame
BenchmarkCreateAndError/Wrap_WithFrame/depth-1-4       	  988268      1215 ns/op     160 B/op    5 allocs/op
BenchmarkCreateAndError/Wrap_WithFrame/depth-2-4       	  565093      1826 ns/op     256 B/op    7 allocs/op
BenchmarkCreateAndError/Wrap_WithFrame/depth-3-4       	  476220      2418 ns/op     336 B/op    9 allocs/op
BenchmarkCreateAndError/Wrap_WithFrame/depth-5-4       	  320095      3554 ns/op     496 B/op   13 allocs/op
BenchmarkCreateAndError/Wrap_WithFrame/depth-10-4      	  180004      6413 ns/op     912 B/op   23 allocs/op
BenchmarkCreateAndError/Wrap_WithFrame/depth-20-4      	   97352     12166 ns/op    1776 B/op   43 allocs/op

# zerrors.SWrap
BenchmarkCreateAndError/SWrap/depth-1-4                	 6000632       198 ns/op     112 B/op    3 allocs/op
BenchmarkCreateAndError/SWrap/depth-2-4                	 4334427       281 ns/op     176 B/op    4 allocs/op
BenchmarkCreateAndError/SWrap/depth-3-4                	 3434900       347 ns/op     224 B/op    5 allocs/op
BenchmarkCreateAndError/SWrap/depth-5-4                	 2336950       510 ns/op     352 B/op    7 allocs/op
BenchmarkCreateAndError/SWrap/depth-10-4               	 1297386       916 ns/op     624 B/op   12 allocs/op
BenchmarkCreateAndError/SWrap/depth-20-4               	  662431      1695 ns/op    1200 B/op   22 allocs/op

# zerrors.SWrap, with Frame
BenchmarkCreateAndError/SWrap_WithFrame/depth-1-4      	 1000000      1015 ns/op     160 B/op    3 allocs/op
BenchmarkCreateAndError/SWrap_WithFrame/depth-2-4      	  671530      1534 ns/op     256 B/op    4 allocs/op
BenchmarkCreateAndError/SWrap_WithFrame/depth-3-4      	  557664      2042 ns/op     336 B/op    5 allocs/op
BenchmarkCreateAndError/SWrap_WithFrame/depth-5-4      	  340326      2963 ns/op     496 B/op    7 allocs/op
BenchmarkCreateAndError/SWrap_WithFrame/depth-10-4     	  208437      5323 ns/op     912 B/op   12 allocs/op
BenchmarkCreateAndError/SWrap_WithFrame/depth-20-4     	  115428     10040 ns/op    1776 B/op   22 allocs/op

# fmt.Errorf
BenchmarkCreateAndError/fmt_Errorf/depth-1-4           	 3911899       302 ns/op      64 B/op    3 allocs/op
BenchmarkCreateAndError/fmt_Errorf/depth-2-4           	 2109055       566 ns/op     128 B/op    5 allocs/op
BenchmarkCreateAndError/fmt_Errorf/depth-3-4           	 1426414       837 ns/op     192 B/op    7 allocs/op
BenchmarkCreateAndError/fmt_Errorf/depth-5-4           	  797191      1396 ns/op     368 B/op   11 allocs/op
BenchmarkCreateAndError/fmt_Errorf/depth-10-4          	  405751      2826 ns/op     944 B/op   21 allocs/op
BenchmarkCreateAndError/fmt_Errorf/depth-20-4          	  195168      5940 ns/op    2768 B/op   41 allocs/op

# errors.New
BenchmarkCreateAndError/errors_New/depth-1-4           	 9423854       116 ns/op      48 B/op    3 allocs/op
BenchmarkCreateAndError/errors_New/depth-2-4           	 5731902       204 ns/op      96 B/op    5 allocs/op
BenchmarkCreateAndError/errors_New/depth-3-4           	 4211803       286 ns/op     144 B/op    7 allocs/op
BenchmarkCreateAndError/errors_New/depth-5-4           	 2313610       485 ns/op     288 B/op   11 allocs/op
BenchmarkCreateAndError/errors_New/depth-10-4          	 1000000      1042 ns/op     784 B/op   21 allocs/op
BenchmarkCreateAndError/errors_New/depth-20-4          	  514698      2240 ns/op    2448 B/op   41 allocs/op
```
