package main

import (
	"busquedas"
	"fmt"
	"time"
	"utiles"
)

func main() {
	tamaño := []int{1000, 10000, 100000, 1000000, 10000000, 100000000}
	for i := 0; i < len(tamaño); i++ {
		fmt.Println("Pruebas con tamaño:", tamaño[i])
		datos(tamaño[i])
	}
}
func datos(valorref int) {
	arreglo := utiles.GenerarArreglo(10, valorref)
	buscado := -1

	//fmt.Println(arreglo)

	inicio := time.Now()
	// Busqueda Lineal
	fmt.Println(busquedas.BusLineal(arreglo, buscado))
	fmt.Println("Busqueda Lineal: ", time.Since(inicio))

	inicio = time.Now()
	// Ordenar el arreglo para la busqueda binaria
	//sort.Ints(arreglo)
	busquedas.Burbujeo(arreglo)
	fmt.Println("Ordenamiento: ", time.Since(inicio))
	//fmt.Println(arreglo)

	inicio = time.Now()
	// Busqueda Binaria
	fmt.Println(busquedas.BusquedaBinaria(arreglo, buscado))
	fmt.Println("Busqueda Binaria: ", time.Since(inicio))

}
