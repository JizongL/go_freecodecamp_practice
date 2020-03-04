package main

import (
	"fmt"
	// "math"
)

type virusCount interface {
	infectedRate() float64
	// recoveredRate() float64
	deathRate() float64
	// testedRate() float64
}

type china struct {
	population, infected, died float64
}

func (ch china) infectedRate() float64 {
	return (ch.infected / ch.population) * 100
}

func (ch china) deathRate() float64 {
	return (ch.died / ch.population) * 100
}

func virusCensus(vc virusCount) {
	fmt.Println(vc.infectedRate())
	fmt.Println(vc.deathRate())
}

func main() {
	c := china{population: 1000, infected: 200}
	virusCensus(c)
}
