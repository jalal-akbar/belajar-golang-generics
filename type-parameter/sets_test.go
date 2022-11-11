package belajar_golang_generics

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

type Number interface {
	int | int8 | int32 | int64 |
		float32 | float64
}

func Min[T Number](v1, v2 T) T {
	if v1 < v2 {
		return v1
	} else {
		return v2
	}
}

func TestSets(t *testing.T) {
	assert.Equal(t, 100, Min(100, 100))                                // int
	assert.Equal(t, int8(1), Min(int8(1), int8(1)))                    // int8
	assert.Equal(t, int32(10), Min(int32(10), int32(20)))              // int32
	assert.Equal(t, int64(1000), Min(int64(1000), int64(2000)))        // int64
	assert.Equal(t, float32(1.0), Min(float32(1.0), float32(9.0)))     // float32
	assert.Equal(t, float64(99.0), Min(float64(100.0), float64(99.0))) // float64
}
