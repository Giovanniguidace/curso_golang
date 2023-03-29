package main

import "fmt"

func main() {
	// var aprovados map[int]string !!! Mapas devem ser inicializados
	aprovados := make(map[int]string)

	aprovados[4234234234] = "Maria"
	aprovados[4534543534] = "Pedro"
	aprovados[5464564566] = "Ana"

	fmt.Println(aprovados)

	for cpf, nome := range aprovados {
		fmt.Printf("%s (CPF: %d)\n", nome, cpf)
	}

	fmt.Println(aprovados[4234234234])
	delete(aprovados, 4234234234)
	fmt.Println(aprovados[4234234234])

}
