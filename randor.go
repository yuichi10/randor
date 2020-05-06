package randor

import (
	"math/rand"
	"time"
)

var random Rand

// Rand is math/rand Rand struct
type Rand interface {
	Int() int
	Int63() int64
}

func init() {
	random = rand.New(rand.NewSource(time.Now().UnixNano()))
}

// Int return random int
func Int() int {
	return random.Int()
}

// Int64 return random int64
func Int64() int64 {
	return random.Int63()
}

// Uint64 return random uint64
func Uint64() uint64 {
	return uint64(random.Int63())
}
