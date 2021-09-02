package main

import (
	"fmt"
)

func main() {
	// reserva 5
	paises := make(map[string]string, 5)
	fmt.Println(paises)
	paises["Mexico"] = "D.F"
	paises["Argentina"] = "Buenos Aires"
	fmt.Println(paises["Mexico"])
	fmt.Println(paises)

	//ordena alfabeticamente por la clave
	campeonato := map[string]int{
		"Barcelona":   23,
		"Real Madrid": 24,
		"Atletico":    22}

	campeonato["Betis"] = 2
	campeonato["Real Madrid"] = 33
	fmt.Println(campeonato)
	delete(campeonato, "Betis")
	fmt.Println(campeonato)

	for equipo, puntaje := range campeonato {
		fmt.Printf("El equipo %s tiene un puntaje de : %d ", equipo, puntaje)
	}

	puntaje, existe := campeonato["Real Sociedad"]
	fmt.Printf("El puntaje es %d y el equipo existe %t", puntaje, existe)
}
