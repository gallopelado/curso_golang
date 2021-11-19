package main

import "fmt"

func main() {
	// Declaración de variables
	helloMessage := "Hello"
	worldMessage := "world"

	// Println --impresion en linea
	fmt.Println(helloMessage, worldMessage)

	// Printf
	nombre := "Juan José"
	anios := 29
	// si no sabemos que tipo de variable sera %v
	fmt.Printf("%s tiene %d años\n", nombre, anios)
	fmt.Printf("%v tiene %v años\n", nombre, anios)

	// Sprintf --guarda el valor a la variable
	message := fmt.Sprintf("%s tiene %d años\n", nombre, anios)
	fmt.Println(message)

	// Tipo de datos
	fmt.Printf("helloMessage es de tipo: %T\n", helloMessage)
}
