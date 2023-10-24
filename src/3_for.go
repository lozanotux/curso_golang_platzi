package main

import "fmt"

func main() {
	// For condicional
	for i := 0; i < 10; i++ {
		fmt.Println(i)
	}

	// For while
	counter := 0
	for counter < 10 {
		fmt.Println(counter)
		counter++
	}

	// For forever
	/*counterForever := 0
	for {
		fmt.Println(counterForever)
		counterForever++
	}*/

	// For range
	lista_numeros_pares := []int{2, 4, 6, 8, 10, 12, 14, 16, 18}
	for i, numero_par := range lista_numeros_pares {
		fmt.Printf("Position [%d] --> pair number: %d \n", i, numero_par)
	}

}
