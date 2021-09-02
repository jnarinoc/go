package main

import (
	"fmt"
)

func main() {

	for i := 1; i <= 10; i++ {
		fmt.Println(i)
	}

	numero := 0
	for {
		fmt.Println("Continuo")
		fmt.Println("Ingrese numero secreto")
		fmt.Scanln(&numero)
		if numero == 100 {
			break
		}
	}

	var i = 0
	for i < 10 {
		fmt.Printf(" Valor de i: %d", i)
		if i == 5 {
			fmt.Println(" multiplicamos por 2 ")
			i = i * 2
			continue
		}
		fmt.Println(" Paso por aqui")
		i++
	}

	var x int = 0
RUTINA:
	for x < 10 {
		if x == 4 {
			x = x + 2
			fmt.Println("Voy a rutina sumando 2 a x")
			goto RUTINA
		}
		fmt.Printf("Valor de x: %d", x)
		x++
	}
}
