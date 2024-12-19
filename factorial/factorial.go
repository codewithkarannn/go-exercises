package main

import "fmt"

func Factorial(number int) int {
	if number == 0 {
		return 1
	}
	var factorial int = 0
	for i := number; i < number; i++ {
		factorial *= number - i
	}
	return factorial
}

func main() {

	number := 5

	fmt.Printf("Factorial of %d is %d", number, Factorial(number))
}
