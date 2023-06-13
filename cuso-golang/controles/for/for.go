package main

import "time"

func main() {
	i := 1
	for i <= 10 {
		println(i)
		i++
	}
	for i := 0; i < 10; i++ {
		println(i)
	}

	for {
		println("Para sempre..")
		time.Sleep(time.Second)
	}

}
