# Testing performance

## Detecting a substring in a string

Run to compare `go test -bench . substring_test.go -benchmem`

Example output:
```
BenchmarkStringsContains-4              100000000               10.5 ns/op             0 B/op          0 allocs/op
BenchmarkStringsIndex-4                 117090943               10.1 ns/op             0 B/op          0 allocs/op
BenchmarkStringsSplit-4                  6958126               152 ns/op              32 B/op          1 allocs/op
BenchmarkStringsCount-4                 42397729                29.1 ns/op             0 B/op          0 allocs/op
BenchmarkStringsRegExp-4                  461696              2467 ns/op            1326 B/op         16 allocs/op
BenchmarkStringsRegExpCompiled-4         7109509               168 ns/op               0 B/op          0 allocs/op
```