package main

import "fmt"

func main() {

	type Pessoa struct {
		Nome  string
		Idade int
	}

	p1 := Pessoa{Nome: "João", Idade: 16}
	p2 := Pessoa{Nome: "João", Idade: 17}

	fmt.Println(p1 == p2)
}
