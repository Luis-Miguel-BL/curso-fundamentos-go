package main

import "fmt"

func main() {
	i := 1

	var p *int

	p = &i //  pegando o endereço de memória da variável i
	*p++
	i++

	fmt.Println(p, *p, i, &i)
}
