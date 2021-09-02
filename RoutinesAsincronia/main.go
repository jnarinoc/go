package main

import (
	"fmt"
	"strings"
	"time"
)

func main() {
	// en este caso la ejecucion no espera a que se termine
	go miNombreLento("Felipe Nariño")
	fmt.Println("estoy aqui")
	var x string
	fmt.Scanln(&x)
}

func miNombreLento(nombre string) {
	letras := strings.Split(nombre, "")
	for _, letra := range letras {
		time.Sleep(1000 * time.Millisecond)
		fmt.Println(letra)
	}
}
