package main

import "fmt"

func main() {
	nome := "Rafael Dias Garcia"

	// chave ; valor ; durante ; item
	for _, v := range nome {
		fmt.Println(string(v))
	}
}
