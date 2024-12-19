package main

import "fmt"

func IsPrime(number int) bool {

	if number < 1 {
		return false
	}
	for i := 2; i*i <= number; i++ {
		if number%i == 0 {
			return false
		}

	}

	return true
}

func main() {

	number := 9

	if IsPrime(number) {
		fmt.Printf("%d is a prime number.\n", number)
	} else {
		fmt.Printf("%d is not a prime number.\n", number)
	}

}
