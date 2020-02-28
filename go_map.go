package main

import (
	"fmt"
)

func main() {
	statePopulations := map[string]int{
		"California": 12345,
		"Florida":    12345,
		"New York":   22333,
	}
	fmt.Println(statePopulations)
}
