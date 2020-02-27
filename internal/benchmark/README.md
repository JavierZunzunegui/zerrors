# Benchmark results

**go test ./ -test.bench=Bench --test.benchmem**

## frameless

Without frame info, all errors produce the same string.

```
goos: linux
goarch: amd64
pkg: github.com/JavierZunzunegui/zerrors/internal/benchmark/frameless
```

### BenchmarkZerrors_Wrap
```
BenchmarkZerrors_Wrap/depth_1-4                 	 4825610	       248 ns/op	     112 B/op	       5 allocs/op
BenchmarkZerrors_Wrap/depth_2-4                 	 3373310	       356 ns/op	     176 B/op	       7 allocs/op
BenchmarkZerrors_Wrap/depth_3-4                 	 2696180	       451 ns/op	     224 B/op	       9 allocs/op
BenchmarkZerrors_Wrap/depth_5-4                 	 1844936	       660 ns/op	     352 B/op	      13 allocs/op
BenchmarkZerrors_Wrap/depth_10-4                   	 1001796	      1184 ns/op	     624 B/op	      23 allocs/op
BenchmarkZerrors_Wrap/depth_20-4                   	  450079	      2336 ns/op	    1200 B/op	      43 allocs/op
```

### BenchmarkZerrors_SWrap
```
BenchmarkZerrors_SWrap/depth_1-4                   	 5830510	       207 ns/op	     112 B/op	       3 allocs/op
BenchmarkZerrors_SWrap/depth_2-4                   	 4171662	       290 ns/op	     176 B/op	       4 allocs/op
BenchmarkZerrors_SWrap/depth_3-4                   	 3316587	       367 ns/op	     224 B/op	       5 allocs/op
BenchmarkZerrors_SWrap/depth_5-4                   	 2275506	       516 ns/op	     352 B/op	       7 allocs/op
BenchmarkZerrors_SWrap/depth_10-4                 	 1297681	       932 ns/op	     624 B/op	      12 allocs/op
BenchmarkZerrors_SWrap/depth_20-4                 	  645165	      1793 ns/op	    1200 B/op	      22 allocs/op
```

### BenchmarkFmt_Errorf
```
BenchmarkFmt_Errorf/depth_1-4                     	 3905769	       306 ns/op	      64 B/op	       3 allocs/op
BenchmarkFmt_Errorf/depth_2-4                     	 2058244	       579 ns/op	     128 B/op	       5 allocs/op
BenchmarkFmt_Errorf/depth_3-4                     	 1375432	       878 ns/op	     192 B/op	       7 allocs/op
BenchmarkFmt_Errorf/depth_5-4                     	  788132	      1446 ns/op	     368 B/op	      11 allocs/op
BenchmarkFmt_Errorf/depth_10-4                   	  397296	      2973 ns/op	     944 B/op	      21 allocs/op
BenchmarkFmt_Errorf/depth_20-4                   	  184642	      6232 ns/op	    2768 B/op	      41 allocs/op
```

### BenchmarkErrors_New
```
BenchmarkErrors_New/depth_1-4                     	 9944774	       124 ns/op	      48 B/op	       3 allocs/op
BenchmarkErrors_New/depth_2-4                     	 5698030	       212 ns/op	      96 B/op	       5 allocs/op
BenchmarkErrors_New/depth_3-4                     	 3808365	       308 ns/op	     144 B/op	       7 allocs/op
BenchmarkErrors_New/depth_5-4                     	 2289410	       513 ns/op	     288 B/op	      11 allocs/op
BenchmarkErrors_New/depth_10-4                   	 1000000	      1076 ns/op	     784 B/op	      21 allocs/op
BenchmarkErrors_New/depth_20-4                   	  517017	      2503 ns/op	    2448 B/op	      41 allocs/op
```

### BenchmarkPkgErrors_WithMessage
```
BenchmarkPkgErrors_WithMessage/depth_1-4         	 8012893	       145 ns/op	      64 B/op	       3 allocs/op
BenchmarkPkgErrors_WithMessage/depth_2-4         	 4857700	       254 ns/op	     128 B/op	       5 allocs/op
BenchmarkPkgErrors_WithMessage/depth_3-4         	 3350054	       358 ns/op	     192 B/op	       7 allocs/op
BenchmarkPkgErrors_WithMessage/depth_5-4         	 1967660	       596 ns/op	     368 B/op	      11 allocs/op
BenchmarkPkgErrors_WithMessage/depth_10-4        	  968436	      1262 ns/op	     944 B/op	      21 allocs/op
BenchmarkPkgErrors_WithMessage/depth_20-4        	  428677	      2835 ns/op	    2768 B/op	      41 allocs/op
```

