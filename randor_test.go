package randor

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

type spyRandom struct {
	callIntFunc   int
	callInt63Func int
}

func (f *spyRandom) Int() int {
	f.callIntFunc++
	return 0
}

func (f *spyRandom) Int63() int64 {
	f.callInt63Func++
	return 0
}

func TestInt(t *testing.T) {
	r := random
	defer func() {
		random = r
	}()

	fake := &spyRandom{}
	random = fake

	Int()

	assert.Equal(t, 1, fake.callIntFunc)
}

func TestInt64(t *testing.T) {
	r := random
	defer func() {
		random = r
	}()

	fake := &spyRandom{}
	random = fake

	Int64()

	assert.Equal(t, 1, fake.callInt63Func)
}

func TestUint64(t *testing.T) {
	r := random
	defer func() {
		random = r
	}()

	fake := &spyRandom{}
	random = fake

	Uint64()

	assert.Equal(t, 1, fake.callInt63Func)
}
