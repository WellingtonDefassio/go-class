package main

import (
	"fmt"
	"time"
)

func fale(pessoa string, text string, qtde int) {
	for i := 0; i < qtde; i++ {
		time.Sleep(time.Second)
		fmt.Printf("%s: %s (iretação %d)\n", pessoa, text, i+1)
	}
}

func main() {
	//fale("Maria", "Pq vc não fala comigo", 3)
	//fale("João", "Só posso falar depois de você", 1)

	//go fale("Maria", "Ei...", 500)
	//go fale("João", "Opa...", 500)
	//time.Sleep(time.Second * 5)
	//fmt.Println("Fim")

	go fale("Maria", "Entendi!!!", 10)
	fale("João", "Parabéns!!", 5)

}
