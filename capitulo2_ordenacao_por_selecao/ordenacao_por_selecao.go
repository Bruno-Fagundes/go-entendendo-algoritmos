package main

import (
	"fmt"
)

func main() {

	minhaLista := []int{21, 43, 55, 6, 01}
	listaOrdenada := ordenarPorSelecao(minhaLista)
	fmt.Println(listaOrdenada)
}

func buscaMenor(arr []int) int {

	menor := arr[0]
	menorIndice := 0

	for i := 1; i < len(arr); i++ {
		if arr[i] < menor {
			menor = arr[i]
			menorIndice = i
		}
	}

	return menorIndice
}

func ordenarPorSelecao(arr []int) []int {
	
	arrCopia := make([]int, len(arr))
	copy(arrCopia, arr)

	novoArr := []int{}
	
	for len(arrCopia) > 0 { 

		menorIndice := buscaMenor(arrCopia) 
		menorElemento := arrCopia[menorIndice] 
		arrCopia = append(arrCopia[:menorIndice], arrCopia[menorIndice+1:]...)
		novoArr = append(novoArr, menorElemento)
	}

	return novoArr
}
