package benchmark

import (
	"fmt"
	"time"
)

func Benchmark(fn func(), times int) {
	startedAt := time.Now()

	for i := 0; i < times; i++ {
		fn()
	}

	endedAt := time.Now()
	duration := endedAt.Sub(startedAt)

	avg := float64(duration.Nanoseconds()) / 1000.0 / float64(times)

	fmt.Printf("%d runs took %v; average per run %.4f µs\n", times, duration, avg)
}