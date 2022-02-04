package main

import (
	"fmt"
	"time"
)

func fale(pessoa, texto string, qtde int) {
	for i := 0; i < qtde; i++ {
		time.Sleep(time.Second)
		fmt.Printf("%s:  %s (iteração %d)\n", pessoa, texto, i+1)
	}
}

func main() {
	// fale("maria", "teste maria", 3)
	// fale("joão", "teste joão", 1)

	go fale("maria", "teste maria 2 ", 500)
	go fale("joão", "teste joão 2", 500)

	fmt.Println("Fim!")

}
