# String/Rune/Grapheme Cluster counting in Go

```console
$ go test -bench .
goos: linux
goarch: amd64
pkg: github.com/dolanor/gomisc/benchlen
cpu: Intel(R) Core(TM) i7-6700 CPU @ 3.40GHz
BenchmarkLen-8                    	1000000000	         0.3022 ns/op
BenchmarkRuneLen-8                	1000000000	         0.2902 ns/op
BenchmarkRuneCountInString-8      	69226528	        17.71 ns/op
BenchmarkRuneCount-8              	71911046	        16.71 ns/op
BenchmarkGraphemeClusterCount-8   	 5341496	       222.5 ns/op
PASS
ok  	github.com/dolanor/gomisc/benchlen	4.548s
```
