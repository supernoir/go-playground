package main

import "fmt"

func main() {

	x := [5]float64{
		98,
		93,
		77,
		82,
		83,
	}

	var total float64 = 0
	for _, value := range x {
		total += value
	}
	fmt.Println(total / float64(len(x)))

	fmt.Println(x[3])

	elements := make(map[string]string)
	elements["H"] = "Hydrogen"
	elements["He"] = "Helium"
	elements["Li"] = "Lithium"
	elements["Be"] = "Beryllium"
	elements["B"] = "Boron"
	elements["C"] = "Carbon"
	elements["N"] = "Nitrogen"
	elements["O"] = "Oxygen"
	elements["F"] = "Fluorine"
	elements["Ne"] = "Neon"

	fmt.Println(elements["Li"])

	presAge := make(map[string]int)
	presAge["Roosevelt"] = 42
	presAge["Kennedy"] = 37
	fmt.Println(presAge["Roosevelt"])
	delete(presAge, "Kennedy")
}
