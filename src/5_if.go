package main

import "fmt"

func main() {
	// Defer
	defer fmt.Println("Hello")
	fmt.Println("World")

	// Break and Continue
	for i := 0; i < 10; i++ {

		fmt.Println(i)

		// Continue
		if i == 2 {
			fmt.Println("Is 2")
			continue
			fmt.Println("Is Pair")
		}

		// Break
		if i == 7 {
			fmt.Println("Break")
			break
		}
	}
}
