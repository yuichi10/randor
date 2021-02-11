package randor

import (
	"math/rand"
	"testing"

	"github.com/stretchr/testify/assert"
)

type spyRandom struct {
	r                   *rand.Rand
	callIntFunc         int
	callInt63Func       int
	callIntnFunc        int
	callNormFloat64Func int
	argIntn             int
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

func (f *spyRandom) NormFloat64() float64 {
	f.callNormFloat64Func++
	return f.r.Float64()
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

func TestIntn(t *testing.T) {
	r := random
	defer func() {
		random = r
	}()

	fake := newSpyRandom(5)
	random = fake

	Intn(3)

	assert.Equal(t, 1, fake.callIntnFunc)
	assert.Equal(t, 3, fake.argIntn)
}

func TestIntRange(t *testing.T) {
	r := random
	defer func() {
		random = r
	}()

	t.Run("positive number will return", func(t *testing.T) {
		fake := newSpyRandom(1)

		random = fake

		v := IntRange(10)

		assert.Equal(t, 2, fake.callIntnFunc)
		assert.Equal(t, 7, v)

	})

	t.Run("negative number will return", func(t *testing.T) {
		fake := newSpyRandom(2)

		random = fake

		v := IntRange(10)

		assert.Equal(t, 2, fake.callIntnFunc)
		assert.Equal(t, -6, v)
	})
}

func TestFloat64(t *testing.T) {
	r := random
	defer func() {
		random = r
	}()

	testCases := []struct {
		name    string
		decimal int
		expect  float64
	}{
		{
			name:    "decimal > 0",
			decimal: 0,
			expect:  0.6046602879796196,
		},
		{
			name:    "1 decimal",
			decimal: 1,
			expect:  0.6,
		},
		{
			name:    "5 decimal",
			decimal: 5,
			expect:  0.60466,
		},
	}

	for _, tt := range testCases {
		t.Run(tt.name, func(t *testing.T) {
			fake := newSpyRandom(1)
			random = fake

			result := Float64(tt.decimal)

			assert.Equal(t, 1, fake.callNormFloat64Func)
			assert.Equal(t, tt.expect, result)
		})
	}

	t.Run("use NormFloat64", func(t *testing.T) {
		fake := newSpyRandom(1)
		random = fake

		result := Float64(-1)

		assert.Equal(t, 1, fake.callNormFloat64Func)
		assert.Equal(t, 0.6046602879796196, result)
	})

	t.Run("return 1 decimal float", func(t *testing.T) {
		fake := newSpyRandom(1)
		random = fake

		result := Float64(1)

		assert.Equal(t, float64(6), result*10)
	})

	t.Run("return 5 decimal float", func(t *testing.T) {
		fake := newSpyRandom(1)
		random = fake

		result := Float64(5)

		assert.Equal(t, float64(60466), result*100000)
	})
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
