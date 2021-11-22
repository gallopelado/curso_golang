package main

import "fmt"

func main() {
	// Antes de finalizar la función se ejecuta defer, usar con cuidado, un defer por función
	defer fmt.Println("Inicio diferido")
	fmt.Println("Sin defer, se ejecuta primero")

	// Continue y break
	for i := 1; i <= 20; i++ {
		fmt.Println(i)
		if i == 2 {
			fmt.Println("Se usó continue")
			continue
		} else if i == 8 {
			fmt.Println("Condición finalizada por break")
			break
		}

	}
}
