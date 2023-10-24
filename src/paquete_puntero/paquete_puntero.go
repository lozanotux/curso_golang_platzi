package paquete_puntero

func (miPC *PC) UpdateRam(amount int) {
	miPC.Ram = miPC.Ram + amount
}

type PC struct {
	Serial string
	Marca  string
	Modelo string
	Ram    int
	Disco  int
}
