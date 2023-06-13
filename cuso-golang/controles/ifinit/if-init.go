package main

import (
	"math/rand"
	"time"
)

func numeroAleatorio() int {
	numeroRando := rand.NewSource(time.Now().UnixNano())
	r := rand.New(numeroRando)
	return r.Intn(10)
}

func main() {

	if i := numeroAleatorio(); i > 5 {
		println("Ganhou!!")
	} else {
		println("Perdeu!!")
	}

}
