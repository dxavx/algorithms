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

func TestBinarySearch(t *testing.T) {

	index := 3
	value := 555

	arr := RandomFillArray(MakeArray(100), 100)

	arr[index] = value

	position := BinarySearch(arr, value)
	assert.Equal(t, value, arr[position])
}

func TestFact(t *testing.T) {
	assert.Equal(t, 5040, Fact(7))
}

func TestMinMax(t *testing.T) {
	aaa := []float64{4.0, 34, -28883, 25, 32, 324, -23, 2, -333, 0, 99, 88888}
	min, max := MinMax(aaa)
	assert.Equal(t, min, aaa[2])
	assert.Equal(t, max, aaa[11])
}
