package belajar_golang_generics

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func FindMin[T interface{ int | int64 | float64 }](v1, v2 T) T {
	if v1 < v2 {
		return v1
	} else {
		return v2
	}
}

func TestInLine(t *testing.T) {
	assert.Equal(t, 1, FindMin(1, 2))
	assert.Equal(t, int64(100), FindMin[int64](100, 200))
	assert.Equal(t, float64(100.0), FindMin(100.0, 200.0))
}

// Generic Di Type Parameter
func GetFirst[T []E, E any](name T) E {
	first := name[0]
	return first
}

func TestGetFirst(t *testing.T) {
	names := []string{"jalal", "muh", "akbar"}
	assert.Equal(t, "jalal", GetFirst(names))
}
