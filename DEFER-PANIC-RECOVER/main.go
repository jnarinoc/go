package main

import (
	"fmt"
	"log"
	"os"
)

func main() {
	archivo := "prueba.txt"
	f, err := os.Open(archivo)
	// defer no se ejecuta secuencialmente, solo lo hace al final
	defer f.Close()
	if err != nil {
		fmt.Println("error abriendo el archivo")
		// os.Exit(1)
	}

	ejemploPanic()
}

func ejemploPanic() {
	// defer ejecuta una sola cosa por eso, si quiero mas instrucciones debo crear una funcion anonima
	defer func() {
		// con recover capturo el panic
		reco := recover()
		if reco != nil {
			log.Fatalf("Ocurrio un error generado por panic %v", reco)
		}
	}()
	a := 1
	if a == 1 {
		// forzar error
		panic("Se encontro valor 1")
	}
}
