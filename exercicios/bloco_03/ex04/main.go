package main

import (
	"fmt"
	"strings"
)

func main() {
	fmt.Println("=== Verificador de Família ===")
	fmt.Print("Digite seu sobrenome: ")

	var sobrenome string
	fmt.Scanf("%s", &sobrenome)

	sobrenome = strings.ToLower(sobrenome)

	if sobrenome == "calvo" {
		fmt.Println("Você pertence à família Calvo!")

	} else {
		fmt.Println("Você não pertence à família Calvo.")

	}
}
