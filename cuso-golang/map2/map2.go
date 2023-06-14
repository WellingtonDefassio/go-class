package main

import "fmt"

func main() {
	funcEndSalary := map[string]float64{}

	funcEndSalary["Wellington Defassio"] = 7700.50
	funcEndSalary["José da silva"] = 1250.70
	funcEndSalary["Matheus Carlos"] = 4433.85
	funcEndSalary["Valéria Assunção"] = 7300.5
	funcEndSalary["Wellington Defassio"] = 8250.73

	for name, salary := range funcEndSalary {

		fmt.Printf("O funcionario %s recebe um salário de R$%v por mês \n", name, salary)

	}

}
