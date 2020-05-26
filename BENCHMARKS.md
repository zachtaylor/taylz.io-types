# Benchmarks

This file contains benchmark comparisons to the standard library with various payload sizes

## Results

```
goos: linux
goarch: amd64
pkg: taylz.io/types

# stringify bytes
BenchmarkBuiltinStringBytes-4    	47493156	        24.9 ns/op	       0 B/op	       0 allocs/op
BenchmarkCastStringBytes-4       	307831524	         3.85 ns/op	       0 B/op	       0 allocs/op

# encode slice to string
BenchmarkSTDFmtSprintSlice-4     	  179725	      6446 ns/op	     384 B/op	      11 allocs/op
BenchmarkSTDJsonMarshalSlice-4   	  245917	      4843 ns/op	     320 B/op	       8 allocs/op
BenchmarkTypesStringSlice-4      	  404456	      3029 ns/op	     144 B/op	       5 allocs/op

# encode payload#1 to map/json format
BenchmarkFmtSprintf1-4           	   37407	     33095 ns/op	    1920 B/op	      41 allocs/op
BenchmarkJsonMarshal1-4          	   42236	     29010 ns/op	    2336 B/op	      43 allocs/op
BenchmarkTypesDictString1-4      	  163870	      7260 ns/op	     384 B/op	       1 allocs/op
BenchmarkTypesMapString1-4       	  123717	      9132 ns/op	     384 B/op	       1 allocs/op

# encode payload#2 to map/json format
BenchmarkFmtSprintf2-4           	    1644	    713180 ns/op	   44627 B/op	     992 allocs/op
BenchmarkJsonMarshal2-4          	    1843	    624040 ns/op	   50241 B/op	    1005 allocs/op
BenchmarkTypesDictString2-4      	    6207	    190847 ns/op	    9651 B/op	     170 allocs/op
BenchmarkTypesMapString2-4       	    5827	    192605 ns/op	    9651 B/op	     170 allocs/op

# encode payload#3 to map/json format
BenchmarkFmtSprintf3-4           	    1105	   1063150 ns/op	   66238 B/op	    1508 allocs/op
BenchmarkJsonMarshal3-4          	    1168	    950726 ns/op	   77170 B/op	    1538 allocs/op
BenchmarkTypesDictString3-4      	    4119	    289326 ns/op	   15163 B/op	     253 allocs/op
BenchmarkTypesMapString3-4       	    4166	    291579 ns/op	   15163 B/op	     253 allocs/op
```

