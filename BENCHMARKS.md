# Benchmarks

This file contains benchmark comparisons to the standard library with various payload sizes

## Results

```
goos: linux
goarch: amd64
pkg: taylz.io/types
BenchmarkBuiltinStringBytes-4        	43050932	        25.1 ns/op	       0 B/op	       0 allocs/op
BenchmarkTypesStringBytes-4          	304471928	         3.86 ns/op	       0 B/op	       0 allocs/op
BenchmarkFmtSprintSlice-4            	  358131	      3173 ns/op	      80 B/op	       4 allocs/op
BenchmarkJsonMarshalSlice-4          	  569418	      2158 ns/op	      64 B/op	       2 allocs/op
BenchmarkTypesStringSliceGeneric-4   	  620298	      1908 ns/op	     112 B/op	       5 allocs/op
BenchmarkTypesStringSliceOptimal-4   	  707545	      1673 ns/op	      80 B/op	       4 allocs/op
BenchmarkFmtSprint1-4                	   37072	     30927 ns/op	    1920 B/op	      41 allocs/op
BenchmarkJsonMarshal1-4              	   44587	     27779 ns/op	    2336 B/op	      43 allocs/op
BenchmarkTypesStringDict1-4          	  158878	      7201 ns/op	     384 B/op	       1 allocs/op
BenchmarkTypesStringMap1-4           	  132950	      8718 ns/op	     384 B/op	       1 allocs/op
BenchmarkFmtSprint2-4                	    1658	    670532 ns/op	   44627 B/op	     992 allocs/op
BenchmarkJsonMarshal2-4              	    1888	    601376 ns/op	   50251 B/op	    1005 allocs/op
BenchmarkTypesStringDict2-4          	    6336	    183910 ns/op	    9651 B/op	     170 allocs/op
BenchmarkTypesStringMap2-4           	    6184	    187285 ns/op	    9651 B/op	     170 allocs/op
BenchmarkFmtSprint3-4                	    1125	   1022986 ns/op	   66236 B/op	    1508 allocs/op
BenchmarkJsonMarshal3-4              	    1238	    907943 ns/op	   77181 B/op	    1538 allocs/op
BenchmarkTypesStringDict3-4          	    4309	    275846 ns/op	   15163 B/op	     253 allocs/op
BenchmarkTypesStringMap3-4           	    4171	    281199 ns/op	   15163 B/op	     253 allocs/op
```
