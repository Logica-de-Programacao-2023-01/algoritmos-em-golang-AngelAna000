package main

import (
	"fmt"
)

func main() {
	var preco int

	fmt.Print("Digite o preço do produto: ")
	fmt.Scan(&preco)

	desconto := (preco * 10) / 100
	precoComDesconto := preco - desconto

	fmt.Print("O preço com desconto é R$", precoComDesconto)
}
