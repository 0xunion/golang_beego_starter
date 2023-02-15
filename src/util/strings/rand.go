package strings

import (
	"math/rand"
	"time"
)

const (
	ALPHAS = "abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ"
	ALPHA_LEN = len(ALPHAS)
	NUMS   = "0123456789"
	NUM_LEN = len(NUMS)
)

func init() {
	rand.Seed(time.Now().UnixNano())
}

func RandomAlphaString(len int) string {
	result := make([]byte, len)
	for i := 0; i < len; i++ {
		result[i] = ALPHAS[rand.Intn(ALPHA_LEN)]
	}

	return Bytes2String(result)
}
