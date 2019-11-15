package main

import (
	"flag"
	"fmt"
	"os"
	"secureRng/benchmark"
	"secureRng/secrets"
)

const DefaultBytes = 32

func main() {
	var bytes = flag.Int("bytes", DefaultBytes, "byte length of generated secret")
	var count = flag.Int("count", 1, "generate this many secrets (default: 1)")
	var quiet = flag.Bool("quiet", false, "do not output anything (only benchmark)")

	flag.Parse()

	loop := func() {
		token := secrets.GenToken(*bytes)
		if !*quiet {
			fmt.Println(token)
		}
	}

	result := benchmark.Run(loop, *count)

	// Ignore errors
	_, _ = fmt.Fprintf(os.Stderr, "Generated %d secrets in %v; average %.4f µs per secret",
		*count,
		result.Duration,
		result.AverageMicroseconds)
}
