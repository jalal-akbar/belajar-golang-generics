package belajar_golang_generics

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"
)

func Length[T any](param T) T {
	fmt.Println(param)
	return param
}

func TestTypeParamaters(t *testing.T) {
	var resultString = Length("Jalal")
	resultInteger := Length(02)
	resultFloat := Length(2.10)
	resultBoolean := Length(true)

	assert.Equal(t, "Jalal", resultString)
	assert.Equal(t, 02, resultInteger)
	assert.Equal(t, 2.10, resultFloat)
	assert.Equal(t, true, resultBoolean)
}

// Multiple Type Parameter
func LengthMultiple[T1 any, T2 any](param1 T1, param2 T2) (T1, T2) {
	fmt.Println(param1)
	fmt.Println(param2)
	return param1, param2
}

func TestMultipleParameters(t *testing.T) {
	t1, t2 := LengthMultiple("Jalal", 10)

	assert.Equal(t, "Jalal", t1)
	assert.Equal(t, 10, t2)
}
