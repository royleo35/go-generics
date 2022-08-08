package myslice

// 利用泛型实现了slice
// 与builtin slice bench mark 对比结论: set 和 append性能稍差，get性能超出built-in slice 很多倍
// cpu: Intel(R) Core(TM) i7-8750H CPU @ 2.20GHz
//BenchmarkMySliceSet1
//BenchmarkMySliceSet1-12       	     698	   1455532 ns/op
//BenchmarkSliceSet1
//BenchmarkSliceSet1-12         	    1096	   1033981 ns/op
//BenchmarkMySliceGet1
//BenchmarkMySliceGet1-12       	100000000	        11.20 ns/op
//BenchmarkSliceGet1
//BenchmarkSliceGet1-12         	 1326634	       903.8 ns/op
//BenchmarkMySliceAppend1
//BenchmarkMySliceAppend1-12    	     624	   1775521 ns/op
//BenchmarkSliceAppend1
//BenchmarkSliceAppend1-12      	    1021	   1183980 ns/op
//BenchmarkMySliceAppend2
//BenchmarkMySliceAppend2-12    	    3001	    359658 ns/op
//BenchmarkSliceAppend2
//BenchmarkSliceAppend2-12      	    4987	    225115 ns/op
