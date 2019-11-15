package benchmark

import "secureRng/rng"

func RngBenchmark(rng rng.Rng, times int) {
	RunAndDisplay(func() {
		rng.GenBigInt()
	}, times)
}
