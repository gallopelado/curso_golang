package main

import (
	"fmt"
	"math"
)

// Retos
// Area del rectangulo, trapecio, circulo
func areaRectangulo(base, altura float64) float64 {
	return base * altura
}

// Area del trapecio: El área del trapecio es igual a la suma de las bases por la altura, y dividido por dos.
func areaTrapecio(baseMayor, baseMenor, altura float64) float64 {
	return ((baseMayor + baseMenor) * altura) / 2
}

// Area del círculo: A = PI x r(cuadrado)
func areaCirculo(radio float64) float64 {
	return math.Pow(math.Pi, 2)
}

func main() {
	area_rectangulo := areaRectangulo(2, 2)
	area_trapecio := areaTrapecio(10.4, 4.9, 50)
	area_circulo := areaCirculo(6.33)
	fmt.Println("Rectangle area is ", area_rectangulo)
	fmt.Println("Trapezium area is ", area_trapecio)
	fmt.Println("Circle area is ", area_circulo)
}
