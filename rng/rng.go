package rng

import "math/big"

type Rng interface {
	GenBigInt() *big.Int
	GenBytes() []byte
}
