# comparing between map and array in go

```
go test -bench=.
```

outpout
```
goos: linux
goarch: amd64
pkg: map_vs_array
cpu: AMD Ryzen 7 PRO 5850U with Radeon Graphics     
Benchmark1M/map-16                             0.0000003 ns/op
Benchmark1M/array-16                           0.0000001 ns/op
Benchmark100M/map-16                           0.0000047 ns/op
Benchmark100M/array-16                         0.0000002 ns/op
```

