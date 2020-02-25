package main

import(
	"fmt"
)

func main(){
	// every time a variable is initialized, it's false by default
	var n bool
	fmt.Printf("%v %T\n",n,n)
	fmt.Println(5%4)
	a:=10 // 1010
	b:=3 // 0011
	fmt.Println(a & b) // 0010
	fmt.Println(a | b) // 1011
	fmt.Println(a ^ b) // 1001
	fmt.Println(a &^ b) // 0100
}