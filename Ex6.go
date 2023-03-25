package main

import "fmt"

func main() {
	var salario float64

	fmt.Print("Digite o salario do funcionário:")
	fmt.Scan(&salario)

	NovoSalario := (salario*15)/100 + salario
	fmt.Print("O novo salário do funcionário comum aumento de 15% é R$", NovoSalario)

}
