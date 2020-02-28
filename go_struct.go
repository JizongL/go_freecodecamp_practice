package main

import (
	"fmt"
)

// upper case for Doctor, and lower cases for all the fields.
type Doctor struct {
	number     int
	actorName  string
	companions []string
}

func main() {
	// independent data type.

	aDoctor := Doctor{
		number:    3,
		actorName: "Jon",
		companions: []string{
			"Liz Shaw",
			"Jo Grant",
		},
	}
	fmt.Println(aDoctor.actorName)

}
