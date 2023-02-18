package main

import (
	"fmt"
	"reflect"
)

type Doctor struct {
	number     int
	actorName  string
	companions []string
}

type ExportedDoctor struct {
	Number     int
	ActorName  string
	Companions []string
}

// GoLang does NOT have inheritance
// It uses composition through embedding
type Animal struct {
	Name   string `required max:"100"` //tag
	Origin string
}

type Bird struct {
	Animal //B.Name
	// Animal   Animal   B.Animal.Name
	SpeedKPH float32
	CanFly   bool
}

func main() {
	t := reflect.TypeOf(Animal{})
	field, _ := t.FieldByName("Name")
	fmt.Println(field.Tag)

	//Struct is NOT referenced type
	//Passing it as an argument = copying
	//Want to pass by ref? add an ampersand

	aDoctor := Doctor{
		number:     3,
		actorName:  "Jon Pertwee",
		companions: []string{"a", "b"},
	}

	fmt.Println(aDoctor)
	fmt.Println(aDoctor.actorName)
	fmt.Println(aDoctor.companions[0])

	b := Bird{}
	b.Name = "Emu"
	b.Origin = "Australia"
	b.SpeedKPH = 48
	b.CanFly = false

	// SAME THING
	// b := Bird{
	// 	Animal: Animal{Name: "Emu", Origin: "Australia"},
	// 	SpeedKPH: 48,
	// 	CanFly: false,
	// }

}
