package main

import "fmt"

func main() {
	nome := "Rafael Dias Garcia"

	for _, v := range nome {

		letra := string(v)
		if letra == "l" {
			break
		}

		if letra == " " {
			continue
		}

		fmt.Println(letra)
	}
}
