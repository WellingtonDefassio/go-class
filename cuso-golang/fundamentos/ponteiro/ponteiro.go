package main

func main() {
	i := 1

	var ponteiro *int = nil

	ponteiro = &i

	println(*ponteiro, ponteiro, &i)
}
