package main

import "fmt"

func main() {

	var (
		peso1 int
		peso2 int
		peso3 int
	)
	peso1 = 2
	peso2 = 3
	peso3 = 5
	var (
		num1 int
		num2 int
		num3 int
	)

	fmt.Print("Qual o primeiro número?")
	fmt.Scan(&num1)
	fmt.Print("QUal o segundo número?")
	fmt.Scan(&num2)
	fmt.Print("Qual o terceiro número?")
	fmt.Scan(&num3)

	média1 := (num1+num2+num3)/peso1 + peso2 + peso3

	fmt.Println("As médias ponderadas são respectivamente:", média1)
}
