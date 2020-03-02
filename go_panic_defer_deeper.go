package main

import (
	"fmt"
	"log"
)

func main() {
	fmt.Println("start")
	panicker()
	fmt.Println("end")
}

func panicker() {
	fmt.Println("About to panic")
	//  arguments evaluuated at time defer is excuted, not at time
	// of called function excution
	defer func() {
		if err := recover(); err != nil {
			log.Println("Error in panicker:", err)
			// repanic
			// if in situation, you try to recover fro the panic but
			// you realize you can't handle it,
			// you can recall panic and defer the management of that panic
			// in higher up stack.
			// function

			// only use panic when it's not recoverable. eg. divided by 0.
			panic(err)
		}
	}()
	panic("Something bad happened")
	fmt.Println("Done panicking")
}
