package main

import "fmt"

func fatorial(n uint) uint {
	switch {
	case n == 0:
		return 1
	default:
		u := n * fatorial(n-1)
		return u
	}
}

func main() {

	u := fatorial(5)

	fmt.Println(u)
}
