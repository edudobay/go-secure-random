package main

import (
	"crypto/rand"
	"math/big"
	"secureRng/benchmark"
)

// ----------------------------------------------------------------------------

var ByteMaxInt = big.NewInt(0xff)

func GenBytes() []byte {
	buf := make([]byte, benchmark.Bytes)
	for i := range buf {
		generated, err := rand.Int(rand.Reader, ByteMaxInt)
		if err != nil {
			panic(err)
		}

		buf[i] = byte(generated.Int64())
	}

	return buf
}

// ----------------------------------------------------------------------------

func main() {
	benchmark.Benchmark(func() {
		GenBytes()
	}, benchmark.Times)
}
