package main

import "fmt"

type car struct {
	name  string
	speed int
}
type radio struct {
	name   string
	volume int
}

type ferrari struct {
	car
	radio
	turbo bool
}

func main() {
	f := ferrari{}

	f.car.name = "F40"
	f.radio.name = "AMBRIOSIO"
	f.speed = 0
	f.turbo = true

	fmt.Printf("A ferrari %s está com o turbo ligado? %v\n", f.car.name, f.turbo)
	fmt.Println(f)
}
