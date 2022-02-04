package main

import "fmt"

func main() {
	ch := make(chan int, 1)

	ch <- 2 // enviando dados para um canal (escrita)

	<-ch // recebendo dados do canal(leitura)

	ch <- 3
	fmt.Println(<-ch)

}
