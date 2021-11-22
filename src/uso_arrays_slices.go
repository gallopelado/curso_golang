package main

import "fmt"

func main() {
	// Array
	var array [4]int
	fmt.Println("Este array no tiene elementos ", array)
	// Su naturaleza es la inmutabilidad, no podemos añadir nuevos elementos
	// Pero es posible editarlos
	array[0] = 1000
	array[1] = 2000
	array[2] = 3000
	array[3] = 4000
	fmt.Println("Elementos editados del array ", array)
	fmt.Println("Tamaño del array ", len(array))
	fmt.Println("Capacidad de elementos del array ", cap(array))

	fmt.Println("****************************")

	// Slices(Mutables)
	slice := []string{"Tomate", "Cebolla", "Lechuga", "Zanahoria", "Patatas"}
	fmt.Println(slice)
	fmt.Println("Elementos editados del slice ", slice)
	fmt.Println("Tamaño del slice ", len(slice))
	fmt.Println("Capacidad de elementos del slice ", cap(slice))

	fmt.Println("****************************")
	fmt.Println("metodos del slice")
	// metodos del slice
	fmt.Println(slice[0])
	fmt.Println(slice[:3])  // el 3 es exclusivo, no va traer Zanahoria
	fmt.Println(slice[2:4]) // el 4 es exclusivo, no va traer Patatas, a partir de 2 si porque es inclusivo
	fmt.Println(slice[4:])  // el 4 inclusivo, Patatas

	// usos del append
	fmt.Println("****************************")
	fmt.Println("usos del append")
	fmt.Println("Agregar rábanos")
	//Agregar un elemento
	slice = append(slice, "Rábanos")
	fmt.Println(slice)

	// Agregar una lista completa a otra lista
	fmt.Println("Agregar otra lista")
	nuevoSlice := []string{"Pepino", "Remolacha"}
	slice = append(slice, nuevoSlice...)
	fmt.Println(slice)
}
