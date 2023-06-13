package main

import "fmt"

func compras(trab1 bool, trab2 bool) (bool, bool, bool, bool) {
	comprarTv50 := trab1 && trab2
	comprarTv32 := trab1 != trab2
	comprarSorvete := trab1 || trab2
	saudavel := !comprarSorvete
	return comprarTv50, comprarTv32, comprarSorvete, saudavel
}

func main() {

	tb50, tv32, sorvete, saudavel := compras(false, false)

	fmt.Printf("Tv50:%t, tv32:%t, sorvete:%t, saudavel:%t", tb50, tv32, sorvete, saudavel)

}
