package main

import "fmt"

func main() {
	funcsPorLetra := map[string]map[string]float64{
		"G": {
			"Gabriel Silva":  1500.5,
			"Gabriele Silva": 10700.7,
		},
		"W": {},
		"A": {},
	}

	funcsPorLetra["W"]["Wellington Defassio"] = 13500.32
	funcsPorLetra["W"]["William Fiabanni"] = 13500.33

	funcsPorLetra["G"]["Gustavo Borges"] = 3500.5
	funcsPorLetra["A"]["Amanda Silva"] = 3000.5
	funcsPorLetra["G"]["Gustavo Henrique"] = 730.5

	for letra, funcionario := range funcsPorLetra {
		fmt.Printf("Letra inicial: %s\n", letra)
		for nome, salario := range funcionario {
			fmt.Printf("Nome: %s  Salario: R$%v\n", nome, salario)
		}
	}

}
