package main

import "fmt"

type produto struct {
	name     string
	preco    float64
	desconto float64
}

func (p produto) precoComDesconto() float64 {
	return p.preco * (1 - p.desconto)
}

func main() {

	prod1 := produto{name: "Laps", preco: 10, desconto: 0.05}
	prod2 := produto{"TV", 1000, 0.2}

	fmt.Println(prod1.precoComDesconto(), prod1.name)
	fmt.Println(prod2.precoComDesconto(), prod2.name)

}
