package main

import (
	"fmt"
	"time"
)

func rotina(ch chan int) {
	fmt.Println("exec 1")
	ch <- 1
	fmt.Println("exec 2")
	ch <- 2
	fmt.Println("exec 3")
	ch <- 3
	fmt.Println("exec 4")
	ch <- 4
	fmt.Println("exec 5")
	ch <- 5
}

func main() {
	ch := make(chan int, 2)
	go rotina(ch)

	time.Sleep(time.Second)
	fmt.Println(<-ch)
}
