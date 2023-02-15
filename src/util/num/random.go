package num

import (
	"math/rand"
	"time"
)

func init() {
	rand.Seed(time.Now().UnixNano())
}

func Random(min, max int) int {
	return min + rand.Intn(max-min)
}
