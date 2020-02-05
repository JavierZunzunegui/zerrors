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
BenchmarkWrappingError_Error/Wrap/depth-1-4 	            13254200	      87.6 ns/op	      16 B/op	       1 allocs/op
BenchmarkWrappingError_Error/Wrap/depth-2-4 	            10297276	       111 ns/op	      32 B/op	       1 allocs/op
BenchmarkWrappingError_Error/Wrap/depth-3-4 	             8879667	       133 ns/op	      32 B/op	       1 allocs/op
BenchmarkWrappingError_Error/Wrap/depth-5-4 	             6529983	       181 ns/op	      64 B/op	       1 allocs/op
BenchmarkWrappingError_Error/Wrap/depth-10-4         	     3885102	       304 ns/op	      96 B/op	       1 allocs/op
BenchmarkWrappingError_Error/Wrap/depth-20-4         	     2222133	       539 ns/op	     192 B/op	       1 allocs/op

# zerrors.Wrap, with Frame
BenchmarkWrappingError_Error/Wrap_WithFrame/depth-1-4      	 1378066	       866 ns/op	     480 B/op	       3 allocs/op
BenchmarkWrappingError_Error/Wrap_WithFrame/depth-2-4      	  891384	      1354 ns/op	     736 B/op	       4 allocs/op
BenchmarkWrappingError_Error/Wrap_WithFrame/depth-3-4      	  618862	      1847 ns/op	     976 B/op	       5 allocs/op
BenchmarkWrappingError_Error/Wrap_WithFrame/depth-5-4      	  412141	      2703 ns/op	    1456 B/op	       7 allocs/op
BenchmarkWrappingError_Error/Wrap_WithFrame/depth-10-4     	  232555	      4978 ns/op	    2672 B/op	      12 allocs/op
BenchmarkWrappingError_Error/Wrap_WithFrame/depth-20-4     	  121351	      9485 ns/op	    5136 B/op	      22 allocs/op

# zerrors.SWrap
BenchmarkWrappingError_Error/SWrap/depth-1-4               	13419445	      88.5 ns/op	      16 B/op	       1 allocs/op
BenchmarkWrappingError_Error/SWrap/depth-2-4               	10583445	       112 ns/op	      32 B/op	       1 allocs/op
BenchmarkWrappingError_Error/SWrap/depth-3-4               	 8643290	       131 ns/op	      32 B/op	       1 allocs/op
BenchmarkWrappingError_Error/SWrap/depth-5-4               	 6458634	       180 ns/op	      64 B/op	       1 allocs/op
BenchmarkWrappingError_Error/SWrap/depth-10-4              	 3882927	       305 ns/op	      96 B/op	       1 allocs/op
BenchmarkWrappingError_Error/SWrap/depth-20-4              	 2189319	       538 ns/op	     192 B/op	       1 allocs/op

# zerrors.SWrap, with Frame
BenchmarkWrappingError_Error/SWrap_WithFrame/depth-1-4     	 1683946	       722 ns/op	     480 B/op	       3 allocs/op
BenchmarkWrappingError_Error/SWrap_WithFrame/depth-2-4     	 1000000	      1056 ns/op	     736 B/op	       4 allocs/op
BenchmarkWrappingError_Error/SWrap_WithFrame/depth-3-4     	  706592	      1423 ns/op	     976 B/op	       5 allocs/op
BenchmarkWrappingError_Error/SWrap_WithFrame/depth-5-4     	  537345	      2144 ns/op	    1456 B/op	       7 allocs/op
BenchmarkWrappingError_Error/SWrap_WithFrame/depth-10-4    	  278078	      3873 ns/op	    2672 B/op	      12 allocs/op
BenchmarkWrappingError_Error/SWrap_WithFrame/depth-20-4    	  165957	      7318 ns/op	    5136 B/op	      22 allocs/op
```

## External benchmarks

Comparisons to external error frameworks.
All benchmarks produce the same error message at each depth level.
Except for frames, which carry additional information.

```
# zerrors.Wrap
BenchmarkCreateAndError/Wrap/depth-1-4                  4938819	       238 ns/op	     112 B/op	       5 allocs/op
BenchmarkCreateAndError/Wrap/depth-2-4                  3412208	       342 ns/op	     176 B/op	       7 allocs/op
BenchmarkCreateAndError/Wrap/depth-3-4                  2678944	       444 ns/op	     224 B/op	       9 allocs/op
BenchmarkCreateAndError/Wrap/depth-5-4                  1854499	       654 ns/op	     352 B/op	      13 allocs/op
BenchmarkCreateAndError/Wrap/depth-10-4                 1036364	      1148 ns/op	     624 B/op	      23 allocs/op
BenchmarkCreateAndError/Wrap/depth-20-4                  541490	      2137 ns/op	    1200 B/op	      43 allocs/op

