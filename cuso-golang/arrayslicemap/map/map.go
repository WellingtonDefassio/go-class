package main

import "fmt"

func main() {

	aprovados := make(map[int]string)

	aprovados[123456789] = "Maria"
	aprovados[564897988] = "João"
	aprovados[889666666] = "Paulo"

	fmt.Println(aprovados)

	for cpf, nome := range aprovados {

		fmt.Printf("os aprovados são nome: %s (CPF: %d)\n", nome, cpf)
	}

	println(aprovados[123456789])

	delete(aprovados, 123456789)

	fmt.Println(aprovados)

}
