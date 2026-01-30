package main

import "fmt"

func main() {
	idades := make(map[string]int) // inicializa o mapa

	idades["rafael"] = 22
	idades["teo"] = 34
	fmt.Println("idades:", idades)

	alturas := map[string]float64{} // também declara um novo mapa
	alturas["rafael"] = 1.73
	fmt.Println("alturas:", alturas)

	alturaRafael := alturas["rafael"]
	fmt.Println("Altura rafael:", alturaRafael)

	// alturas["teo"] = 1.82
	alturaTeo, ok := alturas["teo"]
	if ok {
		fmt.Println("Altura teo:", alturaTeo)
	} else {
		fmt.Println("Não encontrei!")
	}

	// outra forma de tratar o ok
	if valor, ok := alturas["lara"]; ok {
		fmt.Println("Altura da lara é:", valor)
	} else {
		fmt.Println("Não encontrei!")
	}
}
