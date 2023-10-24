package main

import (
	pp "curso_golang_platzi/src/paquete_puntero"
	"fmt"
)

type pc struct {
	ram   int
	disk  int
	brand string
}

func (miPC pc) pnig() {
	fmt.Println(miPC.brand, "Pong")
}

func (miPC *pc) duplicateRAM() {
	miPC.ram = miPC.ram * 2
}

func main() {
	a := 50
	b := &a

	fmt.Println(b)
	fmt.Println(*b)

	*b = 100
	fmt.Println(a)

	miPc := pc{ram: 16, disk: 200, brand: "msi"}
	fmt.Println(miPc)

	miPc.pnig()

	fmt.Println(miPc)
	miPc.duplicateRAM()

	fmt.Println(miPc)
	miPc.duplicateRAM()

	fmt.Println(miPc)

	pc1 := pp.PC{Serial: "ABC123", Marca: "Acer", Modelo: "D255E", Ram: 8, Disco: 120}
	pc1.UpdateRam(8)
	fmt.Println(pc1)
}
