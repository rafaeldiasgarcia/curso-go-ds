package main

import (
	"fmt"
	"strings"
)

func main() {
	fmt.Println("=== Verificador de Item na Lista ===")
	fmt.Println("Itens disponíveis: laranja, cerveja, miojo, carvão, picanha")
	fmt.Print("Digite o item que deseja comprar: ")

	var item string
	fmt.Scanf("%s", &item)

	item = strings.ToLower(item)

	switch item {
	case "laranja", "cerveja", "miojo", "carvão", "carvao", "picanha":
		fmt.Printf("O item '%s' está disponível na loja!\n", item)

	default:
		fmt.Printf("O item '%s' não está disponível na loja.\n", item)

	}
}
