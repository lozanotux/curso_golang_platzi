package main

import "fmt"

func main() {
	// Defer: ejecuta el codigo cuando el programa esta por terminar
	//        por ejemplo: funcion para cerrar base de datos.
	defer fmt.Println("Hola")
	fmt.Println("Mundo")

	// Continue y break (por lo general se usan dentro de for)
	for i := 0; i < 10; i++ {
		fmt.Println(i)

		// Continue
		if i == 2 {
			fmt.Println("Es 2")
			continue
		}

		// Break
		if i == 8 {
			fmt.Println("Break")
			break
		}
	}
}
