package randor

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

type spyRandom struct {
	callIntFunc int
}

func (f *spyRandom) Int() int {
	f.callIntFunc++
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
