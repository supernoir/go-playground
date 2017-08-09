package main

import "fmt"

func main() {
	rect1 := Rectangle{leftX: 0, topY: 50, height: 10, width: 10}
	fmt.Println("Rectangle is", rect1.width, "wide")
	fmt.Println("Area of Rectangle is", rect1.area())
}

type Rectangle struct {
	leftX  float64
	topY   float64
	height float64
	width  float64
}

func (rect *Rectangle) area() float64 {
	return rect.width * rect.height
}
