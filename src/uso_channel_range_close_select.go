package main

import (
	"fmt"
)

/*
	Uso de Channels con Range, Close y Select
*/

func message(text string, c chan string) {
	c <- text
}

func main() {
	// Dos datos en simultáneo
	c := make(chan string, 2)
	c <- "Mensaje1"
	c <- "Mensaje2"
	// len para ver cuantas datos goroutines hay en ese channel
	// cap para ver la cantidad maxima que puede almacenar ese channel
	fmt.Println(len(c), cap(c))

	/*
		Puedes comentar una de las partes donde se asigna el valor a c para ver la cuanto se está usando y la capacidad máxima
	*/

	// Range y close
	// cierra el canal, ese canal no va recibir otro dato adicional, es ideal cerrar siempre el canal
	close(c)

	// Esto dará error porque supera la capacidad máxima
	//c <- "Mensaje3"

	// Range Nos ayuda a hacer un recorrido de la lista(No olvides descomentar el close)
	for message := range c {
		fmt.Println(message)
	}

	// Select
	email1 := make(chan string)
	email2 := make(chan string)
	// Enviar información a los canales
	go message("mensaje1", email1)
	go message("mensaje2", email2)
	for i := 0; i < 2; i++ {
		select {
		case m1 := <-email1:
			fmt.Println("Email recibido de email1", m1)
		case m2 := <-email2:
			fmt.Println("Email recbido de email2", m2)
		}
	}
}
