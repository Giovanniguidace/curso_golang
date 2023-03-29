package main

import "fmt"

func main() {
	s := make([]int, 10)
	s[9] = 12
	fmt.Println(s)

	// Criando um slice tamahno 10 e um array tamanho 20
	s = make([]int, 10, 20)
	fmt.Println(s, len(s), cap(s))

	// alterando o slice s e adicionando mais 10 valores para preencher capacidade
	s = append(s, 1, 2, 3, 4, 5, 6, 7, 8, 9, 0)
	fmt.Println(s, len(s), cap(s))

	// a capacidade do slice é dobrada casoo limite é atingido e queremos aidicionar mais itens
	s = append(s, 1)
	fmt.Println(s, len(s), cap(s))

}
