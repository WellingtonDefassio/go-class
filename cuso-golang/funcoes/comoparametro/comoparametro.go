package main

func multiplicacao(a, b int) int {
	return a * b
}

func divisao(a, b int) int {
	return a / b
}

func exec(funcao func(int, int) int, p1 int, p2 int) int {
	return funcao(p1, p2)
}

func main() {

	mult := exec(multiplicacao, 10, 20)
	div := exec(divisao, 20, 10)

	println(mult, div)

}
