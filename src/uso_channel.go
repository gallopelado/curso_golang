package main

import "fmt"

// <- solo entrada de datos del lado derecho, chan<-
// <- solo de salida de datos del lado izquierdo, <-chan
func say(text string, c chan<- string) {
	// asignar el texto ingresado al channel
	c <- text
}

func main() {
	// Asigna 1 goroutine, se puede dejar vacío pero se recomienda dejar un valor por buena práctica
	c := make(chan string, 1)
	fmt.Println("Hello")
	go say("Bye", c)
	// extraer la salida del canal
	fmt.Println(<-c)
}
