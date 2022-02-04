package main

import "fmt"

func compras(trab1, trab2 bool) (bool, bool, bool) {
	comprarTV50 := trab1 && trab2
	comprarTV32 := trab1 != trab2 // OU exclusivo
	comprarSorvete := trab1 || trab2

	return comprarTV50,
		comprarTV32,
		comprarSorvete

}

func main() {
	tv50, tv32, sorvete := compras(true, true)
	fmt.Printf("tv50 %t, tv32 %t, sorvete %t, saudável %t", tv50, tv32, sorvete, !sorvete)
}
