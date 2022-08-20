package collections

//Its true that go is kinda a different from traditional OOPs
// It doesnot supports inheritance but we have work arounds

import (
	"fmt"
	"reflect"
)

type Animal struct {
	Name   string `required max:"5"`
	Origin string `required min:"8"`
}

// Embedded in or Bird has Animals Property
type Bird struct {
	Animal
	SpeedKPH float32
	CanFly   bool
}

func StructInheritance() {
	bird1 := Bird{Animal: Animal{Name: "Emu", Origin: "Australia"}}
	bird1.CanFly = false
	bird1.SpeedKPH = 450
	fmt.Println(bird1)
	t := reflect.TypeOf(Animal{})
	field, _ := t.FieldByName("Origin")
	fmt.Println(field.Tag)
	fmt.Println()
}
