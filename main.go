package main

import "fmt"

// Porgrama para encontrar un numero con la recursion

func BuscarNumero(array []int, elem int, i int, f int) int {
	op := (i + f) / 2
	//fmt.Println("i and f and op: ", i, f, op)
	if array[op] == elem {
		fmt.Println(array[op])
		return array[op]
	}

	if len(array) <= 1 {
		fmt.Println("No se encontro elk numero")
		return 2
	}

	fmt.Println("array: ", array)
	fmt.Println("i f", i, f)
	fmt.Println("array[op] array: ", array[op], array)
	if array[op] > elem {
		store := make([]int, 0)
		store = append(store, array[0:op]...)
		return BuscarNumero(store, elem, i, len(store)-1) // Con 5 este sera el retorno principal
	}

	if array[op] < elem {
		store := make([]int, 0)
		store = append(store, array[op:]...)
		op = 0
		fmt.Println(store)
		return BuscarNumero(store, elem, op, len(store))
	}

	return 1

}

func main() {
	array := []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10, 11, 12, 13, 14, 15, 16, 17}
	xd := BuscarNumero(array, 17, 0, len(array)-1)
	fmt.Println("Numero encontrado: ", xd)
}
