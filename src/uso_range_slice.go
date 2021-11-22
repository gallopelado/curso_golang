package main

import "fmt"

func main() {
	frutas := []string{"naranja", "manzana", "pera"}
	for i, valor := range frutas {
		fmt.Println(i, valor)
	}

	fmt.Println("Solo valor")
	for _, valor := range frutas {
		fmt.Println(valor)
	}

	fmt.Println("Solo índice")
	for i := range frutas {
		fmt.Println(i)
	}
}
