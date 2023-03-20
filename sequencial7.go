package main

import "fmt"

// Faça um algoritmo que
// leia um número inteiro e mostre o seu sucessor e antecessor.

func main() {
	var x int
	fmt.Print("Qual o número? ")
	fmt.Scan(&x)

	var antecessor = x - 1
	sucessor := x + 1

	fmt.Println("Seu antecessor é", antecessor, "Seu sucessor é", sucessor)
}
