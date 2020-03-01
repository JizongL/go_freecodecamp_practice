package main

import (
	"fmt"
)

func main() {
	statePopulations :=
		map[string]int{
			"California": 39250017,
			"Texas":      33334343,
		}
	if pop, ok := statePopulations["California"]; ok {
		fmt.Println(pop)
	}
}
