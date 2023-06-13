package main

import (
	"fmt"
	"reflect"
)

func main() {
	x := 1000
	fmt.Println("Literal de inteiro é ", reflect.TypeOf(x))

	// valores sem sinal (só positivos).... unit8 16 32 e 64

	const xunit uint32 = 10000

	s1 := "oi você é um teste de tamanho frase.."

	fmt.Println(len(s1))

}
