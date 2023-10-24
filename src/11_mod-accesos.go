package main

import (
	// opcion1 -->	"curso_golang_platzi/src/paquete1"
	pk "curso_golang_platzi/src/paquete1"
	"fmt"
)

func main() {
	var auto1 pk.AutoPublic
	auto1.Marca = "Ferrari"
	auto1.Modelo = 2020
	fmt.Println(auto1)

	pk.PrintMessage("Hola Platzi")
}
