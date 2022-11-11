package belajar_golang_generics

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"
)

// Generic Struct
type Data[T any] struct {
	First  T
	Second T
}

func TestGenericStruct(t *testing.T) {
	data := Data[string]{
		First:  "Jalal",
		Second: "Akbar",
	}
	fmt.Println(data)
}

// Generic Method
func (d *Data[_]) SayHello(name string) string {
	return "Hello " + name
}

func (d *Data[T]) ChangeFirst(first T) T {
	d.First = first
	return first
}

func TestGenericMethod(t *testing.T) {
	data := Data[string]{
		First:  "Jalal",
		Second: "Akbar",
	}

	assert.Equal(t, "Sonia", data.ChangeFirst("Sonia"))
	assert.Equal(t, "Hello Sonia", data.SayHello("Sonia"))
}
