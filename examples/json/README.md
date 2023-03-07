# Benchmarks

## Results

### [gobble](https://github.com/roblovelock/gobble)

```
goos: windows
goarch: amd64
pkg: github.com/roblovelock/gobble/examples/json
cpu: AMD Ryzen 7 2700 Eight-Core Processor
```

| benchmark           | op      | ns/op    | B/op    | allocs/op |
|---------------------|---------|----------|---------|-----------|
| BenchmarkJSONInt    | 3415627 | 340.2    | 56      | 4         |
| BenchmarkJSONFloat  | 2854645 | 426.5    | 80      | 5         |
| BenchmarkJSONString | 3179964 | 371.4    | 72      | 4         |
| BenchmarkJSONBool   | 6407236 | 183.7    | 36      | 2         |
| BenchmarkJSONNull   | 6457402 | 182.7    | 36      | 2         |
| BenchmarkJSONArray  | 535045  | 2105     | 424     | 18        |
| BenchmarkJSONMap    | 194258  | 5956     | 1472    | 42        |
| BenchmarkJSONMedium | 5558    | 204297   | 54258   | 1097      |
| BenchmarkJSONLarge  | 45      | 26507384 | 6427878 | 132411    |

### [goparsec](https://github.com/prataprc/goparsec)

```
goos: windows
goarch: amd64
pkg: github.com/prataprc/goparsec/json
cpu: AMD Ryzen 7 2700 Eight-Core Processor
```

| benchmark           | op      | ns/op    | B/op     | allocs/op |
|---------------------|---------|----------|----------|-----------|
| BenchmarkJSONInt    | 2222426 | 539.1    | 208      | 9         |
| BenchmarkJSONFloat  | 2172085 | 544.6    | 224      | 9         |
| BenchmarkJSONString | 2161797 | 557.8    | 224      | 9         |
| BenchmarkJSONBool   | 2595242 | 440.1    | 184      | 7         |
| BenchmarkJSONNull   | 2648426 | 460.4    | 184      | 7         |
| BenchmarkJSONArray  | 328233  | 3485     | 1482     | 49        |
| BenchmarkJSONMap    | 124335  | 10671    | 4320     | 125       |
| BenchmarkJSONMedium | 3165    | 342298   | 179523   | 3612      |
| BenchmarkJSONLarge  | 30      | 43706757 | 17178652 | 435310    |

### encoding/json

```
goos: windows
goarch: amd64
pkg: encoding/json
cpu: AMD Ryzen 7 2700 Eight-Core Processor
```

| benchmark           | op      | ns/op    | B/op    | allocs/op |
|---------------------|---------|----------|---------|-----------|
| BenchmarkJSONInt    | 1847432 | 620.2    | 184     | 5         |
| BenchmarkJSONFloat  | 1736043 | 678.2    | 200     | 5         |
| BenchmarkJSONString | 1972926 | 609.9    | 208     | 5         |
| BenchmarkJSONBool   | 2549192 | 461.3    | 168     | 3         |
| BenchmarkJSONNull   | 2602298 | 420.7    | 168     | 3         |
| BenchmarkJSONArray  | 528080  | 2044     | 552     | 15        |
| BenchmarkJSONMap    | 268762  | 4477     | 1136    | 25        |
| BenchmarkJSONMedium | 6835    | 169407   | 48895   | 741       |
| BenchmarkJSONLarge  | 56      | 21342052 | 5481636 | 88447     |


### [goparsify](https://github.com/vektah/goparsify)

```
goos: windows
goarch: amd64
pkg: github.com/vektah/goparsify/json
cpu: AMD Ryzen 7 2700 Eight-Core Processor
```

| benchmark           | op      | ns/op    | B/op     | allocs/op |
|---------------------|---------|----------|----------|-----------|
| BenchmarkJSONInt    | 2315056 | 534.3    | 208      | 9         |
| BenchmarkJSONFloat  | 2385096 | 523.4    | 224      | 9         |
| BenchmarkJSONString | 2232732 | 537.2    | 224      | 9         |
| BenchmarkJSONBool   | 2645852 | 438.4    | 184      | 7         |
| BenchmarkJSONNull   | 2719545 | 445.0    | 184      | 7         |
| BenchmarkJSONArray  | 357579  | 3375     | 1482     | 49        |
| BenchmarkJSONMap    | 132651  | 9071     | 4320     | 125       |
| BenchmarkJSONMedium | 3738    | 316172   | 179525   | 3612      |
| BenchmarkJSONLarge  | 30      | 39340097 | 17175839 | 435300    |