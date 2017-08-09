package main

import "fmt"

func main() {
	x := 0
	// pass in actual memory address
	changeXVal(&x)
	fmt.Println("X is", x)
	fmt.Println("Memory Address for X:", &x)
}

func changeXVal(x *int) {
	*x = 2
}
