package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
)

func main() {
	var lista []int = []int{1, 3, 5, 7, 9}

	//fmt.Println(pesquisa_binaria(lista, 3))
	if encontrado := pesquisa_binaria(lista, 3); encontrado != -1 {
		fmt.Println("O número 3 foi encontrado na lista no índice", encontrado)
	} else {
		fmt.Println("O número 3 não foi encontrado na lista")
	}

	//fmt.Println(pesquisa_binaria(lista, -1))
	if encontrado := pesquisa_binaria(lista, -1); encontrado != -1 {
		fmt.Println("O número -1 foi encontrado na lista no índice", encontrado)
	} else {
		fmt.Println("O número -1 não foi encontrado na lista")
	}

	// Leitura e Conversão de valores int64 para int, que será aceito pela função
	scanner := bufio.NewScanner(os.Stdin)
	fmt.Print("Digite um número para pesquisar na lista: ")
	scanner.Scan()
	input := scanner.Text()

	itemProcuradoInt64, err := strconv.ParseInt(input, 10, 64)
	if err != nil {
		fmt.Println("Erro ao converter a entrada para inteiro:", err)
		return
	}

	itemProcuradoInt := int(itemProcuradoInt64)

	if encontrado := pesquisa_binaria(lista, itemProcuradoInt); encontrado != -1 {
		fmt.Println("O número", itemProcuradoInt, "foi encontrado na lista no índice", encontrado, ".")
	} else {
		fmt.Println("O número", itemProcuradoInt, "não foi encontrado na lista")
	}
}

func pesquisa_binaria(lista []int, item int) int {
	baixo := 0
	alto := len(lista) - 1

	for baixo <= alto {
		meio := (baixo + alto) / 2
		chute := lista[meio]
		if chute == item {
			return meio // Retorna o índice do item encontrado
		}
		if chute > item {
			alto = meio - 1
		} else {
			baixo = meio + 1
		}
	}
	return -1 // Retorno de um valor inváliido
}
