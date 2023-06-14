package main

func main() {

	s := make([]int, 10)

	s[9] = 12
	println(&s[0])

	s = make([]int, 10, 20)
	println(&s[0])

	s = append(s, 1, 2, 3, 4, 5, 6, 7, 8, 9, 0)
	println(&s[0])

	s = append(s, 1, 15)

	println(&s[0])

}
