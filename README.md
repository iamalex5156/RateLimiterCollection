<h1 align="center">rlc</h1>
<p align="center">A rate limiter collection for Go.</p>

Pick up one of the rate limiters to throttle requests and control quota.

- [RLC](#rlc)
- [Slider](#slider)
- [TokenBucket](#tokenbucket)

## RLC

`RLC` is a rate limiter that manages quota based on previous requests' statuses and slows down or accelerates accordingly.

### Usage

```go
offset := rlc.NewRandomOffset(96)

// an RLC with the bitmap size of 96 with 16 reserved
// positive bits and the random offset manager
eq := rlc.NewRLC(96, 16, offset)

// non-blocking quota request
haveQuota := eq.Ask()

// update with ten previous successful requests
eq.Notify(true, 10)
```

### Benchmarks

```sh
BenchmarkRLCShortAskStep-16       30607452                37.5 ns/op             0 B/op          0 allocs/op
BenchmarkRLCShortAskRandom-16     31896340                34.5 ns/op             0 B/op          0 allocs/op
BenchmarkRLCShortNotify-16        12715494                81.9 ns/op             0 B/op          0 allocs/op
BenchmarkRLCLongAskStep-16        34627239                35.4 ns/op             0 B/op          0 allocs/op
BenchmarkRLCLongAskRandom-16      32399748                34.0 ns/op             0 B/op          0 allocs/op
BenchmarkRLCLongNotify-16            59935               20343 ns/op             0 B/op          0 allocs/op
```

## Slider

`Slider` rate limiter is based on a sliding window with a specified quota capacity.
Implements the `Limiter` interface.

### Usage

```go
// a Slider with one second window size, 100 millis sliding interval
// and the capacity of 32
slider := rlc.NewSlider(time.Second, time.Millisecond*100, 32)

// non-blocking quota request
haveQuota := slider.Ask()

// blocking call
slider.Take()
```

### Benchmarks

```sh
BenchmarkSliderShortWindow-16           123488035                9.67 ns/op            0 B/op          0 allocs/op
BenchmarkSliderLongerWindow-16          128023276                9.76 ns/op            0 B/op          0 allocs/op
```

## TokenBucket

`TokenBucket` rate limiter is based on the token bucket algorithm with a refill interval.
Implements the `Limiter` interface.

### Usage

```go
// a TokenBucket with the capacity of 32 and 100 millis refill interval
tokenBucket := rlc.NewTokenBucket(32, time.Millisecond*100)

// non-blocking quota request
haveQuota := tokenBucket.Ask()

// blocking call
tokenBucket.Take()
```

### Benchmarks

```sh
BenchmarkTokenBucketDenseRefill-16      212631714                5.64 ns/op            0 B/op          0 allocs/op
BenchmarkTokenBucketSparseRefill-16     211491368                5.63 ns/op            0 B/op          0 allocs/op
```

## License

Licensed under the MIT License.
