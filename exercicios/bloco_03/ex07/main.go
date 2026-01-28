package main

import "fmt"

func main() {
	fmt.Println("=== Verificador de Média ===")
	fmt.Println("Digite as 3 notas:")

	var nota1, nota2, nota3 float64

	fmt.Print("Nota 1: ")
	fmt.Scanf("%f", &nota1)

	fmt.Print("Nota 2: ")
	fmt.Scanf("%f", &nota2)

	fmt.Print("Nota 3: ")
	fmt.Scanf("%f", &nota3)

	media := (nota1 + nota2 + nota3) / 3

	fmt.Printf("\nMédia: %.2f\n", media)

	if media >= 6.0 {
		fmt.Println("Aprovado! Parabéns!")

	} else {
		fmt.Printf("Reprovado. Faltaram %.2f pontos.\n", 6.0-media)

	}
}
