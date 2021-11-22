package main

import "fmt"

func separador() {
	fmt.Println("*****************************")
}

func main() {

	// primer tipo de for
	for i := 1; i <= 10; i++ {
		fmt.Println(i)
	}

	separador()

	// segundo tipo
	contador := 1
	for contador <= 10 {
		fmt.Println(contador)
		contador++
	}

	// tercer tipo, el forever, cuidado aquí
	//for {
	// bloque de codigo a ejecutarse por siempre
	//}

	separador()

	// cuarto tipo
	listaNumerosPares := []int{2, 4, 6, 8, 10, 12, 14, 16, 18, 20}
	for i, par := range listaNumerosPares {
		fmt.Printf("posicion %d, número par: %d\n", i, par)
	}
}
