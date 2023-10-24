package main

import "fmt"

type auto struct {
	marca  string
	modelo int
}

func main() {
	miAuto := auto{marca: "Ford", modelo: 2020}
	fmt.Println(miAuto)

	// Otra manera
	var otroAuto auto
	otroAuto.marca = "Ferrari"
	fmt.Println(otroAuto)
}
