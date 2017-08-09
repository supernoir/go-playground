package main

import (
	"fmt"
	"sort"
	"strings"
)

func main() {
	sampleString := "Hello World"
	fmt.Println(strings.Contains(sampleString, "Hell"))
	fmt.Println(strings.Index(sampleString, "He"))
	fmt.Println(strings.Count(sampleString, "l"))
	fmt.Println(strings.Replace(sampleString, "l", "x", 3))

	csvString := "1,2,3,4,5,6"
	fmt.Println(strings.Split(csvString, ","))
	listOfLetters := []string{"c", "a", "b"}
	sort.Strings(listOfLetters)
	fmt.Println("Letters:", listOfLetters)

	listOfNums := strings.Join([]string{"3", "2", "1"}, ", ")
	fmt.Println(listOfNums)
}
