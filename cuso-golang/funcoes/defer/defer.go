package main

import "fmt"

func obterValorAprovado(numero int) int {
	defer fmt.Println("Fechando a conexão com o banco de dados...")
	if numero > 5000 {
		fmt.Println("valor alto")
		fmt.Println("valor alto")
		fmt.Println("valor alto")
		return 5000
	} else {
		fmt.Println("valor baixo")
		fmt.Println("valor baixo")
		fmt.Println("valor baixo")
		return numero
	}
}

func main() {

	fmt.Println(obterValorAprovado(400))

}