### BenchmarkXerrors_Errorf
```
BenchmarkXerrors_Errorf/depth_1-4                	  545202	      2180 ns/op	     344 B/op	      11 allocs/op
BenchmarkXerrors_Errorf/depth_2-4                	  359118	      3271 ns/op	     480 B/op	      16 allocs/op
BenchmarkXerrors_Errorf/depth_3-4                	  264258	      4317 ns/op	     600 B/op	      21 allocs/op
BenchmarkXerrors_Errorf/depth_5-4                	  185964	      6399 ns/op	     872 B/op	      31 allocs/op
BenchmarkXerrors_Errorf/depth_10-4               	  102651	     11717 ns/op	    1648 B/op	      57 allocs/op
BenchmarkXerrors_Errorf/depth_20-4               	   53280	     22251 ns/op	    3232 B/op	     108 allocs/op
```

## withframe

Without frame info, the error string format is different for different errors.

```
goos: linux
goarch: amd64
pkg: github.com/JavierZunzunegui/zerrors/internal/benchmark/withframe
```

### BenchmarkZerrors_WrapWithFrame
```
BenchmarkZerrors_WrapWithFrame/depth_1-4                   	  560906	      1850 ns/op	     576 B/op	       7 allocs/op
BenchmarkZerrors_WrapWithFrame/depth_2-4                   	  382780	      2850 ns/op	     864 B/op	      10 allocs/op
BenchmarkZerrors_WrapWithFrame/depth_3-4                   	  282442	      3813 ns/op	    1152 B/op	      13 allocs/op
BenchmarkZerrors_WrapWithFrame/depth_5-4                   	  207756	      5626 ns/op	    1728 B/op	      19 allocs/op
BenchmarkZerrors_WrapWithFrame/depth_10-4                 	  113715	     10601 ns/op	    3168 B/op	      34 allocs/op
BenchmarkZerrors_WrapWithFrame/depth_20-4                 	   57622	     20084 ns/op	    6081 B/op	      64 allocs/op
```

### BenchmarkZerrors_SWrapWithFrame
```
BenchmarkZerrors_SWrapWithFrame/depth_1-4                 	  711238	      1505 ns/op	     576 B/op	       5 allocs/op
BenchmarkZerrors_SWrapWithFrame/depth_2-4                 	  486979	      2262 ns/op	     864 B/op	       7 allocs/op
BenchmarkZerrors_SWrapWithFrame/depth_3-4                 	  388880	      2999 ns/op	    1152 B/op	       9 allocs/op
BenchmarkZerrors_SWrapWithFrame/depth_5-4                 	  251052	      4429 ns/op	    1728 B/op	      13 allocs/op
BenchmarkZerrors_SWrapWithFrame/depth_10-4               	  141558	      8272 ns/op	    3168 B/op	      23 allocs/op
BenchmarkZerrors_SWrapWithFrame/depth_20-4               	   73591	     15649 ns/op	    6081 B/op	      43 allocs/op
```

### BenchmarkPkgErrors_WithStackAndMessage
```
BenchmarkPkgErrors_WithStackAndMessage/depth_1-4         	  139720	      7938 ns/op	    1352 B/op	      17 allocs/op
BenchmarkPkgErrors_WithStackAndMessage/depth_2-4         	  124333	      9278 ns/op	    1528 B/op	      20 allocs/op
BenchmarkPkgErrors_WithStackAndMessage/depth_3-4         	  105999	     11089 ns/op	    1832 B/op	      23 allocs/op
BenchmarkPkgErrors_WithStackAndMessage/depth_5-4         	   83389	     14121 ns/op	    2440 B/op	      29 allocs/op
BenchmarkPkgErrors_WithStackAndMessage/depth_10-4        	   54831	     21824 ns/op	    3963 B/op	      44 allocs/op
BenchmarkPkgErrors_WithStackAndMessage/depth_20-4        	   30078	     39509 ns/op	    6753 B/op	      74 allocs/op
```

### BenchmarkXerrors_Errorf
```
BenchmarkXerrors_Errorf/depth_1-4                        	  217130	      5417 ns/op	    2104 B/op	      27 allocs/op
BenchmarkXerrors_Errorf/depth_2-4                        	  138242	      8814 ns/op	    3705 B/op	      40 allocs/op
BenchmarkXerrors_Errorf/depth_3-4                        	  109992	     11046 ns/op	    4345 B/op	      52 allocs/op
BenchmarkXerrors_Errorf/depth_5-4                        	   69910	     16868 ns/op	    7546 B/op	      77 allocs/op
BenchmarkXerrors_Errorf/depth_10-4                       	   38083	     30735 ns/op	   15167 B/op	     138 allocs/op
BenchmarkXerrors_Errorf/depth_20-4                       	   19984	     59274 ns/op	   30029 B/op	     259 allocs/op
```