package main

import "fmt"

func f1() {
	fmt.Println("F1")
}

func f2(p1 string, p2 string) {
	fmt.Printf("P1:%s P2:%s\n", p1, p2)
}

func f3() string {
	return "retorno da função F3"
}

func f4(p1, p2 string) string {
	return fmt.Sprintf("F4: %s %s", p1, p2)
}

func f5() (string, string) {
	return "f5 retorno 1", "f5 retorno 2"
}

func main() {
	f1()
	f2("param1", "param2")
	r3, r4 := f3(), f4("param1", "param2")
	fmt.Println(r3, r4)
	s, s2 := f5()
	fmt.Println(s, s2)

}
