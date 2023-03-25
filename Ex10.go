package main

import "fmt"

func main() {
	var pesoKg float64

	fmt.Print("Digite o peso em quilos: ")
	fmt.Scan(&pesoKg)

	pesoLb := pesoKg * 2.2046

	fmt.Print("O peso em libras é:", pesoLb)
}
