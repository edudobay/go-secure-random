package rng

import (
	"crypto/rand"
	"math/big"
)

type BytesRng struct {
	bytes int
}

var ByteMaxInt = big.NewInt(0xff)

func NewBytesRng(bytes int) *BytesRng {
	rng := new(BytesRng)
	rng.bytes = bytes
	return rng
}

func GenBytes(bytes int) []byte {
	buf := make([]byte, bytes)
	for i := range buf {
		buf[i] = randByte()
	}

	return buf
}

func (t *BytesRng) GenBytes() []byte {
	return GenBytes(t.bytes)
}

func (t *BytesRng) GenBigInt() *big.Int {
	bigNum := new(big.Int)
	bigNum.SetBytes(t.GenBytes())
	return bigNum
}

func randByte() byte {
	generated, err := rand.Int(rand.Reader, ByteMaxInt)
	if err != nil {
		panic(err)
	}
	return byte(generated.Int64())
}
