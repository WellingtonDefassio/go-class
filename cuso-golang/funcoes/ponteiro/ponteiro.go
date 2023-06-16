package main

import "fmt"

func inc1(n int) {
	n++
}
func inc2(n *int) {
	*n++
}

func main() {

	i := 10
	inc1(i)
	fmt.Println(i)
	inc2(&i)
	fmt.Println(i)
}
