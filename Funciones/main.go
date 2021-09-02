package main

import (
	"fmt"
)

func main() {
	fmt.Println(uno(5))
	numero, estado := dos(1)
	fmt.Println(numero, estado)
	fmt.Println(tres())
	fmt.Println(calculo(1, 2, 3))
}

func uno(numero int) int {
	return numero * 2
}

func dos(numero int) (int, bool) {
	if numero == 1 {
		return 5, false
	}
	return 3, true
}

// al hacer esto, retorna la variable z sin necesidar se especificarla
func tres() (z int) {
	z = 2
	return
}

// _ se usa para variables que no voy a usar
func calculo(numero ...int) int {
	total := 0
	for _, num := range numero {
		total = total + num
	}
	return total
}
