package main

import "fmt"

func main() {
	fmt.Println("=== Sorveteria ===")

	fmt.Println("\nTipo de sorvete:")
	fmt.Println("1 - Casquinha (R$ 1,00)")
	fmt.Println("2 - Cascão (R$ 2,50)")
	fmt.Println("3 - Cestinha (R$ 4,00)")
	fmt.Print("Escolha o tipo: ")

	var tipoEscolha int
	fmt.Scanf("%d", &tipoEscolha)

	var precoTipo float64
	var nomeTipo string

	switch tipoEscolha {
	case 1:
		precoTipo = 1.00
		nomeTipo = "Casquinha"

	case 2:
		precoTipo = 2.50
		nomeTipo = "Cascão"

	case 3:
		precoTipo = 4.00
		nomeTipo = "Cestinha"

	default:
		fmt.Println("Tipo inválido!")
		return

	}

	fmt.Println("\nSabor do sorvete:")
	fmt.Println("1 - Morango")
	fmt.Println("2 - Creme")
	fmt.Println("3 - Chocolate")
	fmt.Print("Escolha o sabor: ")

	var saborEscolha int
	fmt.Scanf("%d", &saborEscolha)

	var nomeSabor string

	switch saborEscolha {
	case 1:
		nomeSabor = "Morango"

	case 2:
		nomeSabor = "Creme"

	case 3:
		nomeSabor = "Chocolate"

	default:
		fmt.Println("Sabor inválido!")
		return

	}

	fmt.Println("\nCobertura:")
	fmt.Println("1 - Caramelo (R$ 1,50)")
	fmt.Println("2 - Morango (R$ 1,50)")
	fmt.Println("3 - Chocolate (R$ 1,50)")
	fmt.Println("4 - Sem cobertura (R$ 0,00)")
	fmt.Print("Escolha a cobertura: ")

	var coberturaEscolha int
	fmt.Scanf("%d", &coberturaEscolha)

	var precoCobertura float64
	var nomeCobertura string

	switch coberturaEscolha {
	case 1:
		precoCobertura = 1.50
		nomeCobertura = "Caramelo"

	case 2:
		precoCobertura = 1.50
		nomeCobertura = "Morango"

	case 3:
		precoCobertura = 1.50
		nomeCobertura = "Chocolate"

	case 4:
		precoCobertura = 0.00
		nomeCobertura = "Sem cobertura"

	default:
		fmt.Println("Cobertura inválida!")
		return

	}

	total := precoTipo + precoCobertura
	fmt.Println("\n=== Resumo do Pedido ===")
	fmt.Printf("Tipo: %s\n", nomeTipo)
	fmt.Printf("Sabor: %s\n", nomeSabor)
	fmt.Printf("Cobertura: %s\n", nomeCobertura)
	fmt.Printf("\nTotal a pagar: R$ %.2f\n", total)
}
