package randor

import (
	"math"
	"math/rand"
	"time"
)

var random Rand
var str *strSettings

// Rand is math/rand Rand struct
type Rand interface {
	Int() int
	Int63() int64
	Intn(int) int
	NormFloat64() float64
}

type strSettings struct {
	chars  []rune
	length int
}

func init() {
	random = rand.New(rand.NewSource(time.Now().UnixNano()))
	str = &strSettings{chars: []rune("abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ"), length: 8}
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

// Intn returns, as an int, a non-negative pseudo-random number in [0,n) from the default Source. It panics if n <= 0.
func Intn(n int) int {
	return random.Intn(n)
}


// IntRange return negative or positive number between -n < 0 < n
func IntRange(n int) int {
	sign := 1
	if random.Intn(2) == 0 {
		sign = -1
	}
	return random.Intn(n) * sign
}

// Float64 return random float64
func Float64(decimal int) float64 {
	r := random.NormFloat64()
	if decimal > 0 {
		pow := math.Pow10(decimal)
		return math.Floor(r*pow) / pow
	}
	return r
}

// StrOption able to customize what kind of charactor return
type StrOption func(s *strSettings) error

// StrLength able to handle string length
func StrLength(l int) StrOption {
	return func(s *strSettings) error {
		s.length = l
		return nil
	}
}

// String return random string
func String(options ...StrOption) string {
	for _, o := range options {
		o(str)
	}
	res := make([]rune, str.length)
	for i := range res {
		res[i] = str.chars[random.Intn(len(str.chars))]
	}
	return string(res)
}
