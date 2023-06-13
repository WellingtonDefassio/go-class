package main

import "fmt"

func main() {
	fmt.Print("Mesma")
	fmt.Print(" Linha. ")
	fmt.Println(" Nova ")

	x := 3.141516

	fmt.Println("0 de x é", x)

	xs := fmt.Sprint(x)

	fmt.Println("0 de x é " + xs)

}