# zerrors.Wrap, with Frame
BenchmarkCreateAndError/Wrap_WithFrame/depth-1-4         651660	      1819 ns/op	     576 B/op	       7 allocs/op
BenchmarkCreateAndError/Wrap_WithFrame/depth-2-4         418525	      2759 ns/op	     880 B/op	      10 allocs/op
BenchmarkCreateAndError/Wrap_WithFrame/depth-3-4         322153	      3717 ns/op	    1168 B/op	      13 allocs/op
BenchmarkCreateAndError/Wrap_WithFrame/depth-5-4         211983	      5538 ns/op	    1744 B/op	      19 allocs/op
BenchmarkCreateAndError/Wrap_WithFrame/depth-10-4        115669	     10208 ns/op	    3200 B/op	      34 allocs/op
BenchmarkCreateAndError/Wrap_WithFrame/depth-20-4         62293	     19184 ns/op	    6145 B/op	      64 allocs/op

# zerrors.SWrap
BenchmarkCreateAndError/SWrap/depth-1-4                 6355956	       194 ns/op	     112 B/op	       3 allocs/op
BenchmarkCreateAndError/SWrap/depth-2-4                 4132869	       273 ns/op	     176 B/op	       4 allocs/op
BenchmarkCreateAndError/SWrap/depth-3-4                 3511677	       345 ns/op	     224 B/op	       5 allocs/op
BenchmarkCreateAndError/SWrap/depth-5-4                 2319687	       525 ns/op	     352 B/op	       7 allocs/op
BenchmarkCreateAndError/SWrap/depth-10-4                1308480	       909 ns/op	     624 B/op	      12 allocs/op
BenchmarkCreateAndError/SWrap/depth-20-4                 690384	      1696 ns/op	    1200 B/op	      22 allocs/op

# zerrors.SWrap, with Frame
BenchmarkCreateAndError/SWrap_WithFrame/depth-1-4        770977	      1450 ns/op	     576 B/op	       5 allocs/op
BenchmarkCreateAndError/SWrap_WithFrame/depth-2-4        523915	      2165 ns/op	     880 B/op	       7 allocs/op
BenchmarkCreateAndError/SWrap_WithFrame/depth-3-4        400105	      2876 ns/op	    1168 B/op	       9 allocs/op
BenchmarkCreateAndError/SWrap_WithFrame/depth-5-4        272775	      4335 ns/op	    1744 B/op	      13 allocs/op
BenchmarkCreateAndError/SWrap_WithFrame/depth-10-4       151970	      7787 ns/op	    3200 B/op	      23 allocs/op
BenchmarkCreateAndError/SWrap_WithFrame/depth-20-4        81220	     14978 ns/op	    6145 B/op	      43 allocs/op

# fmt.Errorf
BenchmarkCreateAndError/fmt_Errorf/depth-1-4            3987376	       302 ns/op	      64 B/op	       3 allocs/op
BenchmarkCreateAndError/fmt_Errorf/depth-2-4            2114754	       566 ns/op	     128 B/op	       5 allocs/op
BenchmarkCreateAndError/fmt_Errorf/depth-3-4            1420866	       843 ns/op	     192 B/op	       7 allocs/op
BenchmarkCreateAndError/fmt_Errorf/depth-5-4             774889	      1402 ns/op	     368 B/op	      11 allocs/op
BenchmarkCreateAndError/fmt_Errorf/depth-10-4            406684	      2837 ns/op	     944 B/op	      21 allocs/op
BenchmarkCreateAndError/fmt_Errorf/depth-20-4            199404	      6033 ns/op	    2768 B/op	      41 allocs/op

# errors.New
BenchmarkCreateAndError/errors_New/depth-1-4            9212372	       118 ns/op	      48 B/op	       3 allocs/op
BenchmarkCreateAndError/errors_New/depth-2-4            5857843	       205 ns/op	      96 B/op	       5 allocs/op
BenchmarkCreateAndError/errors_New/depth-3-4            4079690	       294 ns/op	     144 B/op	       7 allocs/op
BenchmarkCreateAndError/errors_New/depth-5-4            2411808	       504 ns/op	     288 B/op	      11 allocs/op
BenchmarkCreateAndError/errors_New/depth-10-4            984384	      1062 ns/op	     784 B/op	      21 allocs/op
BenchmarkCreateAndError/errors_New/depth-20-4            492328	      2282 ns/op	    2448 B/op	      41 allocs/op
```
