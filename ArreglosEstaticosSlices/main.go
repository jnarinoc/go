package main

import (
	"fmt"
)

var tabla [10]int

func main() {
	tabla[0] = 1
	tabla[5] = 15
	fmt.Println(tabla)

	vector := [10]int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}
	fmt.Println(vector)

	for i := 0; i < len(vector); i++ {
		fmt.Println(vector[i])
	}

	//matrices
	var matriz [3][3]int
	matriz[2][2] = 1
	fmt.Println(matriz)

	//slices -> se asemejan a las listas , puedo agregar elementos
	matriz2 := []int{1, 2, 3}
	fmt.Println(matriz2)
	variante2()
}

func variante2() {
	elementos := [5]int{1, 2, 3, 4, 5}
	// desde la posicion 3 hasta el final
	porcion := elementos[3:]
	fmt.Println(porcion)
	variante3()
}

func variante3() {
	// me reserva 20 unidades de memoria pero me la crea de 5 posiciones
	elementos := make([]int, 5, 20)
	// cap -> capacidad de un slice
	fmt.Printf("Largo %d , capacidad %d", len(elementos), cap(elementos))
	variante4()
}

func variante4() {
	fmt.Println("_______________________")
	nums := make([]int, 0, 0)
	for i := 0; i < 100; i++ {
		// añade un elemento cuando sobrepasa la capacidad (cuando me quedo sin espacio)
		// go va aumentando la capacidad en base 2 a la n
		nums = append(nums, i)
		fmt.Printf("Largo %d , capacidad %d", len(nums), cap(nums))
	}
}
