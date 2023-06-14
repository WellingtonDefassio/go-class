package main

import "fmt"

func main() {

	var notas [3]float64

	fmt.Println(notas)

	notas[0], notas[1], notas[2] = 10.5, 15.5, 20.5

	fmt.Println(notas)

	total := 0.0
	for i, nota := range notas {

		total += nota

		i++
	}
	fmt.Println(total)

	total = 0.0
	for i := 0; i < len(notas); i++ {

		total += notas[i]

	}
	fmt.Println(total)

}
