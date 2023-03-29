package main

import "fmt"

func main() {
	numeros := [...]int{1, 2, 3, 4, 5} // compilador realiza a contagem do tamanho do array

	// Para percorrer o array mostrando o index e o seu conteúdo
	for i, numero := range numeros {
		fmt.Printf("index: %d) valor: %d\n", i, numero)
	}

	// Para percorrer o array e se preocupar apenas com o conteúdo, ignorando seu index.
	for _, num := range numeros {
		fmt.Println(num)
	}

}
