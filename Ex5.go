package main

import "fmt"

func main() {
	var IdadesAnos int

	fmt.Print("Digite sua idade em anos:")
	fmt.Scan(&IdadesAnos)

	IdadesDias := IdadesAnos * 365

	fmt.Print(" A idade em dias é:", IdadesDias)

}
