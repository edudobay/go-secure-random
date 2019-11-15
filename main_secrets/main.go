package main

import (
	"flag"
	"fmt"
	"os"
	"secureRng/benchmark"
	"secureRng/secrets"
)

const DefaultBytes = 32
const PrintQueueSize = 4096

func printTokens(tokens chan string, printer func(string), quit chan bool) {
	for token := range tokens {
		printer(token)
	}
	quit <- true
}

func printer(quiet bool) func(string) {
	if !quiet {
		return func(token string) {
			fmt.Println(token)
		}
	} else {
		return func(_ string) {}
	}
}

func main() {
	var bytes = flag.Int("bytes", DefaultBytes, "byte length of generated secret")
	var count = flag.Int("count", 1, "generate this many secrets")
	var quiet = flag.Bool("quiet", false, "do not output anything (only benchmark)")

	flag.Parse()

	tokens := make(chan string, PrintQueueSize)
	quit := make(chan bool)

	loop := func() {
		tokens <- secrets.GenToken(*bytes)
	}

	go printTokens(tokens, printer(*quiet), quit)

	result := benchmark.Run(loop, *count)

	close(tokens)
	<-quit

	// Ignore errors
	_, _ = fmt.Fprintf(os.Stderr, "Generated %d secrets in %v; average %.4f µs per secret",
		*count,
		result.Duration,
		result.AverageMicroseconds)
}
