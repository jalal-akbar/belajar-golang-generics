package belajar_golang_generics

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func Comparable[T comparable](value1, value2 T) bool {
	if value1 == value2 {
		return true
	} else {
		return false
	}
}

func TestComparable(t *testing.T) {
	assert.True(t, Comparable("Jalal", "Jalal"))
	assert.True(t, Comparable(100, 100))
	assert.True(t, Comparable(true, true))
}
