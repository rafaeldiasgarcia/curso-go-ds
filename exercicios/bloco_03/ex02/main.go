package main

import "fmt"

func main() {
	fmt.Println("=== Venda de Água ===")
	fmt.Println("1 - Água mineral natural (R$ 1,50)")
	fmt.Println("2 - Água mineral com gás (R$ 2,50)")
	fmt.Print("Escolha o tipo de água: ")

	var escolha int
	fmt.Scanf("%d", &escolha)

	fmt.Print("Quantas garrafas deseja comprar? ")
	var quantidade int
	fmt.Scanf("%d", &quantidade)

	var precoUnitario float64

	switch escolha {
	case 1:
		precoUnitario = 1.50
		fmt.Printf("Você escolheu %d garrafa(s) de água mineral natural.\n", quantidade)

	case 2:
		precoUnitario = 2.50
		fmt.Printf("Você escolheu %d garrafa(s) de água mineral com gás.\n", quantidade)

	default:
		fmt.Println("Opção inválida!")
		return

	}

	total := precoUnitario * float64(quantidade)
	fmt.Printf("Total a pagar: R$ %.2f\n", total)
}
