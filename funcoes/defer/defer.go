package main

import "fmt"

func obeterValorAprovado(numero int) int {
	defer fmt.Println("fim")

	if numero > 5000 {
		fmt.Println("Valor alto..")
		return 5000
	} else {
		fmt.Println("Valor baixo..")
		return numero
	}
}

func main() {
	fmt.Println(obeterValorAprovado(50000))
	fmt.Println(obeterValorAprovado(2000))
}
