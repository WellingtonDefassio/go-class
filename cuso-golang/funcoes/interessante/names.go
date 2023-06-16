package main

import "strings"

func justLowerCase(funcao func(string) string, palavra string) string {
	toLower := funcao(palavra)
	opa := toLower + " opa"
	return strings.ToLower(opa)
}

func gritar(fala string) string {
	return fala
}

func main() {

	lowerCase := justLowerCase(gritar, "CADEEEEEEE")
	println(lowerCase)
}
