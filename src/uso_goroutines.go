package main

import (
	"fmt"
	"sync"
	"time"
)

func say(text string, wg *sync.WaitGroup) {
	defer wg.Done()
	fmt.Println(text)
}

func main() {
	var wg sync.WaitGroup // acumular un conjunto de goroutines que se libera poco a poco

	fmt.Println("Hello") //luego este
	wg.Add(1)            // agregar 1 goroutine

	go say("world", &wg) // con el puntero del wait group

	wg.Wait() //decirle al main que espere hasta que todas las goroutines del wait group finalicen

	// funciones anónimas
	go func(text string) {
		fmt.Println(text)
	}("Adios")
	time.Sleep(time.Second * 1) // 1 segundo agregado
}
