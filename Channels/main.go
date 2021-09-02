package main

import (
	"fmt"
	"time"
)

// UN CANAL ES UN espacio de memoria para dialogo entre rutinas

func main() {
	canal1 := make(chan time.Duration)
	go bucle(canal1)
	fmt.Println("llegue hasta aca")
	// para poner a alguien a la espera de que la rutina ha terminado
	//espera a que canal1 tenga valor (parece promesas / async await)
	msg := <-canal1
	fmt.Println(msg)
}

func bucle(canal1 chan time.Duration) {
	inicio := time.Now()
	for i := 0; i < 100000000000; i++ {

	}
	final := time.Now()
	// asignarle valor al canal
	canal1 <- final.Sub(inicio)
}
