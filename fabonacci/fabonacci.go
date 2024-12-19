package main

import "fmt"

func main() {
	n := 7

	// Validate input
	if n <= 0 {
		fmt.Println("Please enter a positive integer.")
		return
	}

	// Generate and print Fibonacci series
	fmt.Println("Fibonacci Series:")
	if n == 1 {
		fmt.Println(0)
		return
	}

	a, b := 0, 1
	fmt.Print(a, ", ", b) // Print the first two terms

	for i := 3; i <= n; i++ { // Start from the 3rd term
		c := a + b
		fmt.Print(", ", c)
		a, b = b, c
	}

	fmt.Println()
}
