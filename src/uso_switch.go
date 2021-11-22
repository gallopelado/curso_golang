package main

import "fmt"

func elegirFruta(fruta string) string {
	switch fruta {
	case "pera":
		return fruta
	case "manzana":
		return fruta
	default:
		return "No existe esta fruta"
	}
}

func main() {
	fmt.Println(elegirFruta("pera"))
}
