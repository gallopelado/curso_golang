package main

import "fmt"

type pc struct {
	ram   int
	disk  int
	brand string
}

func (myPC pc) ping() {
	fmt.Println(myPC.brand, "Pong")
}

func (myPC *pc) duplicateRAM() {
	myPC.ram = myPC.ram * 2
}

func main() {
	a := 50
	b := &a

	fmt.Println("Valor de a:", a)
	fmt.Println("Dirección de memoria de a en b:", b)
	fmt.Println("Valor de la Dirección de memoria de a en b:", *b)

	/*
		Si modificas el valor de la dirección en memoria, las demas variables que apuntan a esa dirección van a cambiar
	*/

	*b = 100
	fmt.Printf("El valor de a ahora es %d, ha cambiado porque su valor apuntaba a b\n", a)

	myPC := pc{ram: 16, disk: 200, brand: "msi"}
	fmt.Println(myPC)

	myPC.ping()

	fmt.Println(myPC)
	myPC.duplicateRAM()

	fmt.Println(myPC)
	myPC.duplicateRAM()
	fmt.Println(myPC)
}
