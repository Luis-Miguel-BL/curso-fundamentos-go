package main

import "fmt"

func main() {
	aprovados := make(map[int]string)

	aprovados[12345678] = "maria"
	aprovados[12345679] = "joão"
	aprovados[12345688] = "ana"
	fmt.Println(aprovados)
	for cpf, name := range aprovados {
		fmt.Printf("%s (CPF: %d) \n", name, cpf)
	}

	fmt.Println(aprovados[12345678])
	delete(aprovados, 12345678)
	fmt.Println(aprovados[12345678])

}
