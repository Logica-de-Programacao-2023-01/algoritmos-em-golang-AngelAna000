package main

import "fmt"

func main() {
	var num int
	fmt.Print("Digite um número inteiro:")

	fmt.Scan(&num)

	fmt.Println("O sucessor de", num, "é", num+1)

	fmt.Println("O antecessor de", num, "é", num-1)
}
