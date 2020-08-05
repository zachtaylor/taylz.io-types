# Benchmarks

This file contains benchmark comparisons to the standard library with various payload sizes

## Results

```
goos: linux
goarch: amd64
pkg: taylz.io/types
BenchmarkBuiltinStringBytes-4        	46986300	        25.3 ns/op	       0 B/op	       0 allocs/op
BenchmarkTypesStringBytes-4          	305424681	         3.87 ns/op	       0 B/op	       0 allocs/op
BenchmarkFmtSprintSlice-4            	  173694	      6679 ns/op	     384 B/op	      11 allocs/op
BenchmarkJsonMarshalSlice-4          	  213543	      5568 ns/op	     336 B/op	       8 allocs/op
BenchmarkTypesStringSliceGeneric-4   	  372691	      3083 ns/op	     144 B/op	       5 allocs/op
BenchmarkTypesStringSliceOptimal-4   	  418513	      2826 ns/op	     112 B/op	       4 allocs/op
BenchmarkFmtSprint1-4                	   35274	     34690 ns/op	    1920 B/op	      41 allocs/op
BenchmarkJsonMarshal1-4              	   40518	     29609 ns/op	    2336 B/op	      43 allocs/op
BenchmarkTypesStringDict1-4          	  158910	      7283 ns/op	     384 B/op	       1 allocs/op
BenchmarkTypesStringMap1-4           	  119289	      9507 ns/op	     384 B/op	       1 allocs/op
BenchmarkFmtSprint2-4                	    1603	    735648 ns/op	   44641 B/op	     992 allocs/op
BenchmarkJsonMarshal2-4              	    1849	    646364 ns/op	   50251 B/op	    1005 allocs/op
BenchmarkTypesStringDict2-4          	    6219	    194017 ns/op	    9651 B/op	     170 allocs/op
BenchmarkTypesStringMap2-4           	    6117	    194768 ns/op	    9651 B/op	     170 allocs/op
BenchmarkFmtSprint3-4                	    1078	   1106351 ns/op	   66212 B/op	    1508 allocs/op
BenchmarkJsonMarshal3-4              	    1122	    986201 ns/op	   77189 B/op	    1538 allocs/op
BenchmarkTypesStringDict3-4          	    3943	    294110 ns/op	   15164 B/op	     253 allocs/op
BenchmarkTypesStringMap3-4           	    3926	    292863 ns/op	   15164 B/op	     253 allocs/op
```
