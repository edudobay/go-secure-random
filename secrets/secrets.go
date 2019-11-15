package secrets

import (
	"encoding/base64"
	"secureRng/rng"
)

func genTokenFromRng(rng rng.Rng) string {
	b := rng.GenBytes()
	return base64.RawURLEncoding.EncodeToString(b)
}

func GenToken(bytes int) string {
	_rng := rng.NewBigIntRng(bytes)
	return genTokenFromRng(_rng)
}