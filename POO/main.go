package main

import (
	"fmt"
	"time"

	us "./user"
)

type pepe struct {
	us.Usuario
}

func main() {
	u := new(pepe)
	u.AltaUsuario(1, "Felipe", time.Now(), true)
	fmt.Println(u.Usuario)
}

// correr con GO111MODULE=off go run main.go
