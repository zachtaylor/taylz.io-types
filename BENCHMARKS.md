# Benchmarks

This file contains benchmark comparisons to the standard library with various payload sizes

## Results

```
goos: linux
goarch: amd64
pkg: taylz.io/types
BenchmarkBuiltinStringBytes-4     	47693184	        24.0 ns/op	       0 B/op	       0 allocs/op
BenchmarkTypesNewStringBytes-4    	313794432	         3.88 ns/op	       0 B/op	       0 allocs/op
BenchmarkFmtSprintSlice-4         	  183924	      6463 ns/op	     384 B/op	      11 allocs/op
BenchmarkJsonMarshalSlice-4       	  250132	      4723 ns/op	     320 B/op	       8 allocs/op
BenchmarkTypesNewStringSlice1-4   	  403802	      2982 ns/op	     144 B/op	       5 allocs/op
BenchmarkTypesNewStringSlice2-4   	  435655	      2797 ns/op	     112 B/op	       4 allocs/op
BenchmarkFmtSprint1-4             	   34844	     32371 ns/op	    1920 B/op	      41 allocs/op
BenchmarkJsonMarshal1-4           	   40701	     28480 ns/op	    2336 B/op	      43 allocs/op
BenchmarkTypesNewStringDict1-4    	  160755	      7159 ns/op	     384 B/op	       1 allocs/op
BenchmarkTypesNewStringMap1-4     	  125554	      9287 ns/op	     384 B/op	       1 allocs/op
BenchmarkFmtSprint2-4             	    1603	    707752 ns/op	   44628 B/op	     992 allocs/op
BenchmarkJsonMarshal2-4           	    1801	    621979 ns/op	   50242 B/op	    1005 allocs/op
BenchmarkTypesNewStringDict2-4    	    5610	    193277 ns/op	    9652 B/op	     170 allocs/op
BenchmarkTypesNewStringMap2-4     	    6058	    194627 ns/op	    9651 B/op	     170 allocs/op
BenchmarkFmtSprint3-4             	    1095	   1052620 ns/op	   66239 B/op	    1508 allocs/op
BenchmarkJsonMarshal3-4           	    1164	    935566 ns/op	   77187 B/op	    1538 allocs/op
BenchmarkTypesNewStringDict3-4    	    3973	    292551 ns/op	   15164 B/op	     253 allocs/op
BenchmarkTypesNewStringMap3-4     	    3980	    292671 ns/op	   15164 B/op	     253 allocs/op
```
