package main

import (
	"custom"
	"fmt"
)

func encaminhar(origem <-chan string, destino chan<- string) {
	for {
		destino <- <-origem
	}
}
func juntar(entrada1, entrada2 <-chan string) <-chan string {
	c := make(chan string)
	go encaminhar(entrada1, c)
	go encaminhar(entrada2, c)
	return c
}
func main() {
	channel := juntar(
		custom.Titulo("https://www.cod3r.com.br", "https://www.google.com.br"),
		custom.Titulo("https://www.amazon.com.br", "https://www.youtube.com.br"),
	)
	fmt.Println(<-channel, "|", <-channel)
	fmt.Println(<-channel, "|", <-channel)
}
