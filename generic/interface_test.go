package belajar_golang_generics

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

type SetterGetter[T any] interface {
	GetValue() T
	SetValue(value T)
}

// func implementation
func ChangeValue[T any](param SetterGetter[T], value T) T {
	param.SetValue(value)
	return value
}

// struct implementation
type MyData[T any] struct {
	Value T
}

func (d *MyData[T]) GetValue() T {
	return d.Value
}

func (d *MyData[T]) SetValue(value T) {
	d.Value = value
}

func TestGenericInterface(t *testing.T) {
	myData := MyData[string]{}

	res := ChangeValue[string](&myData, "jalal")

	assert.Equal(t, "jalal", res)
	assert.Equal(t, "jalal", myData.Value)
}
