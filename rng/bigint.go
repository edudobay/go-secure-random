package rng

import (
	"crypto/rand"
	"math/big"
)

func GenBigInt(maxInt *big.Int) *big.Int {
	generated, err := rand.Int(rand.Reader, maxInt)
	if err != nil {
		panic(err)
	}
	return generated
}

func NewBigIntRng(bytes int) *BigIntRng {
	rng := new(BigIntRng)
	rng.bytes = bytes
	rng._maxInt = MaxBigInt(bytes)

	return rng
}

func MaxBigInt(bytes int) *big.Int {
	intMaxBytes := make([]byte, bytes)
	for i := range intMaxBytes {
		intMaxBytes[i] = 0xff
	}

	intMax := new(big.Int)
	intMax.SetBytes(intMaxBytes)

	return intMax
}

type BigIntRng struct {
	bytes   int
	_maxInt *big.Int
}

func (t *BigIntRng) GenBigInt() *big.Int {
	return GenBigInt(t._maxInt)
}

func (t *BigIntRng) GenBytes() []byte {
	return t.GenBigInt().Bytes()
}
