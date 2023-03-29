package main

import "fmt"

func main() {
	funcsESalarios := map[string]float64{
		"José João":      11325.45,
		"Gabriela Silva": 15456.78,
		"Pedro Junior":   1200.0,
	}

	funcsESalarios["Rafael Filho"] = 1350.0 // Adicionando nova chave e valor. É como um append.

	delete(funcsESalarios, "Inexistente") // Fará a procura no map e se existir a chave 'Inexistente', ele será deletado.
	// Se não houver o valor para ser deletado, o sistema continuára o processo normalmente. Não trará erros.

	for nome, salario := range funcsESalarios {
		fmt.Println(nome, salario)
	}
}
