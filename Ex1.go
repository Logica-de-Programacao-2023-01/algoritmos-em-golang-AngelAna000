package main

import "fmt"

func main() {

	var (
		num1 int
		num2 int
		num3 int
	)

	fmt.Print("Qual o primeiro numero?")
	fmt.Scan(&num1)
	fmt.Print("Qual o segundo numero?")
	fmt.Scan(&num2)
	fmt.Print("Qual o terceiro numero?")
	fmt.Scan(&num3)

	Soma := num1 + num2 + num3
	fmt.Println("A soma dos inteiros é:", Soma)
}
