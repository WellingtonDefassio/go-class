package main

import (
	"fmt"
	"math"
)

func main() {

	const PI float64 = 3.1415
	var raio = 3.2

	area := PI * math.Pow(raio, 2)

	fmt.Println("A area de circumference é ", area)

	var e, f bool = true, false

	fmt.Println(e, f)

	g, h, i := 2, false, "String"

	fmt.Println(g, h, i)

}
