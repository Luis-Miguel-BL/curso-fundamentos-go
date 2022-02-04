package main

import "fmt"

func teste() bool {
	const (
		a = 123
		b = "abc"
	)
	return false
}

func main() {
	x := teste()

	fmt.Println(x)
}
