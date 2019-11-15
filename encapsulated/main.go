package main

import (
	"encoding/base64"
	"fmt"
	"secureRng/benchmark"
	"secureRng/rng"
)

// ----------------------------------------------------------------------------

func sample(rng rng.Rng, num int) {
	for i := 0; i < num; i++ {
		n := rng.GenBigInt()
		if i < benchmark.PrintSamples {
			fmt.Println(base64.RawURLEncoding.EncodeToString(n.Bytes()))
		}
	}
}

func main() {
	bigIntRng := rng.NewBigIntRng(benchmark.Bytes)
	bytesRng := rng.NewBytesRng(benchmark.Bytes)

	fmt.Println("## sample")
	fmt.Println("### BigIntRng")
	sample(bigIntRng, benchmark.SampleSize)

	fmt.Println("### BytesRng")
	sample(bytesRng, benchmark.SampleSize)

	fmt.Println("## benchmark")

	fmt.Printf("## BigIntRng\n")
	benchmark.RngBenchmark(bigIntRng, benchmark.Times)

	fmt.Printf("## BytesRng\n")
	benchmark.RngBenchmark(bytesRng, benchmark.Times)
}
