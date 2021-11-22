package main

import "fmt"

func main() {

	persona := make(map[string]string)

	persona["1616"] = "Celina"
	persona["123"] = "Juan"

	fmt.Println(persona)

	// Recorrer un map
	for i, v := range persona {
		fmt.Println(i, v)
	}

	// encontrar un valor
	value, ok := persona["1616"]
	fmt.Println(value, ok)

	/*
		Un detalle interesante de los maps es que trabajan de forma concurrente, esto los hace más eficientes, también ende
		van a agregarse de forma desordenada. Si se desea trabajar con cierto ordenamiento
		mejor utilizar slice
	*/
}
