// interface naming convention, put er after the method.
// e.g Write => Writer

package main

import (
	"fmt"
)

func main() {
	var w Writer = ConsoleWriter{}
	w.Write([]byte("Hello Go!"))
}

// only contain behaviors not content.
type Writer interface {
	Write([]byte) (int, error)
}

// implicitly impliment interface
type ConsoleWriter struct{}

func (cw ConsoleWriter) Write(data []byte) (int, error) {
	n, err := fmt.Println(string(data))
	return n, err
}
