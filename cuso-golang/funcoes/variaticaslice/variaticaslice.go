package main

import "fmt"

func imprimirAprovados(aprovados ...string) {

	for _, aprovado := range aprovados {
		fmt.Println(aprovado)
	}

}

func main() {
	aprovados := []string{"Maria", "Silvio", "Santos"}
	imprimirAprovados(aprovados...)
	aprovados2 := append(aprovados, "Jose", "Joe", "Marcia")
	imprimirAprovados(aprovados2...)
}
