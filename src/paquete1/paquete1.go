package paquete1

import "fmt"

// AutoPublic auto con acceso publico
type AutoPublic struct {
	Marca  string
	Modelo int
}

type autoPrivado struct {
	marca  string
	modelo int
}

// PrintMessage imprimir un mensaje
func PrintMessage(text string) {
	fmt.Println(text)
}

func printMessage1(text string) {
	fmt.Println(text)
}

// con Mayusculas es Publico y con minusculas es privado
