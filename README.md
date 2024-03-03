# concurrent__file_searcher_go
a go based repo to benchmark concurrent file based search


# worker count: 1, each worker sleeps 1 millisecond, to simulate complex processing logic

```
cpu: Intel(R) Core(TM) i7-10870H CPU @ 2.20GHz
=== RUN   BenchmarkSum
BenchmarkSum
File already exists: C:\Users\gupta\concurrent_file_searcher_go\LoadDirectory\test.txt
error will be on line: 8890
Content written to the file.
workers:  1
found at line: 8890
stopping writing to channel
BenchmarkSum-16
       1        139235777600 ns/op        439392 B/op       8938 allocs/op
PASS
ok      conccurent_File_searcher/BenchMark      139.460s

```


# worker count: 10, each worker sleeps 1 millisecond, to simulate complex processing logic

```
goos: windows
goarch: amd64
pkg: conccurent_File_searcher/BenchMark
cpu: Intel(R) Core(TM) i7-10870H CPU @ 2.20GHz
=== RUN   BenchmarkSum
BenchmarkSum
File already exists: C:\Users\gupta\concurrent_file_searcher_go\LoadDirectory\test.txt
error will be on line: 8890
Content written to the file.
workers:  10
found at line: 8890
stopping writing to channel
BenchmarkSum-16
       1        13962974100 ns/op         445792 B/op       9010 allocs/op
PASS
ok      conccurent_File_searcher/BenchMark      14.184s
```

# worker count: 100, each worker sleeps 1 millisecond, to simulate complex processing logic

```
Running tool: C:\Program Files\Go\bin\go.exe test -benchmem -run=^$ -bench ^BenchmarkSum$ conccurent_File_searcher/BenchMark

goos: windows
goarch: amd64
pkg: conccurent_File_searcher/BenchMark
cpu: Intel(R) Core(TM) i7-10870H CPU @ 2.20GHz
=== RUN   BenchmarkSum
BenchmarkSum
File already exists: C:\Users\gupta\concurrent_file_searcher_go\LoadDirectory\test.txt
error will be on line: 8890
Content written to the file.
workers:  100
found at line: 8890
stopping writing to channel
BenchmarkSum-16
       1        1494037500 ns/op          527216 B/op       9589 allocs/op
PASS
ok      conccurent_File_searcher/BenchMark      1.726s
```

