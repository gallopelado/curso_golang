package mypackage

import "fmt"

// PersonPublic es persona con acceso público
type PersonPublic struct {
	FirstName string
	LastName  string
	Age       int
	Phone     string
	Document  string
}

//funcion que muestra un mensaje en consola
func PrintMessage() {
	fmt.Println("Holis, estoy en un paquete externo")
}
