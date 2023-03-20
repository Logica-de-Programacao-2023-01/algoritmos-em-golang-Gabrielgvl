package main

import "fmt"

//Faça um algoritmo
//que leia o preço de um produto e mostre o seu valor com desconto de 10%

func main() {
	var valorProduto float64
	fmt.Print("Qual o valor do produto?")
	fmt.Scan(&valorProduto)
	
	desconto := valorProduto * 0.1
	resultado := valorProduto - desconto
	fmt.Println("O valor com desconto é de", resultado)
}
