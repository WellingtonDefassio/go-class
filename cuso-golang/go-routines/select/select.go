package main

import (
	"custom"
	"fmt"
	"time"
)

func oMaisRapido(url1, url2, url3 string) string {
	titulo := custom.Titulo(url1)
	titulo2 := custom.Titulo(url2)
	titulo3 := custom.Titulo(url3)
	select {
	case t1 := <-titulo:
		return t1
	case t2 := <-titulo2:
		return t2
	case t3 := <-titulo3:
		return t3
	case <-time.After(1000 * time.Millisecond):
		return "Todos perderam"
		//default:
		//	return "Sem resposta"

	}

}

func main() {
	rapido := oMaisRapido("https://www.cod3r.com.br", "https://www.google.com.br", "https://www.amazon.com.br")
	fmt.Println(rapido)
}
