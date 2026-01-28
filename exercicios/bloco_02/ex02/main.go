package main

import "fmt"

func main() {
	fmt.Println("Bom dia, qual seu nome?")

	var nome string
	fmt.Scanf("%s", &nome)

	fmt.Printf("É um prazer te conhecer, %s\n", nome)
}
