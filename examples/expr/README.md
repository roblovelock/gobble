# Benchmarks

## Results

### [gobble](https://github.com/roblovelock/gobble)

```
goos: windows
goarch: amd64
pkg: github.com/roblovelock/gobble/examples/expr
cpu: AMD Ryzen 7 2700 Eight-Core Processor
```

| benchmark        | op      | ns/op | B/op | allocs/op |
|------------------|---------|-------|------|-----------|
| BenchmarkExpr1Op | 1000000 | 1017  | 112  | 9         |
| BenchmarkExpr2Op | 958910  | 1281  | 144  | 11        |
| BenchmarkExpr3Op | 664322  | 1778  | 216  | 18        |
| BenchmarkExpr    | 206222  | 5131  | 640  | 47        |

### [goparsec](https://github.com/prataprc/goparsec)

```
goos: windows
goarch: amd64
pkg: github.com/prataprc/goparsec/expr
cpu: AMD Ryzen 7 2700 Eight-Core Processor
```

| benchmark        | op    | ns/op | B/op  | allocs/op |
|------------------|-------|-------|-------|-----------|
| BenchmarkExpr1Op | 48468 | 24869 | 13463 | 170       |
| BenchmarkExpr2Op | 41455 | 28431 | 14978 | 191       |
| BenchmarkExpr3Op | 36715 | 32265 | 16705 | 213       |
| BenchmarkExpr    | 18974 | 62515 | 30292 | 401       |

### [goparsify](https://github.com/vektah/goparsify)

```
goos: windows
goarch: amd64
pkg: github.com/vektah/goparsify/calc
cpu: AMD Ryzen 7 2700 Eight-Core Processor
```

| benchmark        | op     | ns/op | B/op | allocs/op |
|------------------|--------|-------|------|-----------|
| BenchmarkExpr1Op | 432292 | 2369  | 1816 | 17        |
| BenchmarkExpr2Op | 470802 | 2471  | 1936 | 19        |
| BenchmarkExpr3Op | 406051 | 2831  | 2056 | 21        |
| BenchmarkExpr    | 149031 | 8167  | 6472 | 58        |