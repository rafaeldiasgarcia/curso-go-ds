package main

import "fmt"

func main() {
	fmt.Println("Escreva um valor:")

	var x float64
	fmt.Scanf("%f", &x) // tenho que passar o endereço da variavel, usando o "&"

	fmt.Println("Você escolheu o valor", x)
}
