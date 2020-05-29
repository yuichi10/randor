package randor

import (
	"math/rand"
	"testing"

	"github.com/stretchr/testify/assert"
)

type spyRandom struct {
	r             *rand.Rand
	callIntFunc   int
	callInt63Func int
	callIntnFunc  int
	argIntn       int
}

func newSpyRandom(seed int64) *spyRandom {
	return &spyRandom{r: rand.New(rand.NewSource(seed))}
}

func (f *spyRandom) Int() int {
	f.callIntFunc++
	return f.r.Int()
}

func (f *spyRandom) Int63() int64 {
	f.callInt63Func++
	return f.r.Int63()
}

func (f *spyRandom) Intn(i int) int {
	f.callIntnFunc++
	f.argIntn = i
	return f.r.Intn(i)
}

func TestInt(t *testing.T) {
	r := random
	defer func() {
		random = r
	}()

	fake := newSpyRandom(1)
	random = fake

	Int()

	assert.Equal(t, 1, fake.callIntFunc)
}

func TestInt64(t *testing.T) {
	r := random
	defer func() {
		random = r
	}()

	fake := newSpyRandom(1)
	random = fake

	Int64()

	assert.Equal(t, 1, fake.callInt63Func)
}

func TestUint64(t *testing.T) {
	r := random
	defer func() {
		random = r
	}()

	fake := newSpyRandom(1)
	random = fake

	Uint64()

	assert.Equal(t, 1, fake.callInt63Func)
}

func TestString(t *testing.T) {
	t.Run("default return uppercase and lowercase", func(t *testing.T) {
		r := random
		s := str
		defer func() {
			random = r
			str = s
		}()

		fake := newSpyRandom(1)
		random = fake

		val := String()

		assert.Equal(t, s.chars, []rune("abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ"))
		assert.Equal(t, s.length, 8)
		assert.Equal(t, 8, fake.callIntnFunc)
		assert.Equal(t, fake.argIntn, 52)
		assert.Regexp(t, "^[a-zA-Z]{8}$", val)
		assert.Equal(t, "XVlBzgba", val)
	})

	t.Run("able to change length", func(t *testing.T) {
		r := random
		s := str
		defer func() {
			random = r
			str = s
		}()

		fake := newSpyRandom(1)
		random = fake

		digits := rand.Intn(15)
		val := String(StrLength(digits))

		assert.Equal(t, s.chars, []rune("abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ"))
		assert.Equal(t, s.length, digits)
		assert.Equal(t, digits, fake.callIntnFunc)
		assert.Equal(t, fake.argIntn, 52)
		assert.Regexp(t, "^[a-zA-Z].*$", val)
	})
}
