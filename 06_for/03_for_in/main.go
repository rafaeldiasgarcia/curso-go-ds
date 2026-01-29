package main

import "fmt"

func main() {
	nome := "Rafael Dias Garcia"

	for i := range nome {
		fmt.Println(i, nome[i], string(nome[i]))
	}
}
