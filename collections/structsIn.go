package collections

import (
	"fmt"
)

type Doctor struct {
	number     int
	actorName  string
	companions []string
}

func DoctorStruct() {
	aDoctor := Doctor{
		number:     856759,
		actorName:  "John F Kennedy",
		companions: []string{"Lia", "Shawn", "Deva"},
	}
	fmt.Println(aDoctor)
	fmt.Println(aDoctor.actorName) //dot syntax

	bDoctor := &aDoctor
	bDoctor.actorName = "Ajay Route"
	fmt.Println(bDoctor)
	fmt.Println(aDoctor)

	c := aDoctor.companions
	fmt.Println(c[1])
}
