package algorithms

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestMakeArray(t *testing.T) {
	n := 10
	nm := MakeArray(n)
	assert.Equal(t, n, len(nm))
}

func TestRandomFillArray(t *testing.T) {
	value := false
	nm := MakeArray(10)
	RandomFillArray(nm, 100)
	if nm[0] != nm[1] {
		value = true
	}

	assert.True(t, value)
}
