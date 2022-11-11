package belajar_golang_generics

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

type Age int

type NumberWithAppro interface {
	~int | int8 | int32 | int64 | // gunakan tanda ~ pada tipe data untuk memperkirakan tipe data dasarnya
		float32 | float64
}

func Max[T NumberWithAppro](v1, v2 T) T {
	if v1 > v2 {
		return v1
	} else {
		return v2
	}
}

func TestApproxiamtion(t *testing.T) {
	assert.Equal(t, 200, Max(100, 200))                                 // int
	assert.Equal(t, int8(3), Max(int8(1), int8(3)))                     // int8
	assert.Equal(t, int32(20), Max(int32(10), int32(20)))               // int32
	assert.Equal(t, int64(2000), Max(int64(1000), int64(2000)))         // int64
	assert.Equal(t, float32(9.0), Max(float32(1.0), float32(9.0)))      // float32
	assert.Equal(t, float64(100.0), Max(float64(100.0), float64(99.0))) // float64
	assert.Equal(t, Age(4), Max(Age(4), Age(1)))                        // Approximation
}
