package main

func swap(p1, p2 int) (second int, first int) {
	second = p2
	first = p1

	return second, first
}

func main() {

	second, first := swap(10, 20)

	println(second, first)

}
