package main

import (
	"fmt"
	"strings"
)

type pessoa struct {
	nome      string
	sobrenome string
}

func (p pessoa) getFullName() string {
	return p.nome + " " + p.sobrenome
}

func (p *pessoa) setFullName(fullname string) {
	partes := strings.Split(fullname, " ")
	p.nome = partes[0]
	p.sobrenome = partes[1]
}

func main() {

	p1 := pessoa{"Wellington", "Defassio"}

	fmt.Println(p1.getFullName())

	p1.setFullName("Mario Armario")

	fmt.Println(p1.getFullName())

}
