package main

import (
	"fmt"
	"strconv"
)

var numero int
var texto string
var status bool

func main() {
	// Go las inicializa automaticamente en 0
	var num1, num2 int
	num3, num4 := 2, "Texto"
	num5, num6 := 56, 33
	fmt.Println("Hello World")
	numer := 3
	fmt.Println(num1)
	fmt.Println(num2)
	fmt.Println(num3)
	fmt.Println(num4)
	fmt.Println(num5)
	fmt.Println(num6)
	fmt.Println(numer)
	// Go inicializa en false
	fmt.Println(status)
	mostrarStatus()

	// strconv contiene funciones para manipular tipos de datos
	texto = strconv.Itoa(num3)
	fmt.Println(texto)
}

func mostrarStatus() {
	fmt.Println(status)
}
