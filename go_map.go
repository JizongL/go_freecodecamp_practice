package main

import (
	"fmt"
)

type Doctor struct {
	number     int
	actorName  string
	companions []string
}

func main() {
	statePopulations := map[string]int{
		"California": 12345,
		"Florida":    12345,
		"New York":   22333,
	}
	fmt.Println(statePopulations)
	// can declared an empty map with make
	// statePopulation:=make(map[string]int)
	statePopulations["New Jersey"] = 12020
	// delete element of map
	delete(statePopulations, "New York")
	// , ok syntax, to check if key exists
	pop, ok := statePopulations["Florida"]
	fmt.Println(pop, ok) // 12345,true
	pop2, ok := statePopulations["New Mexico"]
	fmt.Println(pop2, ok) // will print 0, false
}
