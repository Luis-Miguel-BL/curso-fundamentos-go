package main

import (
	"fmt"
	"strconv"
)

func main() {
	x := 2.4
	y := 2
	fmt.Println(x / float64(y)) // é preciso converter o tipo int para float64 para realizar as operações

	nota := 6.9
	notaFinal := int(nota)
	fmt.Println(notaFinal)

	fmt.Println("Teste ", string(97))

	//int para string
	fmt.Println("Teste " + strconv.Itoa(123))

	//string para int
	num, _ := strconv.Atoi("123")
	fmt.Println(num)

	b, _ := strconv.ParseBool("true") // para verdadeiro pode se usar "true" ou 1
	if b {
		fmt.Println("Verdadeiro")
	}
}
