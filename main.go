package main

import "fmt"

func main() {

	const divider string = "--- --- ---"

	fmt.Println("Hello World, go explore!")
	name := "Leander"
	fmt.Println("My cat's name is", name)

	fmt.Println(divider)

	const age int = 34
	var favNum float64 = 1.6180339
	randNum := 1
	fmt.Println(favNum, randNum, age)

	fmt.Println(divider)

	var occassion = "New Year"

	if occassion == "New Year" {
		fmt.Println("新年快乐")
		fmt.Println(len(occassion))
		fmt.Printf("%T \n", occassion)
	} else {
		fmt.Println("圣诞节快乐")
		fmt.Println(len(occassion))
	}

	fmt.Println(divider)

	const specialNum = 1337
	fmt.Printf("Binary: %b \n", specialNum)
	fmt.Printf("Decimal: %d \n", specialNum)
	fmt.Printf("Hex: %x \n", specialNum)

}
