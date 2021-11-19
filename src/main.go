package main

import (
	"fmt"
	"math"
)

func main() {
	// Declaración de constantes
	const pi float64 = 3.14
	const pi2 = 3.1415

	fmt.Println("pi:", pi)
	fmt.Println("pi2:", pi2)

	// Declaracion de variables enteras
	base := 12
	var altura int = 14
	var area int

	fmt.Println(base, altura, area)

	// Zero values
	var a int
	var b float64
	var c string
	var d bool

	fmt.Println(a, b, c, d)

	// Area del cuadrado
	const baseCuadrado = 10
	areaCuadrado := baseCuadrado * baseCuadrado
	fmt.Println("Area cuadrado:", areaCuadrado)

	// Operadores matemáticos
	x := 10
	y := 50

	//Suma
	resultado := x + y
	fmt.Println("Suma:", resultado)

	// Resta
	resultado = y - x
	fmt.Println("Resta:", resultado)

	// Multiplicación
	resultado = y * x
	fmt.Println("Multiplicacion:", resultado)

	// Division
	resultado = y / x
	fmt.Println("Division:", resultado)

	// Modulo o residuo
	resultado = y % x
	fmt.Println("Modulo o residuo:", resultado)

	// Incremental
	x++
	fmt.Println("Incremental:", x)

	// Decremental
	x--
	fmt.Println("Decremental:", x)

	// Retos
	// Area del rectangulo, trapecio, circulo
	// Rectángulo A = b x h
	baseR := 10
	hR := 6
	areaRectangulo := baseR * hR
	fmt.Println("Area del rectangulo:", areaRectangulo)

	// Area del trapecio: El área del trapecio es igual a la suma de las bases por la altura, y dividido por dos.
	baseMayor := 10
	baseMenor := 4
	hTrapecio := 4
	areaTrapecio := ((baseMayor + baseMenor) * hTrapecio) / 2
	fmt.Println("Area del trapecio:", areaTrapecio)

	// Area del círculo: A = PI x r(cuadrado)
	areaCirculo := math.Pow(3.14159265358979323846264338327950288419716939937510582097494459, 2)

	fmt.Println("Area del circulo:", areaCirculo)
}
