package main

import "fmt"

// Faça um algoritmo
// que leia o peso de uma pessoa em quilos e converta para libras.

func main() {
	var peso float64
	fmt.Print("Qual seu peso em kg? ")
	fmt.Scan(&peso)

	resultado := peso * 2.2046
	fmt.Println("Seu peso em libras é", resultado)
}
