package main

import "fmt"

// Faça um algoritmo que leia o salário de um
// funcionário e calcule o seu novo salário com um aumento de 15%.

func main() {
	var salario float64
	fmt.Print("Qual é o salário? ")
	fmt.Scan(&salario)

	aumento := salario * 0.15
	total := salario + aumento
	fmt.Println("Seu salário com aumento é", total)
}
