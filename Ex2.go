package main

import "fmt"

func main() {
	var num int

	fmt.Print("Digite seu número inteiro:")
	fmt.Scan(&num)

	dobro := num * 2
	triplo := num * 3
	quadruplo := num * 4

	fmt.Println("o dobro é ", dobro)
	fmt.Println("o dobro é ", triplo)
	fmt.Println("o dobro é ", quadruplo)

}
