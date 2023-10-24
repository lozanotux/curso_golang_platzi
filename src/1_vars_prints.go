package main

import "fmt"

func main() {
	// Declaración de constantes
	const pi float64 = 3.14
	const PI2 = 3.1415
	const MESSAGE = "Un valor"

	fmt.Println("pi:", pi)
	fmt.Println("pi2:", PI2)
	fmt.Println("CONSTANTE:", MESSAGE)

	// Declaración de variables enteras
	algo := 12
	var algo2 int = 14
	var algo3 = 25

	fmt.Println(algo, algo2, algo3)

	// Zero values (default values initialization)
	var a int
	var b float64
	var c string
	var d bool

	fmt.Println(a, b, c, d)

	// Calcular area de un cuadrado
	const baseCuadrado = 10
	areaCuadrado := baseCuadrado * baseCuadrado
	fmt.Println("Area del cuadrado:", areaCuadrado)

	// Operadores Aritmeticos
	x := 10
	y := 50

	// Suma
	result := x + y
	fmt.Println("Suma:", result)

	// Resta
	result = y - x
	fmt.Println("Resta:", result)

	// Multiplicación
	result = x * y
	fmt.Println("Multiplicación:", result)

	// División
	result = y / x
	fmt.Println("División:", result)

	// Modulo
	result = y % x
	fmt.Println("Modulo:", result)

	// Incremental
	x++
	fmt.Println("Incremental:", x)

	// Decremental
	x--
	fmt.Println("Decremental:", x)

	// Retos
	// Area de un rectangulo
	base := 20
	altura := 10
	fmt.Println("Area Rectangulo:", base*altura)

	// Area de un trapecio
	aa := 10
	bb := 20
	hh := 8
	fmt.Println("Area Trapecio:", hh*((aa+bb)/2))

	// Area de un circulo
	rr := 4
	fmt.Println("Area Circulo:", pi*float64(rr*rr))

	// fmt Package
	// Println
	helloMessage := "Hello"
	worldMessage := "world"
	fmt.Println(helloMessage, worldMessage)
	fmt.Println(helloMessage, worldMessage)

	// Printf
	nombre := "Platzi"
	cursos := 500
	fmt.Printf("%s tiene mas de %d cursos\n", nombre, cursos)
	fmt.Printf("%v tiene mas de %v cursos\n", nombre, cursos)

	// Sprintf
	message := fmt.Sprintf("%s tiene mas de %d cursos", nombre, cursos)
	fmt.Println(message)

	// Tipo de datos
	fmt.Printf("helloMessage: %T\n", helloMessage)
	fmt.Printf("cursos: %T\n", cursos)

}
