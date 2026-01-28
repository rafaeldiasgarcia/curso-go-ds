package main

import "fmt"

func main() {
	fmt.Println("=== Venda de Água ===")
	fmt.Println("1 - Água mineral natural (R$ 1,50)")
	fmt.Println("2 - Água mineral com gás (R$ 2,50)")
	fmt.Print("Escolha o tipo de água: ")

	var escolha int
	fmt.Scanf("%d", &escolha)

	var preco float64

	switch escolha {
	case 1:
		preco = 1.50
		fmt.Printf("Você escolheu água mineral natural. Total a pagar: R$ %.2f\n", preco)

	case 2:
		preco = 2.50
		fmt.Printf("Você escolheu água mineral com gás. Total a pagar: R$ %.2f\n", preco)

	default:
		fmt.Println("Opção inválida!")

	}
}
