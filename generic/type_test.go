package belajar_golang_generics

import (
	"fmt"
	"testing"
)

type Bag[T any] []T

func PrintBag[T any](value Bag[T]) {
	for _, bag := range value {
		fmt.Println(bag)
	}
}

func TestBagString(t *testing.T) {
	value := Bag[string]{"Jalal", "Muh", "Akbar"}
	PrintBag(value)
}

func TestBagInt(t *testing.T) {
	value := Bag[int]{1, 2, 3}
	PrintBag(value)
}
