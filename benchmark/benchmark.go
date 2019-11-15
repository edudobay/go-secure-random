package benchmark

import (
	"fmt"
	"time"
)

type Result struct {
	Duration            time.Duration
	AverageMicroseconds float64
	Times               int
}

func Run(fn func(), times int) Result {
	startedAt := time.Now()

	for i := 0; i < times; i++ {
		fn()
	}

	endedAt := time.Now()
	duration := endedAt.Sub(startedAt)

	avg := float64(duration.Nanoseconds()) / 1000.0 / float64(times)

	return Result{Duration: duration, AverageMicroseconds: avg, Times: times}
}

func RunAndDisplay(fn func(), times int) {
	result := Run(fn, times)
	fmt.Printf("%d runs took %v; average per run %.4f µs\n",
		result.Times,
		result.Duration,
		result.AverageMicroseconds)
}
