package main

import (
	"fmt"
)

func main() {
	listNums := []float64{1, 2, 3, 4, 5, 6}
	fmt.Println("Sum :", addThemUp(listNums))
	fmt.Println("Sub :", subtractThem(60, 80, 100, 120, 140))

	num3 := 3
	doubleNum := func() int {
		num3 *= 2
		return num3
	}
	fmt.Println("Double Num: ", doubleNum())
	fmt.Println("Double Num: ", doubleNum())

	fmt.Println("Factorial: ", factorial(3))

	// printHello2 called first,
	// but deferred until main function finished
	defer printHello2()
	printHello1()

	fmt.Println("Safe Division:", safeDivision(3, 0))
	fmt.Println("Safe Division:", safeDivision(3, 2))
}

func addThemUp(numbers []float64) float64 {
	sum := 0.0
	for _, val := range numbers {
		sum += val
	}
	return sum
}

func subtractThem(args ...int) int {
	finalValue := 0
	for _, value := range args {
		finalValue -= value
	}
	return finalValue
}

// Recursive function
func factorial(num int) int {
	if num == 0 {
		return 1
	}
	return num * factorial(num-1)
}

func printHello1() { fmt.Println("大家好") }
func printHello2() { fmt.Println("你們好") }

func safeDivision(num1, num2 int) int {
	defer func() {
		fmt.Println(recover())
	}()

	solution := num1 / num2
	return solution
}
