package main

import "fmt"

func main() {
	var numero float64
	println("Escreva um número: ")
	fmt.Scanf("%f", &numero)

	dobro := numero * 2

	fmt.Printf("O dobro do número %.2f é %.2f\n", numero, dobro)
}
