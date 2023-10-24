package main

import "fmt"

func imprimirMensaje(mensaje string) {
	fmt.Println(mensaje)
}

func tripleArgumentos(a, b int, c string) {
	fmt.Println(a, b, c)
}

func devolverValorPorDos(a int) int {
	return a * 2
}

func dobleRetorno(a int) (c, d int) {
	return a, a * 2
}

func main() {
	imprimirMensaje("Hello world")
	tripleArgumentos(1, 2, "hello")

	valor := devolverValorPorDos(4)
	fmt.Println("Value:", valor)

	valor1, valor2 := dobleRetorno(3)
	fmt.Printf("Double return: %d - %d \n", valor1, valor2)

	valor3, _ := dobleRetorno(6)
	fmt.Printf("Doble return, pero solo agarro el primero: %v\n", valor3)
}
