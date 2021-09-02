package main

import (
	"fmt"
)

//las funciones se pueden comportar como tipos de datos
var calculo func(int, int) int = func(num1 int, num2 int) int {
	return num1 + num2
}

func main() {
	fmt.Printf("Sumo 5 + 7 %d", calculo(5, 7))

	//al usar funciones en variables se puede redefinir
	calculo = func(num1 int, num2 int) int {
		return num1 - num2
	}

	fmt.Printf("Resto 5 - 7 %d", calculo(5, 7))

	// closures
	tablaDel := 2
	miTabla := tabla(tablaDel)
	for i := 1; i < 11; i++ {
		fmt.Println(miTabla())
	}
}

func tabla(valor int) func() int {
	numero := valor
	secuencia := 0
	return func() int {
		secuencia += 1
		return numero * secuencia
	}
}
