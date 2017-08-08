package main

import "fmt"

func main() {
	i := 0
	for i <= 10 {
		fmt.Println(i)
		i++
	}

	for a := 0; a < 5; a++ {
		fmt.Println(a)
	}

	yourAge := 18

	if yourAge >= 16 {
		fmt.Println("You can drive!")
	} else if yourAge >= 18 {
		fmt.Println("You can vote!")
	} else {
		fmt.Println("You can have fun")
	}

	const occassion string = "Birthday"

	switch occassion {
	case "New Year":
		fmt.Println("新年快乐")
	case "Christmas":
		fmt.Println("圣诞节快乐")
	case "Birthday":
		fmt.Println("生日快乐")
	default:
		fmt.Println("你好")
	}

}
