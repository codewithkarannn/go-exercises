package main

import "fmt"

func main() {
	fmt.Println("Area of the rectangle")

	const length = 10
	const breath = 5

	areaOfRectangle := 2 * (length + breath)

	println("Length of the rectangle =", length)
	println("Breath of the rectangle =", breath)
	println("Area of the rectangle =", areaOfRectangle)

}
