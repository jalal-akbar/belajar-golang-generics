package belajargolanggenerics

import (
	"testing"

	"github.com/stretchr/testify/assert"
	"golang.org/x/exp/constraints"
	"golang.org/x/exp/maps"
	"golang.org/x/exp/slices"
)

func ExperimentalMin[T constraints.Ordered](value1, value2 T) T {
	if value1 < value2 {
		return value1
	} else {
		return value2
	}
}

func TestExperimentalMin(t *testing.T) {
	assert.Equal(t, 1, ExperimentalMin(1, 2))
}

// Experimental Maps
func TestExperimentalMaps(t *testing.T) {
	first := map[string]string{
		"Name": "Jalal",
	}
	second := map[string]string{
		"Name": "Jalal",
	}

	assert.True(t, maps.Equal(first, second))
}

// Experimental Slice
func TestExperimentalSlice(t *testing.T) {
	first := []string{"Jalal"}
	second := []string{"Jalal"}

	assert.True(t, slices.Equal(first, second))
}
