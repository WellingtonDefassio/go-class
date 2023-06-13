package main

import (
	"fmt"
	"strconv"
)

func main() {

	x := 2.4
	y := 2
	fmt.Println(x / float64(y))

	num, _ := strconv.Atoi("123")

	fmt.Println(num)
}
