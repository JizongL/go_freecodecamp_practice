package main

import (
	"fmt"
)

const (
	// if error is assigned, it could avoid default 0 problem.
	// errorSpecialist = iota
	// if you don't care 0 value, then use _
	// _ = iota
	catSpecialist = iota
	dogSpecialist
	snakeSpecialist
)

func main() {
	// if specialistType int is created without equalling catSpecialist,
	// then it will be 0, which equals to catSpecialist, which is also 0.
	// then this will cause bug.
	var specialistType int = catSpecialist
	fmt.Printf("%v\n", specialistType == catSpecialist)
}
