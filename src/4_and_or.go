package main

import (
	"fmt"
	"log"
	"strconv"
)

func main() {
	valor1 := 1
	valor2 := 2

	if valor1 == 1 {
		fmt.Println("Is 1")
	} else {
		fmt.Println("Is not 1")
	}

	// AND
	if valor1 == 1 && valor2 == 2 {
		fmt.Println("The AND is true")
	}

	// OR
	if valor1 == 0 || valor2 == 2 {
		fmt.Println("The OR is true")
	}

	value, err := strconv.Atoi("53")
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println(value)
}
