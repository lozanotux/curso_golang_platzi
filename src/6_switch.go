package main

import "fmt"

func main() {
	switch modulo := 4 % 2; modulo {
	case 0:
		fmt.Println("Is pair")
	default:
		fmt.Println("Is odd")
	}

	// Sin condicion
	valor := 200
	switch {
	case valor > 100:
		fmt.Println("Is greater than 100")
	case valor < 0:
		fmt.Println("Is less than 0")
	default:
		fmt.Println("No condition")
	}
}
