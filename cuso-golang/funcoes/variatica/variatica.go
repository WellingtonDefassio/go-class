package main

import "fmt"

func media(numeros ...float64) float64 {
	total := 0.0
	for _, numero := range numeros {
		total += numero
	}
	return total / float64(len(numeros))
}

func main() {

	f := media(5.7, 8.5, 3.9)
	fmt.Println(f)
}
