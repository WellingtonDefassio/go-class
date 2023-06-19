package main

import (
	"fmt"
	"time"
)

func doisTrezQuatroVezes(base int, c chan int) {
	time.Sleep(time.Second)
	c <- 2 * base
	time.Sleep(time.Second)
	c <- 3 * base
	time.Sleep(time.Second)
	c <- 4 * base
}

func main() {
	c := make(chan int)
	go doisTrezQuatroVezes(2, c)
	fmt.Println(<-c)
	fmt.Println(<-c)
	fmt.Println(<-c)
}
