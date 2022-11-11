package belajar_golang_generics

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

type Employee interface {
	GetName() string
}

func GetName[T Employee](param T) string {
	return param.GetName()
}

type Manager interface {
	GetName() string
	GetManagerName() string
}

type MyManager struct {
	Name string
}

func (m *MyManager) GetName() string {
	return m.Name
}

func (m *MyManager) GetManagerName() string {
	return m.Name
}

type VicePresident interface {
	GetName() string
	GetVicePresidentName() string
}

type MyVicePresident struct {
	Name string
}

func (m *MyVicePresident) GetName() string {
	return m.Name
}
func (m *MyVicePresident) GetVicePresidentName() string {
	return m.Name
}

func TestInheritance(t *testing.T) {
	assert.Equal(t, "Jalal", GetName[Manager](&MyManager{Name: "Jalal"}))
	assert.Equal(t, "Jalal", GetName[VicePresident](&MyVicePresident{Name: "Jalal"}))
}
