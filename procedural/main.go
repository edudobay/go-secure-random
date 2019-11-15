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
	benchmark.RunAndDisplay(func() {
		rng.GenBigInt(maxInt)
	}, benchmark.Times)

	fmt.Printf("## BytesRng\n")
	benchmark.RunAndDisplay(func() {
		rng.GenBytes(benchmark.Bytes)
	}, benchmark.Times)
}
