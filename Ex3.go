package main

import (
	"fmt"
)

func main() {
	var peso, altura float64

	fmt.Print("Digite o seu peso em Kg:")
	fmt.Scan(&peso)

	fmt.Print("Digite sua altura em metros:")
	fmt.Scan(&altura)

	IMC := peso / (altura * altura)

	fmt.Println(IMC)

}
