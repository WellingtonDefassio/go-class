package main

func imprimirResultado(nota float64) string {
	if nota >= 7 {
		return "Aprovado"
	}
	return "Reprovado"
}

func main() {

	println(imprimirResultado(6.5))

}
