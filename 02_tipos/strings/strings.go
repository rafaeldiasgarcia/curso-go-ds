package main

import "fmt"

func main() {

	fmt.Println("O que você quiser!!")
	fmt.Println(`Aqui a gente pode escrever
múltiplas linhas
de
codigo!`)

	fmt.Println("Rafael tem", len("rafael"), "letras") // len() me devolve o tamanho da palavra

	fmt.Println("Rafael"[0]) // isso é um byte (int8)

	fmt.Println(string("Rafael"[0]))
	fmt.Println(string("Rafael"[1]))
	fmt.Println(string("Rafael"[2]))

}
