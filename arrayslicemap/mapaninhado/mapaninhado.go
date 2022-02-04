package main

import "fmt"

func main() {
	funcsPorLetra := map[string]map[string]float64{
		"G": {
			"Gabriel": 123.75,
			"Guga":    789.23,
		},
		"J": {
			"José": 645.12,
		},
	}
	fmt.Println(funcsPorLetra)

	delete(funcsPorLetra, "J")

	fmt.Println(funcsPorLetra)
}
