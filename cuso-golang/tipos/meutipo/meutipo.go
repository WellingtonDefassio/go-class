package main

import "fmt"

type nota float64

func (n nota) criteria(start, end float64) bool {
	return float64(n) >= start && float64(n) <= end
}

func notaParaConceito(n nota) string {
	if n.criteria(9.0, 10.0) {
		return "A"
	} else if n.criteria(7.0, 8.99) {
		return "B"
	} else if n.criteria(5.0, 7.99) {
		return "C"
	} else if n.criteria(3.0, 4.99) {
		return "D"
	} else {
		return "E"
	}
}

func main() {
	fmt.Println(notaParaConceito(9.9))
	fmt.Println(notaParaConceito(4.9))
	fmt.Println(notaParaConceito(3.9))
	fmt.Println(notaParaConceito(2.9))
}
