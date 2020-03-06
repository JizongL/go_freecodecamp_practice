package main

import "fmt"

// import "time"
import "sync"

var wg = sync.WaitGroup{}

// scope is applied. msg is available in side the go routine.
func main() {
	// decouple the msg in outside from the inner scope.
	var msg = "hello"
	wg.Add(1)
	go func(msg string) {
		fmt.Println(msg)
		wg.Done()
	}(msg)
	// race condition is genearated, a bad thing, avoid.
	msg = "Goodbye"
	// time.Sleep(100 * time.Millisecond)
	wg.Wait()
}
