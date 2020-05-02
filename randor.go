package randor

import (
	"math/rand"
	"time"
)

var random Rand

// Rand is math/rand Rand struct
type Rand interface {
	Int() int
}

func init() {
	random = rand.New(rand.NewSource(time.Now().UnixNano()))
}

// Int return random int
func Int() int {
	return random.Int()
}
