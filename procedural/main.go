package main

import (
	"fmt"
	"secureRng/benchmark"
	"secureRng/rng"
)

// ----------------------------------------------------------------------------

func main() {
	fmt.Printf("## BigIntRng\n")
	maxInt := rng.MaxBigInt(benchmark.Bytes)
	benchmark.Benchmark(func() {
		rng.GenBigInt(maxInt)
	}, benchmark.Times)

	fmt.Printf("## BytesRng\n")
	benchmark.Benchmark(func() {
		rng.GenBytes(benchmark.Bytes)
	}, benchmark.Times)
}
