package main

import "fmt"

//Faça um algoritmo
//que leia o número de dias trabalhados por um funcionário
//e o valor da sua diária e calcule o seu salário.

func main() {
	var numeroDias int
	fmt.Print("Quantos dias o funcionário trabalhou? ")
	fmt.Scan(&numeroDias)
	var valorDiaria float64
	fmt.Print("Qual o valor da diária do funcionário? ")
	fmt.Scan(&valorDiaria)

	salario := float64(numeroDias) * valorDiaria
	fmt.Println("O salário é", salario)
}
