### Benchmarking Your Tool

Test files obtained from the resource code of the book's Website.

```bash
time ./colStats -op avg -col 2 testdata/benchmark/*.csv
```

To run the benchmarks, use the go test tool with the `-bench regexp` parameter. The regexp parameter is a regular expression that matches the benchmarks you want to execute. In this case, you can use the `.` (dot) to execute all the benchmarks. In addition, provide the argument `-run ^$` to skip running any of the tests in the test file while executing the benchmark to prevent impacting the results.

```bash
go test -bench . -run ^$
```

You can force additional executions by using the `-benchtime` parameter.

```bash
go test -bench . -benchtime=10x -run ^$
```

Save this output to a file so you can compare it with future executions later and see if it’s improving.

```bash
go test -bench . -benchtime=10x -run ^$ | tee benchresults00.txt
```

### Profiling Your Tool
The Go profiler shows you a breakdown of where your program spends its execution time. By running the profiler, you can determine which functions consume most of the program’s execution time and target them for optimization.

Run the benchmarks again, but this time, enable the CPU profiler:
```bash
go test -bench . -benchtime=10x -run ^$ -cpuprofile cpu00.pprof
```

Analyze the profiling results by using:
```bash
go tool pprof cpu00.pprof
```

commands to use on the `pprof` prompt are: 
- `top`
- `top -cum`
- `list` csv2float
- `web`

You can see how much memory the program is allocating by executing amemory profile.

```bash
go test -bench . -benchtime=10x -run ^$ -memprofile mem00.pprof
```

To view results:

```bash
go tool pprof -alloc_space mem00.pprof
```

Finally, using the parameter `-benchmem` to display the total memory allocation. (best option)
```bash
go test -bench . -benchtime=10x -run ^$ -benchmem | tee benchresults00m.txt
```

### Tracing Your Tool
```bash
go test -bench . -benchtime=10x -run ^$ -trace trace01.out
go tool trace trace01.out
```


