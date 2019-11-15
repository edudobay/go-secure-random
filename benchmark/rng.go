package benchmark

import "secureRng/rng"

func RngBenchmark(rng rng.Rng, times int) {
	Benchmark(func() {
		rng.GenBigInt()
	}, times)
}
