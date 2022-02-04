package main

import (
	"fmt"
	"math"
	"reflect"
)

func main() {
	fmt.Println(1, 2, 1000)
	fmt.Println("Literal interiro é", reflect.TypeOf((32000)))

	var b byte = 255
	fmt.Println("O byte é", reflect.TypeOf(b))

	i1 := math.MaxInt64
	fmt.Println("O valor máximo do int é", i1)
	fmt.Println("O tipe de i1 é", reflect.TypeOf(i1))

	var i2 rune = 'a' // Representa um mapeamento da tabela Unicode (int32)
	fmt.Println("O rune é", reflect.TypeOf(i2))
	fmt.Println(i2)

	//números reais (float32, float64)
	var x float32 = 49.999 // por padrão é "tipado" como float64, caso deseje tipar com float32 é necessário passar manualmente ou usar var x = float32(49.999)
	fmt.Println("O tipo de x é", reflect.TypeOf(x))
	fmt.Println("O tipo literal de 49.99 é", reflect.TypeOf(49.99))

	bo := true
	fmt.Println("O tipo de bo é", reflect.TypeOf(bo))
	fmt.Println(!bo)

	s1 := "Olá"
	fmt.Println(s1 + "!")
	fmt.Println("O tamanho da sintr é", len(s1))

	s2 := `Ola
	tudo
	bem?`
	fmt.Println("O tamanho da string é", len(s2))

}
