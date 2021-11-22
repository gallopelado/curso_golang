package main

import "fmt"

func hola() {
	fmt.Println("Hola mundo")
}

func adios() {
	fmt.Println("Adios")
}

func hablar() {
	fmt.Println("Bla bla bla")
}

func saludo(nombre string) {
	fmt.Println("Hola " + nombre)
}

func sumar(num1, num2 int) int {
	return num1 + num2
}

func frutas(fruta1, fruta2 string) (a, b string) {
	return fruta1, fruta2
}

func main() {
	hola()
	saludo("Juan")
	hablar()
	adios()
	fmt.Println(sumar(1, 1))
	fmt.Println(frutas("banana", "pera"))

	// si quiero solamente uno de los dos returns de frutas
	fruta1, fruta2 := frutas("banana", "pera")

	fmt.Println(fruta1, fruta2)

	fruta1, _ = frutas("banana", "pera")
}
