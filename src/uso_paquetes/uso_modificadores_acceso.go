package main

import (
	"fmt"
	"uso_paquetes/mypackage"
	pk "uso_paquetes/mypackage"
)

/*
	No existe una keyword especial para tratar estos modificadores
	por convención, si la primera letra es mayúscula es pública,
	si es minúscula es privado
*/

/*
	Que dolor de testículos es trabajar con paquetes, tuve que mover muchas cosas.
	generar el go.mod, es parecido al init.py en Python, veremos con otros paquetes
	de otras ubicaciones
	Use dentro del paquete este comando
	go mod init uso_paquetes

	Si lo haces desde la raiz curso_golang vas a generar un flor de error, ¿Por qué?
	Porque en los otros ficheros estamos usando el main y esto hace que detecte varias veces
	marcando en color rojo todos. Que pendejo :)
*/

func main() {

	// 1 forma
	p1 := mypackage.PersonPublic{
		FirstName: "Juan Jose",
		LastName:  "Gonzalez Ramirez",
		Age:       29,
		Phone:     "123456",
		Document:  "1111111",
	}
	fmt.Println(p1)

	// 2 forma
	var p2 pk.PersonPublic
	p2.FirstName = "Sandra"
	p2.LastName = "Lopez"
	p2.Age = 30
	p2.Phone = "11223366"
	p2.Document = "99643183"

	fmt.Println(p2)

	pk.PrintMessage()
}
